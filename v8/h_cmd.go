package redismock

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func (m *ClientMock) HDel(ctx context.Context, key string, fields ...string) *redis.IntCmd {
	if !m.hasStub("HDel") {
		return m.client.HDel(ctx, key, fields...)
	}

	return m.Called(ctx, key, fields).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) HExists(ctx context.Context, key, field string) *redis.BoolCmd {
	if !m.hasStub("HExists") {
		return m.client.HExists(ctx, key, field)
	}

	return m.Called(ctx, key, field).Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) HGet(ctx context.Context, key, field string) *redis.StringCmd {
	if !m.hasStub("HGet") {
		return m.client.HGet(ctx, key, field)
	}

	return m.Called(ctx, key, field).Get(0).(*redis.StringCmd)
}

func (m *ClientMock) HGetAll(ctx context.Context, key string) *redis.StringStringMapCmd {
	if !m.hasStub("HGetAll") {
		return m.client.HGetAll(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.StringStringMapCmd)
}

func (m *ClientMock) HIncrBy(ctx context.Context, key, field string, incr int64) *redis.IntCmd {
	if !m.hasStub("HIncrBy") {
		return m.client.HIncrBy(ctx, key, field, incr)
	}

	return m.Called(ctx, key, field, incr).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) HIncrByFloat(ctx context.Context, key, field string, incr float64) *redis.FloatCmd {
	if !m.hasStub("HIncrByFloat") {
		return m.client.HIncrByFloat(ctx, key, field, incr)
	}

	return m.Called(ctx, key, field, incr).Get(0).(*redis.FloatCmd)
}

func (m *ClientMock) HKeys(ctx context.Context, key string) *redis.StringSliceCmd {
	if !m.hasStub("HKeys") {
		return m.client.HKeys(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) HLen(ctx context.Context, key string) *redis.IntCmd {
	if !m.hasStub("HLen") {
		return m.client.HLen(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) HMGet(ctx context.Context, key string, fields ...string) *redis.SliceCmd {
	if !m.hasStub("HMGet") {
		return m.client.HMGet(ctx, key, fields...)
	}

	return m.Called(ctx, key, fields).Get(0).(*redis.SliceCmd)
}

func (m *ClientMock) HMSet(ctx context.Context, key string, fields ...interface{}) *redis.BoolCmd {
	if !m.hasStub("HMSet") {
		return m.client.HMSet(ctx, key, fields)
	}

	return m.Called(ctx, key, fields).Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	if !m.hasStub("HSet") {
		return m.client.HSet(ctx, key, values)
	}

	return m.Called(ctx, key, values).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) HSetNX(ctx context.Context, key, field string, value interface{}) *redis.BoolCmd {
	if !m.hasStub("HSetNX") {
		return m.client.HSetNX(ctx, key, field, value)
	}

	return m.Called(ctx, key, field, value).Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) HVals(ctx context.Context, key string) *redis.StringSliceCmd {
	if !m.hasStub("HVals") {
		return m.client.HVals(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	if !m.hasStub("HScan") {
		return m.client.HScan(ctx, key, cursor, match, count)
	}

	return m.Called(ctx, key, cursor, match, count).Get(0).(*redis.ScanCmd)
}
