package redismock

import "github.com/go-redis/redis/v7"

func (m *ClientMock) Watch(fn func(*redis.Tx) error, keys ...string) error {
	if !m.hasStub("Watch") {
		return m.client.Watch(fn, keys...)
	}

	args := m.Called()

	return args.Get(0).(error)
}
