package redismock

import "github.com/go-redis/redis"

func (m *ClientMock) SAdd(key string, members ...interface{}) *redis.IntCmd {
	if !m.hasStub("SAdd") {
		return m.client.SAdd(key, members...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) SCard(key string) *redis.IntCmd {
	if !m.hasStub("SCard") {
		return m.client.SCard(key)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) SDiff(keys ...string) *redis.StringSliceCmd {
	if !m.hasStub("SDiff") {
		return m.client.SDiff()
	}

	return m.Called().Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) SDiffStore(destination string, keys ...string) *redis.IntCmd {
	if !m.hasStub("SDiffStore") {
		return m.client.SDiffStore(destination, keys...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) SInter(keys ...string) *redis.StringSliceCmd {
	if !m.hasStub("SInter") {
		return m.client.SInter()
	}

	return m.Called().Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) SInterStore(destination string, keys ...string) *redis.IntCmd {
	if !m.hasStub("SInterStore") {
		return m.client.SInterStore(destination, keys...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) SIsMember(key string, member interface{}) *redis.BoolCmd {
	if !m.hasStub("SIsMember") {
		return m.client.SIsMember(key, member)
	}

	return m.Called().Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) SMembers(key string) *redis.StringSliceCmd {
	if !m.hasStub("SMembers") {
		return m.client.SMembers(key)
	}

	return m.Called().Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) SMembersMap(key string) *redis.StringStructMapCmd {
	if !m.hasStub("SMembersMap") {
		return m.client.SMembersMap(key)
	}

	return m.Called().Get(0).(*redis.StringStructMapCmd)
}

func (m *ClientMock) SMove(source, destination string, member interface{}) *redis.BoolCmd {
	if !m.hasStub("SMove") {
		return m.client.SMove(source, destination, member)
	}

	return m.Called().Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) SPop(key string) *redis.StringCmd {
	if !m.hasStub("SPop") {
		return m.client.SPop(key)
	}

	return m.Called().Get(0).(*redis.StringCmd)
}

func (m *ClientMock) SPopN(key string, count int64) *redis.StringSliceCmd {
	if !m.hasStub("SPopN") {
		return m.client.SPopN(key, count)
	}

	return m.Called().Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) SRandMember(key string) *redis.StringCmd {
	if !m.hasStub("SRandMember") {
		return m.client.SRandMember(key)
	}

	return m.Called().Get(0).(*redis.StringCmd)
}

func (m *ClientMock) SRandMemberN(key string, count int64) *redis.StringSliceCmd {
	if !m.hasStub("SRandMemberN") {
		return m.client.SRandMemberN(key, count)
	}

	return m.Called().Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) SRem(key string, members ...interface{}) *redis.IntCmd {
	if !m.hasStub("SRem") {
		return m.client.SRem(key, members...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) SUnion(keys ...string) *redis.StringSliceCmd {
	if !m.hasStub("SUnion") {
		return m.client.SUnion(keys...)
	}

	return m.Called().Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) SUnionStore(destination string, keys ...string) *redis.IntCmd {
	if !m.hasStub("SUnionStore") {
		return m.client.SUnionStore(destination, keys...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}
