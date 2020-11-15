// Code generated by MockGen. DO NOT EDIT.
// Source: application/usecase/usecase.go

// Package mock_usecase is a generated GoMock package.
package usecase_mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Sample mocks base method.
func (m *MockInterface) Sample() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sample")
	ret0, _ := ret[0].(string)
	return ret0
}

// Sample indicates an expected call of Sample.
func (mr *MockInterfaceMockRecorder) Sample() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sample", reflect.TypeOf((*MockInterface)(nil).Sample))
}
