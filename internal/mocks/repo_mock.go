// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozonva/ova-meeting-api/internal/repo (interfaces: MeetingRepo)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	models "github.com/ozonva/ova-meeting-api/internal/models"
	reflect "reflect"
)

// MockMeetingRepo is a mock of MeetingRepo interface
type MockMeetingRepo struct {
	ctrl     *gomock.Controller
	recorder *MockMeetingRepoMockRecorder
}

// MockMeetingRepoMockRecorder is the mock recorder for MockMeetingRepo
type MockMeetingRepoMockRecorder struct {
	mock *MockMeetingRepo
}

// NewMockMeetingRepo creates a new mock instance
func NewMockMeetingRepo(ctrl *gomock.Controller) *MockMeetingRepo {
	mock := &MockMeetingRepo{ctrl: ctrl}
	mock.recorder = &MockMeetingRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMeetingRepo) EXPECT() *MockMeetingRepoMockRecorder {
	return m.recorder
}

// AddMeetings mocks base method
func (m *MockMeetingRepo) AddMeetings(arg0 []models.Meeting) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMeetings", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMeetings indicates an expected call of AddMeetings
func (mr *MockMeetingRepoMockRecorder) AddMeetings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMeetings", reflect.TypeOf((*MockMeetingRepo)(nil).AddMeetings), arg0)
}

// DescribeMeeting mocks base method
func (m *MockMeetingRepo) DescribeMeeting(arg0 uuid.UUID) (*models.Meeting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeMeeting", arg0)
	ret0, _ := ret[0].(*models.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeMeeting indicates an expected call of DescribeMeeting
func (mr *MockMeetingRepoMockRecorder) DescribeMeeting(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeMeeting", reflect.TypeOf((*MockMeetingRepo)(nil).DescribeMeeting), arg0)
}

// ListMeetings mocks base method
func (m *MockMeetingRepo) ListMeetings(arg0, arg1 uint64) ([]models.Meeting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMeetings", arg0, arg1)
	ret0, _ := ret[0].([]models.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMeetings indicates an expected call of ListMeetings
func (mr *MockMeetingRepoMockRecorder) ListMeetings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMeetings", reflect.TypeOf((*MockMeetingRepo)(nil).ListMeetings), arg0, arg1)
}