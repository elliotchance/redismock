package redismock

import "github.com/go-redis/redis"

func (m *ClientMock) XDel(stream string, ids ...string) *redis.IntCmd {
	if !m.hasStub("XDel") {
		return m.client.XDel(stream, ids...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) XAdd(a *redis.XAddArgs) *redis.StringCmd {
	if !m.hasStub("XAdd") {
		return m.client.XAdd(a)
	}

	return m.Called().Get(0).(*redis.StringCmd)
}

func (m *ClientMock) XLen(stream string) *redis.IntCmd {
	if !m.hasStub("XLen") {
		return m.client.XLen(stream)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) XRange(stream, start, stop string) *redis.XMessageSliceCmd {
	if !m.hasStub("XRange") {
		return m.client.XRange(stream, start, stop)
	}

	return m.Called().Get(0).(*redis.XMessageSliceCmd)
}

func (m *ClientMock) XRangeN(stream, start, stop string, count int64) *redis.XMessageSliceCmd {
	if !m.hasStub("XRangeN") {
		return m.client.XRangeN(stream, start, stop, count)
	}

	return m.Called().Get(0).(*redis.XMessageSliceCmd)
}

func (m *ClientMock) XRead(a *redis.XReadArgs) *redis.XStreamSliceCmd {
	if !m.hasStub("XRead") {
		return m.client.XRead(a)
	}

	return m.Called().Get(0).(*redis.XStreamSliceCmd)
}

func (m *ClientMock) XRevRange(stream string, start, stop string) *redis.XMessageSliceCmd {
	if !m.hasStub("XRevRange") {
		return m.client.XRevRange(stream, start, stop)
	}

	return m.Called().Get(0).(*redis.XMessageSliceCmd)
}

func (m *ClientMock) XRevRangeN(stream string, start, stop string, count int64) *redis.XMessageSliceCmd {
	if !m.hasStub("XRevRangeN") {
		return m.client.XRevRangeN(stream, start, stop, count)
	}

	return m.Called().Get(0).(*redis.XMessageSliceCmd)
}

func (m *ClientMock) XAck(stream, group string, ids ...string) *redis.IntCmd {
	if !m.hasStub("XAck") {
		return m.client.XAck(stream, group, ids...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) XClaim(a *redis.XClaimArgs) *redis.XMessageSliceCmd {
	if !m.hasStub("XClaim") {
		return m.client.XClaim(a)
	}

	return m.Called().Get(0).(*redis.XMessageSliceCmd)
}

func (m *ClientMock) XClaimJustID(a *redis.XClaimArgs) *redis.StringSliceCmd {
	if !m.hasStub("XClaimJustID") {
		return m.client.XClaimJustID(a)
	}

	return m.Called().Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) XGroupCreate(stream, group, start string) *redis.StatusCmd {
	if !m.hasStub("XGroupCreate") {
		return m.client.XGroupCreate(stream, group, start)
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) XGroupDelConsumer(stream, group, consumer string) *redis.IntCmd {
	if !m.hasStub("XGroupDelConsumer") {
		return m.client.XGroupDelConsumer(stream, group, consumer)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) XGroupDestroy(stream, group string) *redis.IntCmd {
	if !m.hasStub("XGroupDestroy") {
		return m.client.XGroupDestroy(stream, group)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) XGroupSetID(stream, group, start string) *redis.StatusCmd {
	if !m.hasStub("XGroupSetID") {
		return m.client.XGroupSetID(stream, group, start)
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) XPending(stream, group string) *redis.XPendingCmd {
	if !m.hasStub("XPending") {
		return m.client.XPending(stream, group)
	}

	return m.Called().Get(0).(*redis.XPendingCmd)
}

func (m *ClientMock) XPendingExt(a *redis.XPendingExtArgs) *redis.XPendingExtCmd {
	if !m.hasStub("XPendingExt") {
		return m.client.XPendingExt(a)
	}

	return m.Called().Get(0).(*redis.XPendingExtCmd)
}

func (m *ClientMock) XReadGroup(a *redis.XReadGroupArgs) *redis.XStreamSliceCmd {
	if !m.hasStub("XReadGroup") {
		return m.client.XReadGroup(a)
	}

	return m.Called().Get(0).(*redis.XStreamSliceCmd)
}

func (m *ClientMock) XReadStreams(streams ...string) *redis.XStreamSliceCmd {
	if !m.hasStub("XReadStreams") {
		return m.client.XReadStreams(streams...)
	}

	return m.Called().Get(0).(*redis.XStreamSliceCmd)
}

func (m *ClientMock) XTrim(key string, maxLen int64) *redis.IntCmd {
	if !m.hasStub("XTrim") {
		return m.client.XTrim(key, maxLen)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) XTrimApprox(key string, maxLen int64) *redis.IntCmd {
	if !m.hasStub("XTrimApprox") {
		return m.client.XTrimApprox(key, maxLen)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) XGroupCreateMkStream(stream, group, start string) *redis.StatusCmd {
	if !m.hasStub("XGroupCreateMkStream") {
		return m.client.XGroupCreateMkStream(stream, group, start)
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}
