package redismock

import "github.com/go-redis/redis"

func (m *ClientMock) HDel(key string, fields ...string) *redis.IntCmd {
	if !m.hasStub("HDel") {
		return m.client.HDel(key, fields...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) HExists(key, field string) *redis.BoolCmd {
	if !m.hasStub("HExists") {
		return m.client.HExists(key, field)
	}

	return m.Called().Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) HGet(key, field string) *redis.StringCmd {
	if !m.hasStub("HGet") {
		return m.client.HGet(key, field)
	}

	return m.Called().Get(0).(*redis.StringCmd)
}

func (m *ClientMock) HGetAll(key string) *redis.StringStringMapCmd {
	if !m.hasStub("HGetAll") {
		return m.client.HGetAll(key)
	}

	return m.Called().Get(0).(*redis.StringStringMapCmd)
}

func (m *ClientMock) HIncrBy(key, field string, incr int64) *redis.IntCmd {
	if !m.hasStub("HIncrBy") {
		return m.client.HIncrBy(key, field, incr)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) HIncrByFloat(key, field string, incr float64) *redis.FloatCmd {
	if !m.hasStub("HIncrByFloat") {
		return m.client.HIncrByFloat(key, field, incr)
	}

	return m.Called().Get(0).(*redis.FloatCmd)
}

func (m *ClientMock) HKeys(key string) *redis.StringSliceCmd {
	if !m.hasStub("HKeys") {
		return m.client.HKeys(key)
	}

	return m.Called().Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) HLen(key string) *redis.IntCmd {
	if !m.hasStub("HLen") {
		return m.client.HLen(key)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) HMGet(key string, fields ...string) *redis.SliceCmd {
	if !m.hasStub("HMGet") {
		return m.client.HMGet(key, fields...)
	}

	return m.Called().Get(0).(*redis.SliceCmd)
}

func (m *ClientMock) HMSet(key string, fields map[string]interface{}) *redis.StatusCmd {
	if !m.hasStub("HMSet") {
		return m.client.HMSet(key, fields)
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) HSet(key, field string, value interface{}) *redis.BoolCmd {
	if !m.hasStub("HSet") {
		return m.client.HSet(key, field, value)
	}

	return m.Called().Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) HSetNX(key, field string, value interface{}) *redis.BoolCmd {
	if !m.hasStub("HSetNX") {
		return m.client.HSetNX(key, field, value)
	}

	return m.Called().Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) HVals(key string) *redis.StringSliceCmd {
	if !m.hasStub("HVals") {
		return m.client.HVals(key)
	}

	return m.Called().Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) HScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	if !m.hasStub("HScan") {
		return m.client.HScan(key, cursor, match, count)
	}

	return m.Called().Get(0).(*redis.ScanCmd)
}
