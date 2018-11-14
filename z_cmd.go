package redismock

import (
	"time"

	"github.com/go-redis/redis"
)

func (m *ClientMock) BZPopMax(timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	if !m.hasStub("BZPopMax") {
		return m.client.BZPopMax(timeout, keys...)
	}

	return m.Called().Get(0).(*redis.ZWithKeyCmd)
}

func (m *ClientMock) BZPopMin(timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	if !m.hasStub("BZPopMin") {
		return m.client.BZPopMin(timeout, keys...)
	}

	return m.Called().Get(0).(*redis.ZWithKeyCmd)
}

func (m *ClientMock) ZPopMax(key string, count ...int64) *redis.ZSliceCmd {
	if !m.hasStub("ZPopMax") {
		return m.client.ZPopMax(key, count...)
	}

	return m.Called().Get(0).(*redis.ZSliceCmd)
}

func (m *ClientMock) ZPopMin(key string, count ...int64) *redis.ZSliceCmd {
	if !m.hasStub("ZPopMin") {
		return m.client.ZPopMin(key, count...)
	}

	return m.Called().Get(0).(*redis.ZSliceCmd)
}

func (m *ClientMock) ZScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	if !m.hasStub("ZScan") {
		return m.client.ZScan(key, cursor, match, count)
	}

	return m.Called().Get(0).(*redis.ScanCmd)
}

