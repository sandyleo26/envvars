// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import envvars "github.com/flemay/envvars/pkg/envvars"
import mock "github.com/stretchr/testify/mock"

// EnvfileWriter is an autogenerated mock type for the EnvfileWriter type
type EnvfileWriter struct {
	mock.Mock
}

// Write provides a mock function with given fields: c
func (_m *EnvfileWriter) Write(c envvars.EnvvarCollection) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(envvars.EnvvarCollection) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
