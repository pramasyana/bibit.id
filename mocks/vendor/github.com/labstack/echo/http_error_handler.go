// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo"
	mock "github.com/stretchr/testify/mock"
)

// HTTPErrorHandler is an autogenerated mock type for the HTTPErrorHandler type
type HTTPErrorHandler struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *HTTPErrorHandler) Execute(_a0 error, _a1 echo.Context) {
	_m.Called(_a0, _a1)
}
