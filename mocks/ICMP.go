// Code generated by mockery v2.30.17. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	icmp "github.com/tonobo/mtr/pkg/icmp"

	net "net"

	time "time"
)

// ICMP is an autogenerated mock type for the ICMP type
type ICMP struct {
	mock.Mock
}

// SendICMP provides a mock function with given fields: localAddr, dst, target, ttl, id, timeout, seq
func (_m *ICMP) SendICMP(localAddr string, dst net.Addr, target string, ttl int, id int, timeout time.Duration, seq int) (icmp.ICMPReturn, error) {
	ret := _m.Called(localAddr, dst, target, ttl, id, timeout, seq)

	var r0 icmp.ICMPReturn
	var r1 error
	if rf, ok := ret.Get(0).(func(string, net.Addr, string, int, int, time.Duration, int) (icmp.ICMPReturn, error)); ok {
		return rf(localAddr, dst, target, ttl, id, timeout, seq)
	}
	if rf, ok := ret.Get(0).(func(string, net.Addr, string, int, int, time.Duration, int) icmp.ICMPReturn); ok {
		r0 = rf(localAddr, dst, target, ttl, id, timeout, seq)
	} else {
		r0 = ret.Get(0).(icmp.ICMPReturn)
	}

	if rf, ok := ret.Get(1).(func(string, net.Addr, string, int, int, time.Duration, int) error); ok {
		r1 = rf(localAddr, dst, target, ttl, id, timeout, seq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewICMP creates a new instance of ICMP. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewICMP(t interface {
	mock.TestingT
	Cleanup(func())
}) *ICMP {
	mock := &ICMP{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
