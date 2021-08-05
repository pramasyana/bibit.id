// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	io "io"

	log "github.com/labstack/gommon/log"
	mock "github.com/stretchr/testify/mock"
)

// Logger is an autogenerated mock type for the Logger type
type Logger struct {
	mock.Mock
}

// Debug provides a mock function with given fields: i
func (_m *Logger) Debug(i ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, i...)
	_m.Called(_ca...)
}

// Debugf provides a mock function with given fields: format, args
func (_m *Logger) Debugf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Debugj provides a mock function with given fields: j
func (_m *Logger) Debugj(j log.JSON) {
	_m.Called(j)
}

// Error provides a mock function with given fields: i
func (_m *Logger) Error(i ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, i...)
	_m.Called(_ca...)
}

// Errorf provides a mock function with given fields: format, args
func (_m *Logger) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Errorj provides a mock function with given fields: j
func (_m *Logger) Errorj(j log.JSON) {
	_m.Called(j)
}

// Fatal provides a mock function with given fields: i
func (_m *Logger) Fatal(i ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, i...)
	_m.Called(_ca...)
}

// Fatalf provides a mock function with given fields: format, args
func (_m *Logger) Fatalf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Fatalj provides a mock function with given fields: j
func (_m *Logger) Fatalj(j log.JSON) {
	_m.Called(j)
}

// Info provides a mock function with given fields: i
func (_m *Logger) Info(i ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, i...)
	_m.Called(_ca...)
}

// Infof provides a mock function with given fields: format, args
func (_m *Logger) Infof(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Infoj provides a mock function with given fields: j
func (_m *Logger) Infoj(j log.JSON) {
	_m.Called(j)
}

// Level provides a mock function with given fields:
func (_m *Logger) Level() log.Lvl {
	ret := _m.Called()

	var r0 log.Lvl
	if rf, ok := ret.Get(0).(func() log.Lvl); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(log.Lvl)
	}

	return r0
}

// Output provides a mock function with given fields:
func (_m *Logger) Output() io.Writer {
	ret := _m.Called()

	var r0 io.Writer
	if rf, ok := ret.Get(0).(func() io.Writer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.Writer)
		}
	}

	return r0
}

// Panic provides a mock function with given fields: i
func (_m *Logger) Panic(i ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, i...)
	_m.Called(_ca...)
}

// Panicf provides a mock function with given fields: format, args
func (_m *Logger) Panicf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Panicj provides a mock function with given fields: j
func (_m *Logger) Panicj(j log.JSON) {
	_m.Called(j)
}

// Prefix provides a mock function with given fields:
func (_m *Logger) Prefix() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Print provides a mock function with given fields: i
func (_m *Logger) Print(i ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, i...)
	_m.Called(_ca...)
}

// Printf provides a mock function with given fields: format, args
func (_m *Logger) Printf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Printj provides a mock function with given fields: j
func (_m *Logger) Printj(j log.JSON) {
	_m.Called(j)
}

// SetHeader provides a mock function with given fields: h
func (_m *Logger) SetHeader(h string) {
	_m.Called(h)
}

// SetLevel provides a mock function with given fields: v
func (_m *Logger) SetLevel(v log.Lvl) {
	_m.Called(v)
}

// SetOutput provides a mock function with given fields: w
func (_m *Logger) SetOutput(w io.Writer) {
	_m.Called(w)
}

// SetPrefix provides a mock function with given fields: p
func (_m *Logger) SetPrefix(p string) {
	_m.Called(p)
}

// Warn provides a mock function with given fields: i
func (_m *Logger) Warn(i ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, i...)
	_m.Called(_ca...)
}

// Warnf provides a mock function with given fields: format, args
func (_m *Logger) Warnf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Warnj provides a mock function with given fields: j
func (_m *Logger) Warnj(j log.JSON) {
	_m.Called(j)
}
