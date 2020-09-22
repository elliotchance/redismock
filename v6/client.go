// Package redismock is useful for unit testing applications that interact with
// Redis.
//
// The two constructors NewMock and NewNiceMock are explained in more detail on
// the functions themselves.
package redismock

import (
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/mock"
)

type ClientMock struct {
	mock.Mock
	client *redis.Client
}

// NewMock creates a hollow mock. You will need to stub all commands that you
// will need to interact with.
//
// This is most useful when you want to strictly test all Redis interactions.
// See NewNiceMock().
func NewMock() *ClientMock {
	return new(ClientMock)
}

// NewNiceMock will create a mock that falls back to the real client in cases
// where a command has not been stubbed.
//
// This is most useful when you want a real Redis instance, but you need to stub
// off certain commands or behaviors. See NewMock().
func NewNiceMock(client *redis.Client) *ClientMock {
	m := NewMock()
	m.client = client

	return m
}

// hasStub check if a method has an expectation on it. This allows the calling
// method to determine if the method call should be passed through to the real
// redis client.
func (m *ClientMock) hasStub(method string) bool {
	for _, call := range m.ExpectedCalls {
		if call.Method == method && call.Repeatability > -1 {
			return true
		}
	}

	return false
}
