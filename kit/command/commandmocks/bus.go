// Code generated by mockery v2.39.1. DO NOT EDIT.

package commandmocks

import (
	command "cabify-code-challenge/kit/command"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Bus is an autogenerated mock type for the Bus type
type Bus struct {
	mock.Mock
}

// Dispatch provides a mock function with given fields: _a0, _a1
func (_m *Bus) Dispatch(_a0 context.Context, _a1 command.Command) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Dispatch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, command.Command) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Register provides a mock function with given fields: _a0, _a1
func (_m *Bus) Register(_a0 command.Type, _a1 command.Handler) {
	_m.Called(_a0, _a1)
}

// NewBus creates a new instance of Bus. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBus(t interface {
	mock.TestingT
	Cleanup(func())
}) *Bus {
	mock := &Bus{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
