package mocks

import "github.com/digitalocean/doctl/do"
import "github.com/stretchr/testify/mock"

import "github.com/digitalocean/godo"

// Generated: please do not edit by hand

type ImagesService struct {
	mock.Mock
}

func (_m *ImagesService) List(public bool) (do.Images, error) {
	ret := _m.Called(public)

	var r0 do.Images
	if rf, ok := ret.Get(0).(func(bool) do.Images); ok {
		r0 = rf(public)
	} else {
		r0 = ret.Get(0).(do.Images)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bool) error); ok {
		r1 = rf(public)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *ImagesService) ListDistribution(public bool) (do.Images, error) {
	ret := _m.Called(public)

	var r0 do.Images
	if rf, ok := ret.Get(0).(func(bool) do.Images); ok {
		r0 = rf(public)
	} else {
		r0 = ret.Get(0).(do.Images)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bool) error); ok {
		r1 = rf(public)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *ImagesService) ListApplication(public bool) (do.Images, error) {
	ret := _m.Called(public)

	var r0 do.Images
	if rf, ok := ret.Get(0).(func(bool) do.Images); ok {
		r0 = rf(public)
	} else {
		r0 = ret.Get(0).(do.Images)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bool) error); ok {
		r1 = rf(public)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *ImagesService) ListUser(public bool) (do.Images, error) {
	ret := _m.Called(public)

	var r0 do.Images
	if rf, ok := ret.Get(0).(func(bool) do.Images); ok {
		r0 = rf(public)
	} else {
		r0 = ret.Get(0).(do.Images)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bool) error); ok {
		r1 = rf(public)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *ImagesService) GetByID(id int) (*do.Image, error) {
	ret := _m.Called(id)

	var r0 *do.Image
	if rf, ok := ret.Get(0).(func(int) *do.Image); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.Image)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *ImagesService) GetBySlug(slug string) (*do.Image, error) {
	ret := _m.Called(slug)

	var r0 *do.Image
	if rf, ok := ret.Get(0).(func(string) *do.Image); ok {
		r0 = rf(slug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.Image)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *ImagesService) Update(id int, iur *godo.ImageUpdateRequest) (*do.Image, error) {
	ret := _m.Called(id, iur)

	var r0 *do.Image
	if rf, ok := ret.Get(0).(func(int, *godo.ImageUpdateRequest) *do.Image); ok {
		r0 = rf(id, iur)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.Image)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, *godo.ImageUpdateRequest) error); ok {
		r1 = rf(id, iur)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *ImagesService) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
