// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// MSIConvertable is an autogenerated mock type for the MSIConvertable type
type MSIConvertable struct {
	mock.Mock
}

// MSI provides a mock function with given fields:
func (_m *MSIConvertable) MSI() map[string]interface{} {
	ret := _m.Called()

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func() map[string]interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}
