// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/dbnode/retention/types.go

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

// Package retention is a generated GoMock package.
package retention

import (
	"reflect"
	"time"

	"github.com/golang/mock/gomock"
)

// MockOptions is a mock of Options interface
type MockOptions struct {
	ctrl     *gomock.Controller
	recorder *MockOptionsMockRecorder
}

// MockOptionsMockRecorder is the mock recorder for MockOptions
type MockOptionsMockRecorder struct {
	mock *MockOptions
}

// NewMockOptions creates a new mock instance
func NewMockOptions(ctrl *gomock.Controller) *MockOptions {
	mock := &MockOptions{ctrl: ctrl}
	mock.recorder = &MockOptionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOptions) EXPECT() *MockOptionsMockRecorder {
	return m.recorder
}

// Validate mocks base method
func (m *MockOptions) Validate() error {
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockOptionsMockRecorder) Validate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockOptions)(nil).Validate))
}

// Equal mocks base method
func (m *MockOptions) Equal(value Options) bool {
	ret := m.ctrl.Call(m, "Equal", value)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockOptionsMockRecorder) Equal(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockOptions)(nil).Equal), value)
}

// SetRetentionPeriod mocks base method
func (m *MockOptions) SetRetentionPeriod(value time.Duration) Options {
	ret := m.ctrl.Call(m, "SetRetentionPeriod", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetRetentionPeriod indicates an expected call of SetRetentionPeriod
func (mr *MockOptionsMockRecorder) SetRetentionPeriod(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRetentionPeriod", reflect.TypeOf((*MockOptions)(nil).SetRetentionPeriod), value)
}

// RetentionPeriod mocks base method
func (m *MockOptions) RetentionPeriod() time.Duration {
	ret := m.ctrl.Call(m, "RetentionPeriod")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// RetentionPeriod indicates an expected call of RetentionPeriod
func (mr *MockOptionsMockRecorder) RetentionPeriod() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetentionPeriod", reflect.TypeOf((*MockOptions)(nil).RetentionPeriod))
}

// SetBlockSize mocks base method
func (m *MockOptions) SetBlockSize(value time.Duration) Options {
	ret := m.ctrl.Call(m, "SetBlockSize", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetBlockSize indicates an expected call of SetBlockSize
func (mr *MockOptionsMockRecorder) SetBlockSize(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBlockSize", reflect.TypeOf((*MockOptions)(nil).SetBlockSize), value)
}

// BlockSize mocks base method
func (m *MockOptions) BlockSize() time.Duration {
	ret := m.ctrl.Call(m, "BlockSize")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// BlockSize indicates an expected call of BlockSize
func (mr *MockOptionsMockRecorder) BlockSize() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockSize", reflect.TypeOf((*MockOptions)(nil).BlockSize))
}

// SetBufferFuture mocks base method
func (m *MockOptions) SetBufferFuture(value time.Duration) Options {
	ret := m.ctrl.Call(m, "SetBufferFuture", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetBufferFuture indicates an expected call of SetBufferFuture
func (mr *MockOptionsMockRecorder) SetBufferFuture(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBufferFuture", reflect.TypeOf((*MockOptions)(nil).SetBufferFuture), value)
}

// BufferFuture mocks base method
func (m *MockOptions) BufferFuture() time.Duration {
	ret := m.ctrl.Call(m, "BufferFuture")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// BufferFuture indicates an expected call of BufferFuture
func (mr *MockOptionsMockRecorder) BufferFuture() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BufferFuture", reflect.TypeOf((*MockOptions)(nil).BufferFuture))
}

// SetBufferPast mocks base method
func (m *MockOptions) SetBufferPast(value time.Duration) Options {
	ret := m.ctrl.Call(m, "SetBufferPast", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetBufferPast indicates an expected call of SetBufferPast
func (mr *MockOptionsMockRecorder) SetBufferPast(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBufferPast", reflect.TypeOf((*MockOptions)(nil).SetBufferPast), value)
}

// BufferPast mocks base method
func (m *MockOptions) BufferPast() time.Duration {
	ret := m.ctrl.Call(m, "BufferPast")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// BufferPast indicates an expected call of BufferPast
func (mr *MockOptionsMockRecorder) BufferPast() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BufferPast", reflect.TypeOf((*MockOptions)(nil).BufferPast))
}

// SetBlockDataExpiry mocks base method
func (m *MockOptions) SetBlockDataExpiry(on bool) Options {
	ret := m.ctrl.Call(m, "SetBlockDataExpiry", on)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetBlockDataExpiry indicates an expected call of SetBlockDataExpiry
func (mr *MockOptionsMockRecorder) SetBlockDataExpiry(on interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBlockDataExpiry", reflect.TypeOf((*MockOptions)(nil).SetBlockDataExpiry), on)
}

// BlockDataExpiry mocks base method
func (m *MockOptions) BlockDataExpiry() bool {
	ret := m.ctrl.Call(m, "BlockDataExpiry")
	ret0, _ := ret[0].(bool)
	return ret0
}

// BlockDataExpiry indicates an expected call of BlockDataExpiry
func (mr *MockOptionsMockRecorder) BlockDataExpiry() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockDataExpiry", reflect.TypeOf((*MockOptions)(nil).BlockDataExpiry))
}

// SetBlockDataExpiryAfterNotAccessedPeriod mocks base method
func (m *MockOptions) SetBlockDataExpiryAfterNotAccessedPeriod(period time.Duration) Options {
	ret := m.ctrl.Call(m, "SetBlockDataExpiryAfterNotAccessedPeriod", period)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetBlockDataExpiryAfterNotAccessedPeriod indicates an expected call of SetBlockDataExpiryAfterNotAccessedPeriod
func (mr *MockOptionsMockRecorder) SetBlockDataExpiryAfterNotAccessedPeriod(period interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBlockDataExpiryAfterNotAccessedPeriod", reflect.TypeOf((*MockOptions)(nil).SetBlockDataExpiryAfterNotAccessedPeriod), period)
}

// BlockDataExpiryAfterNotAccessedPeriod mocks base method
func (m *MockOptions) BlockDataExpiryAfterNotAccessedPeriod() time.Duration {
	ret := m.ctrl.Call(m, "BlockDataExpiryAfterNotAccessedPeriod")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// BlockDataExpiryAfterNotAccessedPeriod indicates an expected call of BlockDataExpiryAfterNotAccessedPeriod
func (mr *MockOptionsMockRecorder) BlockDataExpiryAfterNotAccessedPeriod() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockDataExpiryAfterNotAccessedPeriod", reflect.TypeOf((*MockOptions)(nil).BlockDataExpiryAfterNotAccessedPeriod))
}

// SetColdWritesEnabled mocks base method
func (m *MockOptions) SetColdWritesEnabled(value bool) Options {
	ret := m.ctrl.Call(m, "SetColdWritesEnabled", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetColdWritesEnabled indicates an expected call of SetColdWritesEnabled
func (mr *MockOptionsMockRecorder) SetColdWritesEnabled(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetColdWritesEnabled", reflect.TypeOf((*MockOptions)(nil).SetColdWritesEnabled), value)
}

// ColdWritesEnabled mocks base method
func (m *MockOptions) ColdWritesEnabled() bool {
	ret := m.ctrl.Call(m, "ColdWritesEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// ColdWritesEnabled indicates an expected call of ColdWritesEnabled
func (mr *MockOptionsMockRecorder) ColdWritesEnabled() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ColdWritesEnabled", reflect.TypeOf((*MockOptions)(nil).ColdWritesEnabled))
}
