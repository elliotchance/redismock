package redismock

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func (m *ClientMock) Watch(ctx context.Context, fn func(*redis.Tx) error, keys ...string) error {
	if !m.hasStub("Watch") {
		return m.client.Watch(ctx, fn, keys...)
	}

	args := m.Called()

	return args.Get(0).(error)
}
