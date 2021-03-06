// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spiffe/spire/proto/spire/common/hostservices (interfaces: MetricsService)

// Package mock_hostservices is a generated GoMock package.
package mock_hostservices

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	hostservices "github.com/spiffe/spire/proto/spire/common/hostservices"
	reflect "reflect"
)

// MockMetricsService is a mock of MetricsService interface
type MockMetricsService struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsServiceMockRecorder
}

// MockMetricsServiceMockRecorder is the mock recorder for MockMetricsService
type MockMetricsServiceMockRecorder struct {
	mock *MockMetricsService
}

// NewMockMetricsService creates a new mock instance
func NewMockMetricsService(ctrl *gomock.Controller) *MockMetricsService {
	mock := &MockMetricsService{ctrl: ctrl}
	mock.recorder = &MockMetricsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMetricsService) EXPECT() *MockMetricsServiceMockRecorder {
	return m.recorder
}

// AddSample mocks base method
func (m *MockMetricsService) AddSample(arg0 context.Context, arg1 *hostservices.AddSampleRequest) (*hostservices.AddSampleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSample", arg0, arg1)
	ret0, _ := ret[0].(*hostservices.AddSampleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSample indicates an expected call of AddSample
func (mr *MockMetricsServiceMockRecorder) AddSample(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSample", reflect.TypeOf((*MockMetricsService)(nil).AddSample), arg0, arg1)
}

// EmitKey mocks base method
func (m *MockMetricsService) EmitKey(arg0 context.Context, arg1 *hostservices.EmitKeyRequest) (*hostservices.EmitKeyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EmitKey", arg0, arg1)
	ret0, _ := ret[0].(*hostservices.EmitKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EmitKey indicates an expected call of EmitKey
func (mr *MockMetricsServiceMockRecorder) EmitKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EmitKey", reflect.TypeOf((*MockMetricsService)(nil).EmitKey), arg0, arg1)
}

// IncrCounter mocks base method
func (m *MockMetricsService) IncrCounter(arg0 context.Context, arg1 *hostservices.IncrCounterRequest) (*hostservices.IncrCounterResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrCounter", arg0, arg1)
	ret0, _ := ret[0].(*hostservices.IncrCounterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IncrCounter indicates an expected call of IncrCounter
func (mr *MockMetricsServiceMockRecorder) IncrCounter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrCounter", reflect.TypeOf((*MockMetricsService)(nil).IncrCounter), arg0, arg1)
}

// MeasureSince mocks base method
func (m *MockMetricsService) MeasureSince(arg0 context.Context, arg1 *hostservices.MeasureSinceRequest) (*hostservices.MeasureSinceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MeasureSince", arg0, arg1)
	ret0, _ := ret[0].(*hostservices.MeasureSinceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MeasureSince indicates an expected call of MeasureSince
func (mr *MockMetricsServiceMockRecorder) MeasureSince(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MeasureSince", reflect.TypeOf((*MockMetricsService)(nil).MeasureSince), arg0, arg1)
}

// SetGauge mocks base method
func (m *MockMetricsService) SetGauge(arg0 context.Context, arg1 *hostservices.SetGaugeRequest) (*hostservices.SetGaugeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetGauge", arg0, arg1)
	ret0, _ := ret[0].(*hostservices.SetGaugeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetGauge indicates an expected call of SetGauge
func (mr *MockMetricsServiceMockRecorder) SetGauge(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGauge", reflect.TypeOf((*MockMetricsService)(nil).SetGauge), arg0, arg1)
}
