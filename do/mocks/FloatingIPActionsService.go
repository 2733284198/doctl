// Code generated by MockGen. DO NOT EDIT.
// Source: floating_ip_actions.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	do "github.com/digitalocean/doctl/do"
	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

// MockFloatingIPActionsService is a mock of FloatingIPActionsService interface.
type MockFloatingIPActionsService struct {
	ctrl     *gomock.Controller
	recorder *MockFloatingIPActionsServiceMockRecorder
}

// MockFloatingIPActionsServiceMockRecorder is the mock recorder for MockFloatingIPActionsService.
type MockFloatingIPActionsServiceMockRecorder struct {
	mock *MockFloatingIPActionsService
}

// NewMockFloatingIPActionsService creates a new mock instance.
func NewMockFloatingIPActionsService(ctrl *gomock.Controller) *MockFloatingIPActionsService {
	mock := &MockFloatingIPActionsService{ctrl: ctrl}
	mock.recorder = &MockFloatingIPActionsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFloatingIPActionsService) EXPECT() *MockFloatingIPActionsServiceMockRecorder {
	return m.recorder
}

// Assign mocks base method.
func (m *MockFloatingIPActionsService) Assign(ip string, dropletID int) (*do.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Assign", ip, dropletID)
	ret0, _ := ret[0].(*do.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Assign indicates an expected call of Assign.
func (mr *MockFloatingIPActionsServiceMockRecorder) Assign(ip, dropletID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Assign", reflect.TypeOf((*MockFloatingIPActionsService)(nil).Assign), ip, dropletID)
}

// Unassign mocks base method.
func (m *MockFloatingIPActionsService) Unassign(ip string) (*do.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unassign", ip)
	ret0, _ := ret[0].(*do.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unassign indicates an expected call of Unassign.
func (mr *MockFloatingIPActionsServiceMockRecorder) Unassign(ip interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unassign", reflect.TypeOf((*MockFloatingIPActionsService)(nil).Unassign), ip)
}

// Get mocks base method.
func (m *MockFloatingIPActionsService) Get(ip string, actionID int) (*do.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ip, actionID)
	ret0, _ := ret[0].(*do.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockFloatingIPActionsServiceMockRecorder) Get(ip, actionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockFloatingIPActionsService)(nil).Get), ip, actionID)
}

// List mocks base method.
func (m *MockFloatingIPActionsService) List(ip string, opt *godo.ListOptions) ([]do.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ip, opt)
	ret0, _ := ret[0].([]do.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockFloatingIPActionsServiceMockRecorder) List(ip, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockFloatingIPActionsService)(nil).List), ip, opt)
}
