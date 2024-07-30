package redismock

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func (m *ClientMock) BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	if !m.hasStub("BZPopMax") {
		return m.client.BZPopMax(ctx, timeout, keys...)
	}

	return m.Called(ctx, timeout, keys).Get(0).(*redis.ZWithKeyCmd)
}

func (m *ClientMock) BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	if !m.hasStub("BZPopMin") {
		return m.client.BZPopMin(ctx, timeout, keys...)
	}

	return m.Called(ctx, timeout, keys).Get(0).(*redis.ZWithKeyCmd)
}

func (m *ClientMock) ZPopMax(ctx context.Context, key string, count ...int64) *redis.ZSliceCmd {
	if !m.hasStub("ZPopMax") {
		return m.client.ZPopMax(ctx, key, count...)
	}

	return m.Called(ctx, key, count).Get(0).(*redis.ZSliceCmd)
}

func (m *ClientMock) ZPopMin(ctx context.Context, key string, count ...int64) *redis.ZSliceCmd {
	if !m.hasStub("ZPopMin") {
		return m.client.ZPopMin(ctx, key, count...)
	}

	return m.Called(ctx, key, count).Get(0).(*redis.ZSliceCmd)
}

func (m *ClientMock) ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	if !m.hasStub("ZScan") {
		return m.client.ZScan(ctx, key, cursor, match, count)
	}

	return m.Called(ctx, key, cursor, match, count).Get(0).(*redis.ScanCmd)
}

func (m *ClientMock) ZAddArgs(ctx context.Context, key string, args redis.ZAddArgs) *redis.IntCmd {
	if !m.hasStub("ZAddArgs") {
		return m.client.ZAddArgs(ctx, key, args)
	}

	return m.Called(ctx, key, args).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZAddArgsIncr(ctx context.Context, key string, args redis.ZAddArgs) *redis.FloatCmd {
	if !m.hasStub("ZAddArgsIncr") {
		return m.client.ZAddArgsIncr(ctx, key, args)
	}

	return m.Called(ctx, key, args).Get(0).(*redis.FloatCmd)
}

func (m *ClientMock) ZCard(ctx context.Context, key string) *redis.IntCmd {
	if !m.hasStub("ZCard") {
		return m.client.ZCard(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZCount(ctx context.Context, key, min, max string) *redis.IntCmd {
	if !m.hasStub("ZCount") {
		return m.client.ZCount(ctx, key, min, max)
	}

	return m.Called(ctx, key, min, max).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZLexCount(ctx context.Context, key, min, max string) *redis.IntCmd {
	if !m.hasStub("ZLexCount") {
		return m.client.ZLexCount(ctx, key, min, max)
	}

	return m.Called(ctx, key, min, max).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZIncrBy(ctx context.Context, key string, increment float64, member string) *redis.FloatCmd {
	if !m.hasStub("ZIncrBy") {
		return m.client.ZIncrBy(ctx, key, increment, member)
	}

	return m.Called(ctx, key, increment, member).Get(0).(*redis.FloatCmd)
}

func (m *ClientMock) ZInterStore(ctx context.Context, destination string, store *redis.ZStore) *redis.IntCmd {
	if !m.hasStub("ZInterStore") {
		return m.client.ZInterStore(ctx, destination, store)
	}

	return m.Called(ctx, destination, store).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	if !m.hasStub("ZRange") {
		return m.client.ZRange(ctx, key, start, stop)
	}

	return m.Called(ctx, key, start, stop).Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) ZRangeWithScores(ctx context.Context, key string, start, stop int64) *redis.ZSliceCmd {
	if !m.hasStub("ZRangeWithScores") {
		return m.client.ZRangeWithScores(ctx, key, start, stop)
	}

	return m.Called(ctx, key, start, stop).Get(0).(*redis.ZSliceCmd)
}

func (m *ClientMock) ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	if !m.hasStub("ZRangeByScore") {
		return m.client.ZRangeByScore(ctx, key, opt)
	}

	return m.Called(ctx).Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) ZRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	if !m.hasStub("ZRangeByLex") {
		return m.client.ZRangeByLex(ctx, key, opt)
	}

	return m.Called(ctx, key, opt).Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) ZRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd {
	if !m.hasStub("ZRangeByScoreWithScores") {
		return m.client.ZRangeByScoreWithScores(ctx, key, opt)
	}

	return m.Called(ctx, key, opt).Get(0).(*redis.ZSliceCmd)
}

func (m *ClientMock) ZRank(ctx context.Context, key, member string) *redis.IntCmd {
	if !m.hasStub("ZRank") {
		return m.client.ZRank(ctx, key, member)
	}

	return m.Called(ctx, key, member).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd {
	if !m.hasStub("ZRem") {
		return m.client.ZRem(ctx, key, members...)
	}

	return m.Called(ctx, key, members).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *redis.IntCmd {
	if !m.hasStub("ZRemRangeByRank") {
		return m.client.ZRemRangeByRank(ctx, key, start, stop)
	}

	return m.Called(ctx, key, start, stop).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZRemRangeByScore(ctx context.Context, key, min, max string) *redis.IntCmd {
	if !m.hasStub("ZRemRangeByScore") {
		return m.client.ZRemRangeByScore(ctx, key, min, max)
	}

	return m.Called(ctx, key, min, max).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZRemRangeByLex(ctx context.Context, key, min, max string) *redis.IntCmd {
	if !m.hasStub("ZRemRangeByLex") {
		return m.client.ZRemRangeByLex(ctx, key, min, max)
	}

	return m.Called(ctx, key, min, max).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZRevRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	if !m.hasStub("ZRevRange") {
		return m.client.ZRevRange(ctx, key, start, stop)
	}

	return m.Called(ctx, key, start, stop).Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *redis.ZSliceCmd {
	if !m.hasStub("ZRevRangeWithScores") {
		return m.client.ZRevRangeWithScores(ctx, key, start, stop)
	}

	return m.Called(ctx, key, start, stop).Get(0).(*redis.ZSliceCmd)
}

func (m *ClientMock) ZRevRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	if !m.hasStub("ZRevRangeByScore") {
		return m.client.ZRevRangeByScore(ctx, key, opt)
	}

	return m.Called(ctx, key, opt).Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) ZRevRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	if !m.hasStub("ZRevRangeByLex") {
		return m.client.ZRevRangeByLex(ctx, key, opt)
	}

	return m.Called(ctx, key, opt).Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd {
	if !m.hasStub("ZRevRangeByScoreWithScores") {
		return m.client.ZRevRangeByScoreWithScores(ctx, key, opt)
	}

	return m.Called(ctx, key, opt).Get(0).(*redis.ZSliceCmd)
}

func (m *ClientMock) ZRevRank(ctx context.Context, key, member string) *redis.IntCmd {
	if !m.hasStub("ZRevRank") {
		return m.client.ZRevRank(ctx, key, member)
	}

	return m.Called(ctx, key, member).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZScore(ctx context.Context, key, member string) *redis.FloatCmd {
	if !m.hasStub("ZScore") {
		return m.client.ZScore(ctx, key, member)
	}

	return m.Called(ctx, key, member).Get(0).(*redis.FloatCmd)
}

func (m *ClientMock) ZUnionStore(ctx context.Context, dest string, store *redis.ZStore) *redis.IntCmd {
	if !m.hasStub("ZUnionStore") {
		return m.client.ZUnionStore(ctx, dest, store)
	}

	return m.Called(ctx, dest, store).Get(0).(*redis.IntCmd)
}
