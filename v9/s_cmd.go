package redismock

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func (m *ClientMock) SAdd(ctx context.Context, key string, members ...interface{}) *redis.IntCmd {
	if !m.hasStub("SAdd") {
		return m.client.SAdd(ctx, key, members...)
	}

	return m.Called(ctx, key, members).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) SCard(ctx context.Context, key string) *redis.IntCmd {
	if !m.hasStub("SCard") {
		return m.client.SCard(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) SDiff(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	if !m.hasStub("SDiff") {
		return m.client.SDiff(ctx, keys...)
	}

	return m.Called(ctx, keys).Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) SDiffStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	if !m.hasStub("SDiffStore") {
		return m.client.SDiffStore(ctx, destination, keys...)
	}

	return m.Called(ctx, destination, keys).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) SInter(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	if !m.hasStub("SInter") {
		return m.client.SInter(ctx, keys...)
	}

	return m.Called(ctx, keys).Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) SInterStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	if !m.hasStub("SInterStore") {
		return m.client.SInterStore(ctx, destination, keys...)
	}

	return m.Called(ctx, destination, keys).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) SIsMember(ctx context.Context, key string, member interface{}) *redis.BoolCmd {
	if !m.hasStub("SIsMember") {
		return m.client.SIsMember(ctx, key, member)
	}

	return m.Called(ctx, key, member).Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) SMembers(ctx context.Context, key string) *redis.StringSliceCmd {
	if !m.hasStub("SMembers") {
		return m.client.SMembers(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) SMembersMap(ctx context.Context, key string) *redis.StringStructMapCmd {
	if !m.hasStub("SMembersMap") {
		return m.client.SMembersMap(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.StringStructMapCmd)
}

func (m *ClientMock) SMove(ctx context.Context, source, destination string, member interface{}) *redis.BoolCmd {
	if !m.hasStub("SMove") {
		return m.client.SMove(ctx, source, destination, member)
	}

	return m.Called(ctx, source, destination, member).Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) SPop(ctx context.Context, key string) *redis.StringCmd {
	if !m.hasStub("SPop") {
		return m.client.SPop(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.StringCmd)
}

func (m *ClientMock) SPopN(ctx context.Context, key string, count int64) *redis.StringSliceCmd {
	if !m.hasStub("SPopN") {
		return m.client.SPopN(ctx, key, count)
	}

	return m.Called(ctx, key, count).Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) SRandMember(ctx context.Context, key string) *redis.StringCmd {
	if !m.hasStub("SRandMember") {
		return m.client.SRandMember(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.StringCmd)
}

func (m *ClientMock) SRandMemberN(ctx context.Context, key string, count int64) *redis.StringSliceCmd {
	if !m.hasStub("SRandMemberN") {
		return m.client.SRandMemberN(ctx, key, count)
	}

	return m.Called(ctx, key, count).Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) SRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd {
	if !m.hasStub("SRem") {
		return m.client.SRem(ctx, key, members...)
	}

	return m.Called(ctx, key, members).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) SUnion(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	if !m.hasStub("SUnion") {
		return m.client.SUnion(ctx, keys...)
	}

	return m.Called(ctx, keys).Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) SUnionStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	if !m.hasStub("SUnionStore") {
		return m.client.SUnionStore(ctx, destination, keys...)
	}

	return m.Called(ctx, destination, keys).Get(0).(*redis.IntCmd)
}
