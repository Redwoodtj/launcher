// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"
	net "net"

	mock "github.com/stretchr/testify/mock"
)

// IpLookuper is an autogenerated mock type for the ipLookuper type
type IpLookuper struct {
	mock.Mock
}

// LookupIP provides a mock function with given fields: ctx, network, host
func (_m *IpLookuper) LookupIP(ctx context.Context, network string, host string) ([]net.IP, error) {
	ret := _m.Called(ctx, network, host)

	var r0 []net.IP
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []net.IP); ok {
		r0 = rf(ctx, network, host)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]net.IP)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, network, host)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