func (m *ClientMock) ZAdd(key string, members ...redis.Z) *redis.IntCmd {
	if !m.hasStub("ZAdd") {
		return m.client.ZAdd(key, members...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZAddNX(key string, members ...redis.Z) *redis.IntCmd {
	if !m.hasStub("ZAddNX") {
		return m.client.ZAddNX(key, members...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZAddXX(key string, members ...redis.Z) *redis.IntCmd {
	if !m.hasStub("ZAddXX") {
		return m.client.ZAddXX(key, members...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZAddCh(key string, members ...redis.Z) *redis.IntCmd {
	if !m.hasStub("ZAddCh") {
		return m.client.ZAddCh(key, members...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZAddNXCh(key string, members ...redis.Z) *redis.IntCmd {
	if !m.hasStub("ZAddNXCh") {
		return m.client.ZAddNXCh(key, members...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZAddXXCh(key string, members ...redis.Z) *redis.IntCmd {
	if !m.hasStub("ZAddXXCh") {
		return m.client.ZAddXXCh(key, members...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZIncr(key string, member redis.Z) *redis.FloatCmd {
	if !m.hasStub("ZIncr") {
		return m.client.ZIncr(key, member)
	}

	return m.Called().Get(0).(*redis.FloatCmd)
}

func (m *ClientMock) ZIncrNX(key string, member redis.Z) *redis.FloatCmd {
	if !m.hasStub("ZIncrNX") {
		return m.client.ZIncrNX(key, member)
	}

	return m.Called().Get(0).(*redis.FloatCmd)
}

func (m *ClientMock) ZIncrXX(key string, member redis.Z) *redis.FloatCmd {
	if !m.hasStub("ZIncrXX") {
		return m.client.ZIncrXX(key, member)
	}

	return m.Called().Get(0).(*redis.FloatCmd)
}

func (m *ClientMock) ZCard(key string) *redis.IntCmd {
	if !m.hasStub("ZCard") {
		return m.client.ZCard(key)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZCount(key, min, max string) *redis.IntCmd {
	if !m.hasStub("ZCount") {
		return m.client.ZCount(key, min, max)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZLexCount(key, min, max string) *redis.IntCmd {
	if !m.hasStub("ZLexCount") {
		return m.client.ZLexCount(key, min, max)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZIncrBy(key string, increment float64, member string) *redis.FloatCmd {
	if !m.hasStub("ZIncrBy") {
		return m.client.ZIncrBy(key, increment, member)
	}

	return m.Called().Get(0).(*redis.FloatCmd)
}

func (m *ClientMock) ZInterStore(destination string, store redis.ZStore, keys ...string) *redis.IntCmd {
	if !m.hasStub("ZInterStore") {
		return m.client.ZInterStore(destination, store, keys...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZRange(key string, start, stop int64) *redis.StringSliceCmd {
	if !m.hasStub("ZRange") {
		return m.client.ZRange(key, start, stop)
	}

	return m.Called().Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) ZRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd {
	if !m.hasStub("ZRangeWithScores") {
		return m.client.ZRangeWithScores(key, start, stop)
	}

	return m.Called().Get(0).(*redis.ZSliceCmd)
}

func (m *ClientMock) ZRangeByScore(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	if !m.hasStub("ZRangeByScore") {
		return m.client.ZRangeByScore(key, opt)
	}

	return m.Called().Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) ZRangeByLex(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	if !m.hasStub("ZRangeByLex") {
		return m.client.ZRangeByLex(key, opt)
	}

	return m.Called().Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) ZRangeByScoreWithScores(key string, opt redis.ZRangeBy) *redis.ZSliceCmd {
	if !m.hasStub("ZRangeByScoreWithScores") {
		return m.client.ZRangeByScoreWithScores(key, opt)
	}

	return m.Called().Get(0).(*redis.ZSliceCmd)
}

func (m *ClientMock) ZRank(key, member string) *redis.IntCmd {
	if !m.hasStub("ZRank") {
		return m.client.ZRank(key, member)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZRem(key string, members ...interface{}) *redis.IntCmd {
	if !m.hasStub("ZRem") {
		return m.client.ZRem(key, members...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZRemRangeByRank(key string, start, stop int64) *redis.IntCmd {
	if !m.hasStub("ZRemRangeByRank") {
		return m.client.ZRemRangeByRank(key, start, stop)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZRemRangeByScore(key, min, max string) *redis.IntCmd {
	if !m.hasStub("ZRemRangeByScore") {
		return m.client.ZRemRangeByScore(key, min, max)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZRemRangeByLex(key, min, max string) *redis.IntCmd {
	if !m.hasStub("ZRemRangeByLex") {
		return m.client.ZRemRangeByLex(key, min, max)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZRevRange(key string, start, stop int64) *redis.StringSliceCmd {
	if !m.hasStub("ZRevRange") {
		return m.client.ZRevRange(key, start, stop)
	}

	return m.Called().Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) ZRevRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd {
	if !m.hasStub("ZRevRangeWithScores") {
		return m.client.ZRevRangeWithScores(key, start, stop)
	}

	return m.Called().Get(0).(*redis.ZSliceCmd)
}

func (m *ClientMock) ZRevRangeByScore(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	if !m.hasStub("ZRevRangeByScore") {
		return m.client.ZRevRangeByScore(key, opt)
	}

	return m.Called().Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) ZRevRangeByLex(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	if !m.hasStub("ZRevRangeByLex") {
		return m.client.ZRevRangeByLex(key, opt)
	}

	return m.Called().Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) ZRevRangeByScoreWithScores(key string, opt redis.ZRangeBy) *redis.ZSliceCmd {
	if !m.hasStub("ZRevRangeByScoreWithScores") {
		return m.client.ZRevRangeByScoreWithScores(key, opt)
	}

	return m.Called().Get(0).(*redis.ZSliceCmd)
}

func (m *ClientMock) ZRevRank(key, member string) *redis.IntCmd {
	if !m.hasStub("ZRevRank") {
		return m.client.ZRevRank(key, member)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZScore(key, member string) *redis.FloatCmd {
	if !m.hasStub("ZScore") {
		return m.client.ZScore(key, member)
	}

	return m.Called().Get(0).(*redis.FloatCmd)
}

func (m *ClientMock) ZUnionStore(dest string, store redis.ZStore, keys ...string) *redis.IntCmd {
	if !m.hasStub("ZUnionStore") {
		return m.client.ZUnionStore(dest, store, keys...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}
