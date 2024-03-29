// Code generated by MockGen. DO NOT EDIT.
// Source: internal/queue/handler.go

// Package testutil is a generated GoMock package.
package testutil

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	taskq "github.com/vmihailenco/taskq/v3"
)

// MockHandler is a mock of Handler interface.
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler.
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance.
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// AddTask mocks base method.
func (m *MockHandler) AddTask(t *taskq.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTask", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTask indicates an expected call of AddTask.
func (mr *MockHandlerMockRecorder) AddTask(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTask", reflect.TypeOf((*MockHandler)(nil).AddTask), t)
}

// StartConsumer mocks base method.
func (m *MockHandler) StartConsumer() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartConsumer")
	ret0, _ := ret[0].(error)
	return ret0
}

// StartConsumer indicates an expected call of StartConsumer.
func (mr *MockHandlerMockRecorder) StartConsumer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartConsumer", reflect.TypeOf((*MockHandler)(nil).StartConsumer))
}
