// Code generated by MockGen. DO NOT EDIT.
// Source: droplets.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	do "github.com/digitalocean/doctl/do"
	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

// MockDropletsService is a mock of DropletsService interface.
type MockDropletsService struct {
	ctrl     *gomock.Controller
	recorder *MockDropletsServiceMockRecorder
}

// MockDropletsServiceMockRecorder is the mock recorder for MockDropletsService.
type MockDropletsServiceMockRecorder struct {
	mock *MockDropletsService
}

// NewMockDropletsService creates a new mock instance.
func NewMockDropletsService(ctrl *gomock.Controller) *MockDropletsService {
	mock := &MockDropletsService{ctrl: ctrl}
	mock.recorder = &MockDropletsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDropletsService) EXPECT() *MockDropletsServiceMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockDropletsService) List() (do.Droplets, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].(do.Droplets)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockDropletsServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDropletsService)(nil).List))
}

// ListByTag mocks base method.
func (m *MockDropletsService) ListByTag(arg0 string) (do.Droplets, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByTag", arg0)
	ret0, _ := ret[0].(do.Droplets)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByTag indicates an expected call of ListByTag.
func (mr *MockDropletsServiceMockRecorder) ListByTag(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByTag", reflect.TypeOf((*MockDropletsService)(nil).ListByTag), arg0)
}

// Get mocks base method.
func (m *MockDropletsService) Get(arg0 int) (*do.Droplet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*do.Droplet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDropletsServiceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDropletsService)(nil).Get), arg0)
}

// Create mocks base method.
func (m *MockDropletsService) Create(arg0 *godo.DropletCreateRequest, arg1 bool) (*do.Droplet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*do.Droplet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockDropletsServiceMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDropletsService)(nil).Create), arg0, arg1)
}

// CreateMultiple mocks base method.
func (m *MockDropletsService) CreateMultiple(arg0 *godo.DropletMultiCreateRequest) (do.Droplets, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMultiple", arg0)
	ret0, _ := ret[0].(do.Droplets)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMultiple indicates an expected call of CreateMultiple.
func (mr *MockDropletsServiceMockRecorder) CreateMultiple(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMultiple", reflect.TypeOf((*MockDropletsService)(nil).CreateMultiple), arg0)
}

// Delete mocks base method.
func (m *MockDropletsService) Delete(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDropletsServiceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDropletsService)(nil).Delete), arg0)
}

// DeleteByTag mocks base method.
func (m *MockDropletsService) DeleteByTag(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByTag", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByTag indicates an expected call of DeleteByTag.
func (mr *MockDropletsServiceMockRecorder) DeleteByTag(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByTag", reflect.TypeOf((*MockDropletsService)(nil).DeleteByTag), arg0)
}

// Kernels mocks base method.
func (m *MockDropletsService) Kernels(arg0 int) (do.Kernels, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Kernels", arg0)
	ret0, _ := ret[0].(do.Kernels)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Kernels indicates an expected call of Kernels.
func (mr *MockDropletsServiceMockRecorder) Kernels(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kernels", reflect.TypeOf((*MockDropletsService)(nil).Kernels), arg0)
}

// Snapshots mocks base method.
func (m *MockDropletsService) Snapshots(arg0 int) (do.Images, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshots", arg0)
	ret0, _ := ret[0].(do.Images)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Snapshots indicates an expected call of Snapshots.
func (mr *MockDropletsServiceMockRecorder) Snapshots(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshots", reflect.TypeOf((*MockDropletsService)(nil).Snapshots), arg0)
}

// Backups mocks base method.
func (m *MockDropletsService) Backups(arg0 int) (do.Images, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Backups", arg0)
	ret0, _ := ret[0].(do.Images)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Backups indicates an expected call of Backups.
func (mr *MockDropletsServiceMockRecorder) Backups(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Backups", reflect.TypeOf((*MockDropletsService)(nil).Backups), arg0)
}

// Actions mocks base method.
func (m *MockDropletsService) Actions(arg0 int) (do.Actions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Actions", arg0)
	ret0, _ := ret[0].(do.Actions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Actions indicates an expected call of Actions.
func (mr *MockDropletsServiceMockRecorder) Actions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Actions", reflect.TypeOf((*MockDropletsService)(nil).Actions), arg0)
}

// Neighbors mocks base method.
func (m *MockDropletsService) Neighbors(arg0 int) (do.Droplets, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Neighbors", arg0)
	ret0, _ := ret[0].(do.Droplets)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Neighbors indicates an expected call of Neighbors.
func (mr *MockDropletsServiceMockRecorder) Neighbors(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Neighbors", reflect.TypeOf((*MockDropletsService)(nil).Neighbors), arg0)
}
