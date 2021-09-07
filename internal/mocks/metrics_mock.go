// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozonva/ova-meeting-api/internal/metrics (interfaces: Metrics)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMetrics is a mock of Metrics interface
type MockMetrics struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsMockRecorder
}

// MockMetricsMockRecorder is the mock recorder for MockMetrics
type MockMetricsMockRecorder struct {
	mock *MockMetrics
}

// NewMockMetrics creates a new mock instance
func NewMockMetrics(ctrl *gomock.Controller) *MockMetrics {
	mock := &MockMetrics{ctrl: ctrl}
	mock.recorder = &MockMetricsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMetrics) EXPECT() *MockMetricsMockRecorder {
	return m.recorder
}

// IncCreate mocks base method
func (m *MockMetrics) IncCreate() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncCreate")
}

// IncCreate indicates an expected call of IncCreate
func (mr *MockMetricsMockRecorder) IncCreate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncCreate", reflect.TypeOf((*MockMetrics)(nil).IncCreate))
}

// IncDelete mocks base method
func (m *MockMetrics) IncDelete() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncDelete")
}

// IncDelete indicates an expected call of IncDelete
func (mr *MockMetricsMockRecorder) IncDelete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncDelete", reflect.TypeOf((*MockMetrics)(nil).IncDelete))
}

// IncUpdate mocks base method
func (m *MockMetrics) IncUpdate() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncUpdate")
}

// IncUpdate indicates an expected call of IncUpdate
func (mr *MockMetricsMockRecorder) IncUpdate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncUpdate", reflect.TypeOf((*MockMetrics)(nil).IncUpdate))
}