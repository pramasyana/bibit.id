// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// TestingT is an autogenerated mock type for the TestingT type
type TestingT struct {
	mock.Mock
}

// Errorf provides a mock function with given fields: format, args
func (_m *TestingT) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// FailNow provides a mock function with given fields:
func (_m *TestingT) FailNow() {
	_m.Called()
}

// Logf provides a mock function with given fields: format, args
func (_m *TestingT) Logf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}
