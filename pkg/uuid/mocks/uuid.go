// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Uuid is an autogenerated mock type for the Uuid type
type Uuid struct {
	mock.Mock
}

// Generate provides a mock function with given fields:
func (_m *Uuid) Generate() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
