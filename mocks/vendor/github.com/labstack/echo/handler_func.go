// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo"
	mock "github.com/stretchr/testify/mock"
)

// HandlerFunc is an autogenerated mock type for the HandlerFunc type
type HandlerFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *HandlerFunc) Execute(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
