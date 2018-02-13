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

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/m3db/m3db/storage/series (interfaces: DatabaseSeries,QueryableBlockRetriever)

package series

import (
	"time"

	"github.com/m3db/m3db/persist"
	"github.com/m3db/m3db/storage/block"
	"github.com/m3db/m3db/ts"
	"github.com/m3db/m3x/context"
	"github.com/m3db/m3x/ident"
	time0 "github.com/m3db/m3x/time"
	"github.com/m3db/m3db/x/xio"

	"github.com/golang/mock/gomock"
)

// Mock of DatabaseSeries interface
type MockDatabaseSeries struct {
	ctrl     *gomock.Controller
	recorder *_MockDatabaseSeriesRecorder
}

// Recorder for MockDatabaseSeries (not exported)
type _MockDatabaseSeriesRecorder struct {
	mock *MockDatabaseSeries
}

func NewMockDatabaseSeries(ctrl *gomock.Controller) *MockDatabaseSeries {
	mock := &MockDatabaseSeries{ctrl: ctrl}
	mock.recorder = &_MockDatabaseSeriesRecorder{mock}
	return mock
}

func (_m *MockDatabaseSeries) EXPECT() *_MockDatabaseSeriesRecorder {
	return _m.recorder
}

func (_m *MockDatabaseSeries) Bootstrap(_param0 block.DatabaseSeriesBlocks) error {
	ret := _m.ctrl.Call(_m, "Bootstrap", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDatabaseSeriesRecorder) Bootstrap(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Bootstrap", arg0)
}

func (_m *MockDatabaseSeries) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockDatabaseSeriesRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockDatabaseSeries) FetchBlocks(_param0 context.Context, _param1 []time.Time) ([]block.FetchBlockResult, error) {
	ret := _m.ctrl.Call(_m, "FetchBlocks", _param0, _param1)
	ret0, _ := ret[0].([]block.FetchBlockResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDatabaseSeriesRecorder) FetchBlocks(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBlocks", arg0, arg1)
}

func (_m *MockDatabaseSeries) FetchBlocksMetadata(_param0 context.Context, _param1 time.Time, _param2 time.Time, _param3 FetchBlocksMetadataOptions) block.FetchBlocksMetadataResult {
	ret := _m.ctrl.Call(_m, "FetchBlocksMetadata", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(block.FetchBlocksMetadataResult)
	return ret0
}

func (_mr *_MockDatabaseSeriesRecorder) FetchBlocksMetadata(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBlocksMetadata", arg0, arg1, arg2, arg3)
}

func (_m *MockDatabaseSeries) Flush(_param0 context.Context, _param1 time.Time, _param2 persist.Fn) error {
	ret := _m.ctrl.Call(_m, "Flush", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDatabaseSeriesRecorder) Flush(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Flush", arg0, arg1, arg2)
}

func (_m *MockDatabaseSeries) ID() ident.ID {
	ret := _m.ctrl.Call(_m, "ID")
	ret0, _ := ret[0].(ident.ID)
	return ret0
}

func (_mr *_MockDatabaseSeriesRecorder) ID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ID")
}

func (_m *MockDatabaseSeries) IsBootstrapped() bool {
	ret := _m.ctrl.Call(_m, "IsBootstrapped")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockDatabaseSeriesRecorder) IsBootstrapped() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsBootstrapped")
}

func (_m *MockDatabaseSeries) IsEmpty() bool {
	ret := _m.ctrl.Call(_m, "IsEmpty")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockDatabaseSeriesRecorder) IsEmpty() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsEmpty")
}

func (_m *MockDatabaseSeries) NumActiveBlocks() int {
	ret := _m.ctrl.Call(_m, "NumActiveBlocks")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockDatabaseSeriesRecorder) NumActiveBlocks() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NumActiveBlocks")
}

func (_m *MockDatabaseSeries) OnRetrieveBlock(_param0 ident.ID, _param1 time.Time, _param2 ts.Segment) {
	_m.ctrl.Call(_m, "OnRetrieveBlock", _param0, _param1, _param2)
}

func (_mr *_MockDatabaseSeriesRecorder) OnRetrieveBlock(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OnRetrieveBlock", arg0, arg1, arg2)
}

func (_m *MockDatabaseSeries) ReadEncoded(_param0 context.Context, _param1 time.Time, _param2 time.Time) ([][]xio.SegmentReader, error) {
	ret := _m.ctrl.Call(_m, "ReadEncoded", _param0, _param1, _param2)
	ret0, _ := ret[0].([][]xio.SegmentReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDatabaseSeriesRecorder) ReadEncoded(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadEncoded", arg0, arg1, arg2)
}

func (_m *MockDatabaseSeries) Reset(_param0 ident.ID, _param1 QueryableBlockRetriever, _param2 block.OnRetrieveBlock, _param3 Options) {
	_m.ctrl.Call(_m, "Reset", _param0, _param1, _param2, _param3)
}

func (_mr *_MockDatabaseSeriesRecorder) Reset(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset", arg0, arg1, arg2, arg3)
}

func (_m *MockDatabaseSeries) Tick() (TickResult, error) {
	ret := _m.ctrl.Call(_m, "Tick")
	ret0, _ := ret[0].(TickResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDatabaseSeriesRecorder) Tick() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Tick")
}

func (_m *MockDatabaseSeries) Write(_param0 context.Context, _param1 time.Time, _param2 float64, _param3 time0.Unit, _param4 []byte) error {
	ret := _m.ctrl.Call(_m, "Write", _param0, _param1, _param2, _param3, _param4)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDatabaseSeriesRecorder) Write(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0, arg1, arg2, arg3, arg4)
}

// Mock of QueryableBlockRetriever interface
type MockQueryableBlockRetriever struct {
	ctrl     *gomock.Controller
	recorder *_MockQueryableBlockRetrieverRecorder
}

// Recorder for MockQueryableBlockRetriever (not exported)
type _MockQueryableBlockRetrieverRecorder struct {
	mock *MockQueryableBlockRetriever
}

func NewMockQueryableBlockRetriever(ctrl *gomock.Controller) *MockQueryableBlockRetriever {
	mock := &MockQueryableBlockRetriever{ctrl: ctrl}
	mock.recorder = &_MockQueryableBlockRetrieverRecorder{mock}
	return mock
}

func (_m *MockQueryableBlockRetriever) EXPECT() *_MockQueryableBlockRetrieverRecorder {
	return _m.recorder
}

func (_m *MockQueryableBlockRetriever) IsBlockRetrievable(_param0 time.Time) bool {
	ret := _m.ctrl.Call(_m, "IsBlockRetrievable", _param0)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockQueryableBlockRetrieverRecorder) IsBlockRetrievable(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsBlockRetrievable", arg0)
}

func (_m *MockQueryableBlockRetriever) Stream(_param0 context.Context, _param1 ident.ID, _param2 time.Time, _param3 block.OnRetrieveBlock) (xio.SegmentReader, error) {
	ret := _m.ctrl.Call(_m, "Stream", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(xio.SegmentReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockQueryableBlockRetrieverRecorder) Stream(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stream", arg0, arg1, arg2, arg3)
}