// Package redismock is useful for unit testing applications that interact with
// Redis.
//
// The two constructors NewMock and NewNiceMock are explained in more detail on
// the functions themselves.
package redismock

import (
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/mock"
)

type ClientMock struct {
	mu sync.Mutex

	mock.Mock
	redis.Cmdable
	client redis.UniversalClient
	Ctx    context.Context
}

func (m *ClientMock) Context() context.Context {
	return m.Ctx
}

func (m *ClientMock) AddHook(hook redis.Hook) {
	m.client.AddHook(hook)
}

func (m *ClientMock) Do(ctx context.Context, args ...interface{}) *redis.Cmd {
	if !m.hasStub("Do") {
		return m.client.Do(ctx, args)
	}
	return m.Called().Get(0).(*redis.Cmd)
}

func (m *ClientMock) Process(ctx context.Context, cmd redis.Cmder) error {
	if !m.hasStub("Process") {
		return m.client.Process(ctx, cmd)
	}
	args := m.Called()
	return args.Get(0).(error)
}

func (m *ClientMock) Subscribe(ctx context.Context, channels ...string) *redis.PubSub {
	if !m.hasStub("Subscribe") {
		return m.client.Subscribe(ctx, channels...)
	}
	return m.Called().Get(0).(*redis.PubSub)
}

func (m *ClientMock) PSubscribe(ctx context.Context, channels ...string) *redis.PubSub {

	if !m.hasStub("PSubscribe") {
		return m.client.PSubscribe(ctx, channels...)
	}
	return m.Called().Get(0).(*redis.PubSub)
}

func (m *ClientMock) Close() error {
	if !m.hasStub("Close") {
		return m.client.Close()
	}
	args := m.Called()
	return args.Get(0).(error)
}

func (m *ClientMock) PoolStats() *redis.PoolStats {
	if !m.hasStub("PoolStats") {
		return m.client.PoolStats()
	}
	return m.Called().Get(0).(*redis.PoolStats)
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
func NewNiceMock(client redis.UniversalClient) *ClientMock {
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
