package redismock

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func (m *ClientMock) XDel(ctx context.Context, stream string, ids ...string) *redis.IntCmd {
	if !m.hasStub("XDel") {
		return m.client.XDel(ctx, stream, ids...)
	}

	return m.Called(ctx, stream, ids).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) XAdd(ctx context.Context, a *redis.XAddArgs) *redis.StringCmd {
	if !m.hasStub("XAdd") {
		return m.client.XAdd(ctx, a)
	}

	return m.Called(ctx, a).Get(0).(*redis.StringCmd)
}

func (m *ClientMock) XLen(ctx context.Context, stream string) *redis.IntCmd {
	if !m.hasStub("XLen") {
		return m.client.XLen(ctx, stream)
	}

	return m.Called(ctx, stream).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) XRange(ctx context.Context, stream, start, stop string) *redis.XMessageSliceCmd {
	if !m.hasStub("XRange") {
		return m.client.XRange(ctx, stream, start, stop)
	}

	return m.Called(ctx, stream, start, stop).Get(0).(*redis.XMessageSliceCmd)
}

func (m *ClientMock) XRangeN(ctx context.Context, stream, start, stop string, count int64) *redis.XMessageSliceCmd {
	if !m.hasStub("XRangeN") {
		return m.client.XRangeN(ctx, stream, start, stop, count)
	}

	return m.Called(ctx, stream, start, stop, count).Get(0).(*redis.XMessageSliceCmd)
}

func (m *ClientMock) XRead(ctx context.Context, a *redis.XReadArgs) *redis.XStreamSliceCmd {
	if !m.hasStub("XRead") {
		return m.client.XRead(ctx, a)
	}

	return m.Called(ctx, a).Get(0).(*redis.XStreamSliceCmd)
}

func (m *ClientMock) XRevRange(ctx context.Context, stream string, start, stop string) *redis.XMessageSliceCmd {
	if !m.hasStub("XRevRange") {
		return m.client.XRevRange(ctx, stream, start, stop)
	}

	return m.Called(ctx, stream, start, stop).Get(0).(*redis.XMessageSliceCmd)
}

func (m *ClientMock) XRevRangeN(ctx context.Context, stream string, start, stop string, count int64) *redis.XMessageSliceCmd {
	if !m.hasStub("XRevRangeN") {
		return m.client.XRevRangeN(ctx, stream, start, stop, count)
	}

	return m.Called(ctx, stream, start, stop, count).Get(0).(*redis.XMessageSliceCmd)
}

func (m *ClientMock) XAck(ctx context.Context, stream, group string, ids ...string) *redis.IntCmd {
	if !m.hasStub("XAck") {
		return m.client.XAck(ctx, stream, group, ids...)
	}

	return m.Called(ctx, stream, group, ids).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) XClaim(ctx context.Context, a *redis.XClaimArgs) *redis.XMessageSliceCmd {
	if !m.hasStub("XClaim") {
		return m.client.XClaim(ctx, a)
	}

	return m.Called(ctx, a).Get(0).(*redis.XMessageSliceCmd)
}

func (m *ClientMock) XClaimJustID(ctx context.Context, a *redis.XClaimArgs) *redis.StringSliceCmd {
	if !m.hasStub("XClaimJustID") {
		return m.client.XClaimJustID(ctx, a)
	}

	return m.Called(ctx, a).Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) XGroupCreate(ctx context.Context, stream, group, start string) *redis.StatusCmd {
	if !m.hasStub("XGroupCreate") {
		return m.client.XGroupCreate(ctx, stream, group, start)
	}

	return m.Called(ctx, stream, group, start).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) XGroupDelConsumer(ctx context.Context, stream, group, consumer string) *redis.IntCmd {
	if !m.hasStub("XGroupDelConsumer") {
		return m.client.XGroupDelConsumer(ctx, stream, group, consumer)
	}

	return m.Called(ctx, stream, group, consumer).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) XGroupDestroy(ctx context.Context, stream, group string) *redis.IntCmd {
	if !m.hasStub("XGroupDestroy") {
		return m.client.XGroupDestroy(ctx, stream, group)
	}

	return m.Called(ctx, stream, group).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) XGroupSetID(ctx context.Context, stream, group, start string) *redis.StatusCmd {
	if !m.hasStub("XGroupSetID") {
		return m.client.XGroupSetID(ctx, stream, group, start)
	}

	return m.Called(ctx, stream, group, start).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) XPending(ctx context.Context, stream, group string) *redis.XPendingCmd {
	if !m.hasStub("XPending") {
		return m.client.XPending(ctx, stream, group)
	}

	return m.Called(ctx, stream, group).Get(0).(*redis.XPendingCmd)
}

func (m *ClientMock) XPendingExt(ctx context.Context, a *redis.XPendingExtArgs) *redis.XPendingExtCmd {
	if !m.hasStub("XPendingExt") {
		return m.client.XPendingExt(ctx, a)
	}

	return m.Called(ctx, a).Get(0).(*redis.XPendingExtCmd)
}

func (m *ClientMock) XReadGroup(ctx context.Context, a *redis.XReadGroupArgs) *redis.XStreamSliceCmd {
	if !m.hasStub("XReadGroup") {
		return m.client.XReadGroup(ctx, a)
	}

	return m.Called(ctx, a).Get(0).(*redis.XStreamSliceCmd)
}

func (m *ClientMock) XReadStreams(ctx context.Context, streams ...string) *redis.XStreamSliceCmd {
	if !m.hasStub("XReadStreams") {
		return m.client.XReadStreams(ctx, streams...)
	}

	return m.Called(ctx, streams).Get(0).(*redis.XStreamSliceCmd)
}

func (m *ClientMock) XTrimMaxLen(ctx context.Context, key string, maxLen int64) *redis.IntCmd {
	if !m.hasStub("XTrimMaxLen") {
		return m.client.XTrimMaxLen(ctx, key, maxLen)
	}

	return m.Called(ctx, key, maxLen).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) XTrimMaxLenApprox(ctx context.Context, key string, maxLen int64, limit int64) *redis.IntCmd {
	if !m.hasStub("XTrimMaxLenApprox") {
		return m.client.XTrimMaxLenApprox(ctx, key, maxLen, limit)
	}

	return m.Called(ctx, key, maxLen).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) XGroupCreateMkStream(ctx context.Context, stream, group, start string) *redis.StatusCmd {
	if !m.hasStub("XGroupCreateMkStream") {
		return m.client.XGroupCreateMkStream(ctx, stream, group, start)
	}

	return m.Called(ctx, stream, group, start).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) XInfoGroups(ctx context.Context, key string) *redis.XInfoGroupsCmd {
	if !m.hasStub("XInfoGroups") {
		return m.client.XInfoGroups(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.XInfoGroupsCmd)
}

func (m *ClientMock) XInfoStream(ctx context.Context, key string) *redis.XInfoStreamCmd {
	if !m.hasStub("XInfoStream") {
		return m.client.XInfoStream(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.XInfoStreamCmd)
}
