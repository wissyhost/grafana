// Code generated by mockery v2.40.1. DO NOT EDIT.

package oauthtokentest

import (
	context "context"

	identity "github.com/grafana/grafana/pkg/services/auth/identity"
	datasources "github.com/grafana/grafana/pkg/services/datasources"

	login "github.com/grafana/grafana/pkg/services/login"

	mock "github.com/stretchr/testify/mock"

	oauth2 "golang.org/x/oauth2"
)

// MockService is an autogenerated mock type for the OAuthTokenService type
type MockService struct {
	mock.Mock
}

// GetCurrentOAuthToken provides a mock function with given fields: _a0, _a1
func (_m *MockService) GetCurrentOAuthToken(_a0 context.Context, _a1 identity.Requester) *oauth2.Token {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetCurrentOAuthToken")
	}

	var r0 *oauth2.Token
	if rf, ok := ret.Get(0).(func(context.Context, identity.Requester) *oauth2.Token); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oauth2.Token)
		}
	}

	return r0
}

// HasOAuthEntry provides a mock function with given fields: _a0, _a1
func (_m *MockService) HasOAuthEntry(_a0 context.Context, _a1 identity.Requester) (*login.UserAuth, bool, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for HasOAuthEntry")
	}

	var r0 *login.UserAuth
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, identity.Requester) (*login.UserAuth, bool, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, identity.Requester) *login.UserAuth); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*login.UserAuth)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, identity.Requester) bool); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(context.Context, identity.Requester) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// InvalidateOAuthTokens provides a mock function with given fields: _a0, _a1
func (_m *MockService) InvalidateOAuthTokens(_a0 context.Context, _a1 *login.UserAuth) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for InvalidateOAuthTokens")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *login.UserAuth) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsOAuthPassThruEnabled provides a mock function with given fields: _a0
func (_m *MockService) IsOAuthPassThruEnabled(_a0 *datasources.DataSource) bool {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for IsOAuthPassThruEnabled")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(*datasources.DataSource) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// TryTokenRefresh provides a mock function with given fields: _a0, _a1
func (_m *MockService) TryTokenRefresh(_a0 context.Context, _a1 identity.Requester) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for TryTokenRefresh")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, identity.Requester) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockService {
	mock := &MockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
