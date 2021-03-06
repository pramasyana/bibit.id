// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// yaml_read_handler_t is an autogenerated mock type for the yaml_read_handler_t type
type yaml_read_handler_t struct {
	mock.Mock
}

// Execute provides a mock function with given fields: parser, buffer
func (_m *yaml_read_handler_t) Execute(parser *yaml.yaml_parser_t, buffer []byte) (int, error) {
	ret := _m.Called(parser, buffer)

	var r0 int
	if rf, ok := ret.Get(0).(func(*yaml.yaml_parser_t, []byte) int); ok {
		r0 = rf(parser, buffer)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*yaml.yaml_parser_t, []byte) error); ok {
		r1 = rf(parser, buffer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
