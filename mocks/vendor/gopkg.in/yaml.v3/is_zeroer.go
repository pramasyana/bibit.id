// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IsZeroer is an autogenerated mock type for the IsZeroer type
type IsZeroer struct {
	mock.Mock
}

// IsZero provides a mock function with given fields:
func (_m *IsZeroer) IsZero() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
