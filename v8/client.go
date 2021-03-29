// Package redismock is useful for unit testing applications that interact with
// Redis.
//
// The two constructors NewMock and NewNiceMock are explained in more detail on
// the functions themselves.
package redismock

import (
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/mock"
)

type ClientMock struct {
	mu sync.Mutex

	mock.Mock
	redis.Cmdable
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
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, call := range m.ExpectedCalls {
		if call.Method == method && call.Repeatability > -1 {
			return true
		}
	}

	return false
}

//
// Implmenent all the required mock.Mock functions but the sync.Mutex to make the library thread safe
//

func (m *ClientMock) Test(t mock.TestingT) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.Mock.Test(t)
}

func (m *ClientMock) On(methodName string, arguments ...interface{}) *mock.Call {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.Mock.On(methodName, arguments...)
}

func (m *ClientMock) MethodCalled(methodName string, arguments ...interface{}) mock.Arguments {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.Mock.MethodCalled(methodName, arguments...)
}

func (m *ClientMock) AssertExpectations(t mock.TestingT) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.Mock.AssertExpectations(t)
}

func (m *ClientMock) AssertNumberOfCalls(t mock.TestingT, methodName string, expectedCalls int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.Mock.AssertNumberOfCalls(t, methodName, expectedCalls)
}

func (m *ClientMock) AssertCalled(t mock.TestingT, methodName string, arguments ...interface{}) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.Mock.AssertCalled(t, methodName, arguments...)
}

func (m *ClientMock) AssertNotCalled(t mock.TestingT, methodName string, arguments ...interface{}) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.Mock.AssertNotCalled(t, methodName, arguments...)
}

func (m *ClientMock) IsMethodCallable(t mock.TestingT, methodName string, arguments ...interface{}) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.Mock.IsMethodCallable(t, methodName, arguments...)
}
