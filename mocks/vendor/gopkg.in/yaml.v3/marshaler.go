// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Marshaler is an autogenerated mock type for the Marshaler type
type Marshaler struct {
	mock.Mock
}

// MarshalYAML provides a mock function with given fields:
func (_m *Marshaler) MarshalYAML() (interface{}, error) {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
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