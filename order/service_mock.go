// Code generated by MockGen. DO NOT EDIT.
// Source: service.go
//
// Generated by this command:
//
//	mockgen -source service.go -destination service_mock.go -package order order Service
//

// Package order is a generated GoMock package.
package order

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetTotalOrdersSum mocks base method.
func (m *MockService) GetTotalOrdersSum() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalOrdersSum")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetTotalOrdersSum indicates an expected call of GetTotalOrdersSum.
func (mr *MockServiceMockRecorder) GetTotalOrdersSum() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalOrdersSum", reflect.TypeOf((*MockService)(nil).GetTotalOrdersSum))
}