// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package m3msg

import (
	"bytes"
	"context"
	"io"
	"sync/atomic"

	"github.com/m3db/m3/src/metrics/encoding/msgpack"
	"github.com/m3db/m3/src/msg/consumer"
	"github.com/m3db/m3x/instrument"
	"github.com/m3db/m3x/log"
	"github.com/m3db/m3x/pool"

	"github.com/uber-go/tally"
)

// Options for the ingest handler.
type Options struct {
	InstrumentOptions          instrument.Options
	WriteFn                    WriteFn
	AggregatedIteratorOptions  msgpack.AggregatedIteratorOptions
	ProtobufDecoderPoolOptions pool.ObjectPoolOptions
}

type handlerMetrics struct {
	messageReadError             tally.Counter
	metricAccepted               tally.Counter
	droppedMetricDecodeError     tally.Counter
	droppedMetricDecodeMalformed tally.Counter
}

func newHandlerMetrics(scope tally.Scope) handlerMetrics {
	messageScope := scope.SubScope("metric")
	return handlerMetrics{
		messageReadError: scope.Counter("message-read-error"),
		metricAccepted:   messageScope.Counter("accepted"),
		droppedMetricDecodeError: messageScope.Tagged(map[string]string{
			"reason": "decode-error",
		}).Counter("dropped"),
		droppedMetricDecodeMalformed: messageScope.Tagged(map[string]string{
			"reason": "decode-malformed",
		}).Counter("dropped"),
	}
}

type msgpackHandler struct {
	writeFn      WriteFn
	iteratorOpts msgpack.AggregatedIteratorOptions
	logger       log.Logger
	m            handlerMetrics
}

func newMsgpackHandler(opts Options) (*msgpackHandler, error) {
	return &msgpackHandler{
		writeFn:      opts.WriteFn,
		iteratorOpts: opts.AggregatedIteratorOptions,
		logger:       opts.InstrumentOptions.Logger(),
		m:            newHandlerMetrics(opts.InstrumentOptions.MetricsScope()),
	}, nil
}

func (h *msgpackHandler) Handle(c consumer.Consumer) {
	h.newPerConsumerHandler().handle(c)
}

func (h *msgpackHandler) newPerConsumerHandler() *perConsumerHandler {
	return &perConsumerHandler{
		ctx:     context.Background(),
		writeFn: h.writeFn,
		logger:  h.logger,
		m:       h.m,
		it:      msgpack.NewAggregatedIterator(nil, h.iteratorOpts),
		r:       bytes.NewReader(nil),
	}
}

type perConsumerHandler struct {
	// Per server variables, shared across consumers/connections.
	ctx     context.Context
	writeFn WriteFn
	logger  log.Logger
	m       handlerMetrics

	// Per consumer/connection variables.
	it msgpack.AggregatedIterator
	r  *bytes.Reader
}

func (h *perConsumerHandler) handle(c consumer.Consumer) {
	var (
		msgErr error
		msg    consumer.Message
	)
	for {
		msg, msgErr = c.Message()
		if msgErr != nil {
			break
		}

		h.processMessage(msg)
	}
	if msgErr != nil && msgErr != io.EOF {
		h.logger.WithFields(log.NewErrField(msgErr)).Errorf("could not read message from consumer")
		h.m.messageReadError.Inc(1)
	}
	c.Close()
	h.it.Close()
}

func (h *perConsumerHandler) processMessage(
	msg consumer.Message,
) {
	r := NewRefCountedCallback(msg)
	r.IncRef()

	// Decode the bytes in the message.
	h.r.Reset(msg.Bytes())
	h.it.Reset(h.r)
	for h.it.Next() {
		raw, sp, encodeNanos := h.it.Value()
		m, err := raw.Metric()
		if err != nil {
			h.logger.WithFields(log.NewErrField(err)).Error("invalid raw metric")
			h.m.droppedMetricDecodeMalformed.Inc(1)
			continue
		}

		h.m.metricAccepted.Inc(1)

		// TODO: Consider incrementing a wait group for each write and wait on
		// shut down to reduce the number of messages being retried by m3msg.
		r.IncRef()
		h.writeFn(h.ctx, m.ID, m.TimeNanos, encodeNanos, m.Value, sp, r)
	}
	r.decRef()
	if err := h.it.Err(); err != nil && err != io.EOF {
		h.logger.WithFields(log.NewErrField(h.it.Err())).Errorf("could not decode msg %s", msg.Bytes())
		h.m.droppedMetricDecodeError.Inc(1)
	}
}

// RefCountedCallback wraps a message with a reference count, the message will
// be acked once the reference count decrements to zero.
// The implementation is thread safe.
type RefCountedCallback struct {
	ref int32
	msg consumer.Message
}

// NewRefCountedCallback creates a RefCountedCallback.
func NewRefCountedCallback(msg consumer.Message) *RefCountedCallback {
	return &RefCountedCallback{
		ref: 0,
		msg: msg,
	}
}

// IncRef increments the ref count atomically.
func (r *RefCountedCallback) IncRef() {
	atomic.AddInt32(&r.ref, 1)
}

// decRef decrements the ref count atomically. If the decrement causes the
// reference count to reach zero, then the message will be acked.
func (r *RefCountedCallback) decRef() {
	ref := atomic.AddInt32(&r.ref, -1)
	if ref == 0 {
		r.msg.Ack()
		return
	}
	if ref < 0 {
		panic("invalid ref count")
	}
}

// Callback performs the callback.
func (r *RefCountedCallback) Callback(t CallbackType) {
	switch t {
	case OnSuccess, OnNonRetriableError:
		r.decRef()
	}
}
