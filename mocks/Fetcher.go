// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// Fetcher is an autogenerated mock type for the Fetcher type
type Fetcher struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: url
func (_m *Fetcher) Fetch(url string) (io.Reader, error) {
	ret := _m.Called(url)

	var r0 io.Reader
	if rf, ok := ret.Get(0).(func(string) io.Reader); ok {
		r0 = rf(url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.Reader)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}