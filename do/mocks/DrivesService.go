package mocks

import "github.com/digitalocean/doctl/do"
import "github.com/stretchr/testify/mock"

import "github.com/digitalocean/godo"

// Generated: please do not edit by hand

type DrivesService struct {
	mock.Mock
}

func (_m *DrivesService) List() ([]do.Drive, error) {
	ret := _m.Called()

	var r0 []do.Drive
	if rf, ok := ret.Get(0).(func() []do.Drive); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]do.Drive)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *DrivesService) CreateDrive(_a0 *godo.DriveCreateRequest) (*do.Drive, error) {
	ret := _m.Called(_a0)

	var r0 *do.Drive
	if rf, ok := ret.Get(0).(func(*godo.DriveCreateRequest) *do.Drive); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.Drive)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*godo.DriveCreateRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *DrivesService) DeleteDrive(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *DrivesService) Get(_a0 string) (*do.Drive, error) {
	ret := _m.Called(_a0)

	var r0 *do.Drive
	if rf, ok := ret.Get(0).(func(string) *do.Drive); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.Drive)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
