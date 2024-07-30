package redismock

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func (m *ClientMock) Pipeline() redis.Pipeliner {
	if !m.hasStub("Pipeline") {
		return m.client.Pipeline()
	}

	return m.Called().Get(0).(redis.Pipeliner)
}

func (m *ClientMock) Pipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	if !m.hasStub("Pipelined") {
		return m.client.Pipelined(ctx, fn)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).([]redis.Cmder), mockArgs.Get(1).(error)
}

func (m *ClientMock) TxPipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	if !m.hasStub("TxPipelined") {
		return m.client.TxPipelined(ctx, fn)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).([]redis.Cmder), mockArgs.Get(1).(error)
}

func (m *ClientMock) TxPipeline() redis.Pipeliner {
	if !m.hasStub("TxPipeline") {
		return m.client.TxPipeline()
	}

	return m.Called().Get(0).(redis.Pipeliner)
}

func (m *ClientMock) ClientGetName(ctx context.Context) *redis.StringCmd {
	if !m.hasStub("ClientGetName") {
		return m.client.ClientGetName(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StringCmd)
}

func (m *ClientMock) Echo(ctx context.Context, message interface{}) *redis.StringCmd {
	if !m.hasStub("Echo") {
		return m.client.Echo(ctx, message)
	}

	return m.Called(ctx, message).Get(0).(*redis.StringCmd)
}

func (m *ClientMock) Ping(ctx context.Context) *redis.StatusCmd {
	if !m.hasStub("Ping") {
		return m.client.Ping(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) Quit(ctx context.Context) *redis.StatusCmd {
	if !m.hasStub("Quit") {
		return m.client.Quit(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	if !m.hasStub("Del") {
		return m.client.Del(ctx, keys...)
	}

	return m.Called(ctx, keys).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) Unlink(ctx context.Context, keys ...string) *redis.IntCmd {
	if !m.hasStub("Unlink") {
		return m.client.Unlink(ctx, keys...)
	}

	return m.Called(ctx, keys).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) Dump(ctx context.Context, key string) *redis.StringCmd {
	if !m.hasStub("Dump") {
		return m.client.Dump(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.StringCmd)
}

func (m *ClientMock) Exists(ctx context.Context, keys ...string) *redis.IntCmd {
	if !m.hasStub("Exists") {
		return m.client.Exists(ctx, keys...)
	}

	return m.Called(ctx, keys).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	if !m.hasStub("Expire") {
		return m.client.Expire(ctx, key, expiration)
	}

	return m.Called(ctx, key, expiration).Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) ExpireAt(ctx context.Context, key string, tm time.Time) *redis.BoolCmd {
	if !m.hasStub("ExpireAt") {
		return m.client.ExpireAt(ctx, key, tm)
	}

	return m.Called(ctx).Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) Keys(ctx context.Context, pattern string) *redis.StringSliceCmd {
	if !m.hasStub("Keys") {
		return m.client.Keys(ctx, pattern)
	}

	return m.Called(ctx, pattern).Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) Migrate(ctx context.Context, host, port, key string, db int, timeout time.Duration) *redis.StatusCmd {
	if !m.hasStub("Migrate") {
		return m.client.Migrate(ctx, host, port, key, db, timeout)
	}

	return m.Called(ctx, host, port, key, db, timeout).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) Move(ctx context.Context, key string, db int) *redis.BoolCmd {
	if !m.hasStub("Move") {
		return m.client.Move(ctx, key, db)
	}

	return m.Called(ctx, key, db).Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) ObjectRefCount(ctx context.Context, key string) *redis.IntCmd {
	if !m.hasStub("ObjectRefCount") {
		return m.client.ObjectRefCount(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ObjectEncoding(ctx context.Context, key string) *redis.StringCmd {
	if !m.hasStub("ObjectEncoding") {
		return m.client.ObjectEncoding(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.StringCmd)
}

func (m *ClientMock) ObjectIdleTime(ctx context.Context, key string) *redis.DurationCmd {
	if !m.hasStub("ObjectIdleTime") {
		return m.client.ObjectIdleTime(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.DurationCmd)
}

func (m *ClientMock) Persist(ctx context.Context, key string) *redis.BoolCmd {
	if !m.hasStub("Persist") {
		return m.client.Persist(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) PExpire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	if !m.hasStub("PExpire") {
		return m.client.PExpire(ctx, key, expiration)
	}

	return m.Called(ctx, key, expiration).Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) PExpireAt(ctx context.Context, key string, tm time.Time) *redis.BoolCmd {
	if !m.hasStub("PExpireAt") {
		return m.client.PExpireAt(ctx, key, tm)
	}

	return m.Called(ctx, key, tm).Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) PTTL(ctx context.Context, key string) *redis.DurationCmd {
	if !m.hasStub("PTTL") {
		return m.client.PTTL(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.DurationCmd)
}

func (m *ClientMock) RandomKey(ctx context.Context) *redis.StringCmd {
	if !m.hasStub("RandomKey") {
		return m.client.RandomKey(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StringCmd)
}

func (m *ClientMock) Rename(ctx context.Context, key, newkey string) *redis.StatusCmd {
	if !m.hasStub("Rename") {
		return m.client.Rename(ctx, key, newkey)
	}

	return m.Called(ctx, key, newkey).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) RenameNX(ctx context.Context, key, newkey string) *redis.BoolCmd {
	if !m.hasStub("RenameNX") {
		return m.client.RenameNX(ctx, key, newkey)
	}

	return m.Called(ctx, key, newkey).Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) Restore(ctx context.Context, key string, ttl time.Duration, value string) *redis.StatusCmd {
	if !m.hasStub("Restore") {
		return m.client.Restore(ctx, key, ttl, value)
	}

	return m.Called(ctx, key, ttl, value).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) *redis.StatusCmd {
	if !m.hasStub("RestoreReplace") {
		return m.client.RestoreReplace(ctx, key, ttl, value)
	}

	return m.Called(ctx, key, ttl, value).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) Sort(ctx context.Context, key string, sort *redis.Sort) *redis.StringSliceCmd {
	if !m.hasStub("Sort") {
		return m.client.Sort(ctx, key, sort)
	}

	return m.Called(ctx, key, sort).Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) SortStore(ctx context.Context, key, store string, sort *redis.Sort) *redis.IntCmd {
	if !m.hasStub("SortStore") {
		return m.client.SortStore(ctx, key, store, sort)
	}

	return m.Called(ctx, key, store, sort).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) SortInterfaces(ctx context.Context, key string, sort *redis.Sort) *redis.SliceCmd {
	if !m.hasStub("SortInterfaces") {
		return m.client.SortInterfaces(ctx, key, sort)
	}

	return m.Called(ctx, key, sort).Get(0).(*redis.SliceCmd)
}

func (m *ClientMock) Touch(ctx context.Context, keys ...string) *redis.IntCmd {
	if !m.hasStub("Touch") {
		return m.client.Touch(ctx, keys...)
	}

	return m.Called(ctx, keys).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) TTL(ctx context.Context, key string) *redis.DurationCmd {
	if !m.hasStub("TTL") {
		return m.client.TTL(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.DurationCmd)
}

func (m *ClientMock) Type(ctx context.Context, key string) *redis.StatusCmd {
	if !m.hasStub("Type") {
		return m.client.Type(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd {
	if !m.hasStub("Scan") {
		return m.client.Scan(ctx, cursor, match, count)
	}

	return m.Called(ctx, cursor, match, count).Get(0).(*redis.ScanCmd)
}

func (m *ClientMock) SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	if !m.hasStub("SScan") {
		return m.client.SScan(ctx, key, cursor, match, count)
	}

	return m.Called(ctx, key, cursor, match, count).Get(0).(*redis.ScanCmd)
}

func (m *ClientMock) Append(ctx context.Context, key, value string) *redis.IntCmd {
	if !m.hasStub("Append") {
		return m.client.Append(ctx, key, value)
	}

	return m.Called(ctx, key, value).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) BitCount(ctx context.Context, key string, bitCount *redis.BitCount) *redis.IntCmd {
	if !m.hasStub("BitCount") {
		return m.client.BitCount(ctx, key, bitCount)
	}

	return m.Called(ctx, key, bitCount).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) BitOpAnd(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	if !m.hasStub("BitOpAnd") {
		return m.client.BitOpAnd(ctx, destKey, keys...)
	}

	return m.Called(ctx, destKey, keys).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) BitOpOr(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	if !m.hasStub("BitOpOr") {
		return m.client.BitOpOr(ctx, destKey, keys...)
	}

	return m.Called(ctx, destKey, keys).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) BitOpXor(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	if !m.hasStub("BitOpXor") {
		return m.client.BitOpXor(ctx, destKey, keys...)
	}

	return m.Called(ctx, destKey, keys).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) BitOpNot(ctx context.Context, destKey string, key string) *redis.IntCmd {
	if !m.hasStub("BitOpNot") {
		return m.client.BitOpNot(ctx, destKey, key)
	}

	return m.Called(ctx, destKey, key).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) BitPos(ctx context.Context, key string, bit int64, pos ...int64) *redis.IntCmd {
	if !m.hasStub("BitPos") {
		return m.client.BitPos(ctx, key, bit, pos...)
	}

	return m.Called(ctx, key, bit, pos).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) BitField(ctx context.Context, key string, args ...interface{}) *redis.IntSliceCmd {
	if !m.hasStub("BitField") {
		return m.client.BitField(ctx, key, args)
	}

	return m.Called(ctx, key, args).Get(0).(*redis.IntSliceCmd)
}

func (m *ClientMock) Decr(ctx context.Context, key string) *redis.IntCmd {
	if !m.hasStub("Decr") {
		return m.client.Decr(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) DecrBy(ctx context.Context, key string, decrement int64) *redis.IntCmd {
	if !m.hasStub("DecrBy") {
		return m.client.DecrBy(ctx, key, decrement)
	}

	return m.Called(ctx, key, decrement).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) Get(ctx context.Context, key string) *redis.StringCmd {
	if !m.hasStub("Get") {
		return m.client.Get(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.StringCmd)
}

func (m *ClientMock) GetBit(ctx context.Context, key string, offset int64) *redis.IntCmd {
	if !m.hasStub("GetBit") {
		return m.client.GetBit(ctx, key, offset)
	}

	return m.Called(ctx, key, offset).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) GetEx(ctx context.Context, key string, expiration time.Duration) *redis.StringCmd {
	if !m.hasStub("GetEx") {
		return m.client.GetEx(ctx, key, expiration)
	}

	return m.Called(ctx, key, expiration).Get(0).(*redis.StringCmd)
}

func (m *ClientMock) GetRange(ctx context.Context, key string, start, end int64) *redis.StringCmd {
	if !m.hasStub("GetRange") {
		return m.client.GetRange(ctx, key, start, end)
	}

	return m.Called(ctx, key, start, end).Get(0).(*redis.StringCmd)
}

func (m *ClientMock) GetSet(ctx context.Context, key string, value interface{}) *redis.StringCmd {
	if !m.hasStub("GetSet") {
		return m.client.GetSet(ctx, key, value)
	}

	return m.Called(ctx, key, value).Get(0).(*redis.StringCmd)
}

func (m *ClientMock) Incr(ctx context.Context, key string) *redis.IntCmd {
	if !m.hasStub("Incr") {
		return m.client.Incr(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) IncrBy(ctx context.Context, key string, value int64) *redis.IntCmd {
	if !m.hasStub("IncrBy") {
		return m.client.IncrBy(ctx, key, value)
	}

	return m.Called(ctx, key, value).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) IncrByFloat(ctx context.Context, key string, value float64) *redis.FloatCmd {
	if !m.hasStub("IncrByFloat") {
		return m.client.IncrByFloat(ctx, key, value)
	}

	return m.Called(ctx, key, value).Get(0).(*redis.FloatCmd)
}

func (m *ClientMock) MGet(ctx context.Context, keys ...string) *redis.SliceCmd {
	if !m.hasStub("MGet") {
		return m.client.MGet(ctx, keys...)
	}

	return m.Called(ctx, keys).Get(0).(*redis.SliceCmd)
}

func (m *ClientMock) MSet(ctx context.Context, pairs ...interface{}) *redis.StatusCmd {
	if !m.hasStub("MSet") {
		return m.client.MSet(ctx, pairs...)
	}

	return m.Called(ctx, pairs).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) MSetNX(ctx context.Context, pairs ...interface{}) *redis.BoolCmd {
	if !m.hasStub("MSetNX") {
		return m.client.MSetNX(ctx, pairs...)
	}

	return m.Called(ctx, pairs).Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	if !m.hasStub("Set") {
		return m.client.Set(ctx, key, value, expiration)
	}

	return m.Called(ctx, key, value, expiration).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	if !m.hasStub("SetEX") {
		return m.client.Set(ctx, key, value, expiration)
	}

	return m.Called(ctx, key, value, expiration).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) SetBit(ctx context.Context, key string, offset int64, value int) *redis.IntCmd {
	if !m.hasStub("SetBit") {
		return m.client.SetBit(ctx, key, offset, value)
	}

	return m.Called(ctx, key, offset, value).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	if !m.hasStub("SetNX") {
		return m.client.SetNX(ctx, key, value, expiration)
	}

	return m.Called(ctx, key, value, expiration).Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	if !m.hasStub("SetXX") {
		return m.client.SetXX(ctx, key, value, expiration)
	}

	return m.Called(ctx, key, value, expiration).Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) SetRange(ctx context.Context, key string, offset int64, value string) *redis.IntCmd {
	if !m.hasStub("SetRange") {
		return m.client.SetRange(ctx, key, offset, value)
	}

	return m.Called(ctx, key, offset, value).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) StrLen(ctx context.Context, key string) *redis.IntCmd {
	if !m.hasStub("StrLen") {
		return m.client.StrLen(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) BLPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	if !m.hasStub("BLPop") {
		return m.client.BLPop(ctx, timeout, keys...)
	}

	return m.Called(ctx, timeout, keys).Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) BRPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	if !m.hasStub("BRPop") {
		return m.client.BRPop(ctx, timeout, keys...)
	}

	return m.Called(ctx, timeout, keys).Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *redis.StringCmd {
	if !m.hasStub("BRPopLPush") {
		return m.client.BRPopLPush(ctx, source, destination, timeout)
	}

	return m.Called(ctx, source, destination, timeout).Get(0).(*redis.StringCmd)
}

func (m *ClientMock) LIndex(ctx context.Context, key string, index int64) *redis.StringCmd {
	if !m.hasStub("LIndex") {
		return m.client.LIndex(ctx, key, index)
	}

	return m.Called(ctx, key, index).Get(0).(*redis.StringCmd)
}

func (m *ClientMock) LInsert(ctx context.Context, key, op string, pivot, value interface{}) *redis.IntCmd {
	if !m.hasStub("LInsert") {
		return m.client.LInsert(ctx, key, op, pivot, value)
	}

	return m.Called(ctx, key, op, pivot, value).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) LInsertBefore(ctx context.Context, key string, pivot, value interface{}) *redis.IntCmd {
	if !m.hasStub("LInsertBefore") {
		return m.client.LInsertBefore(ctx, key, pivot, value)
	}

	return m.Called(ctx, key, pivot, value).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) LInsertAfter(ctx context.Context, key string, pivot, value interface{}) *redis.IntCmd {
	if !m.hasStub("LInsertAfter") {
		return m.client.LInsertAfter(ctx, key, pivot, value)
	}

	return m.Called(ctx, key, pivot, value).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) LLen(ctx context.Context, key string) *redis.IntCmd {
	if !m.hasStub("LLen") {
		return m.client.LLen(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) LPos(ctx context.Context, key string, value string, args redis.LPosArgs) *redis.IntCmd {
	if !m.hasStub("LPos") {
		return m.client.LPos(ctx, key, value, args)
	}

	return m.Called(ctx, key, value, args).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) LPosCount(ctx context.Context, key string, value string, count int64, args redis.LPosArgs) *redis.IntSliceCmd {
	if !m.hasStub("LPosCount") {
		return m.client.LPosCount(ctx, key, value, count, args)
	}

	return m.Called(ctx, key, value, count, args).Get(0).(*redis.IntSliceCmd)
}

func (m *ClientMock) LPop(ctx context.Context, key string) *redis.StringCmd {
	if !m.hasStub("LPop") {
		return m.client.LPop(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.StringCmd)
}

func (m *ClientMock) LPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	if !m.hasStub("LPush") {
		return m.client.LPush(ctx, key, values...)
	}

	return m.Called(ctx, key, values).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) LPushX(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	if !m.hasStub("LPushX") {
		return m.client.LPushX(ctx, key, values)
	}

	return m.Called(ctx, key, values).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) LRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	if !m.hasStub("LRange") {
		return m.client.LRange(ctx, key, start, stop)
	}

	return m.Called(ctx, key, start, stop).Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) LRem(ctx context.Context, key string, count int64, value interface{}) *redis.IntCmd {
	if !m.hasStub("LRem") {
		return m.client.LRem(ctx, key, count, value)
	}

	return m.Called(ctx, key, count, value).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) LSet(ctx context.Context, key string, index int64, value interface{}) *redis.StatusCmd {
	if !m.hasStub("LSet") {
		return m.client.LSet(ctx, key, index, value)
	}

	return m.Called(ctx, key, index, value).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) LTrim(ctx context.Context, key string, start, stop int64) *redis.StatusCmd {
	if !m.hasStub("LTrim") {
		return m.client.LTrim(ctx, key, start, stop)
	}

	return m.Called(ctx, key, start, stop).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) RPop(ctx context.Context, key string) *redis.StringCmd {
	if !m.hasStub("RPop") {
		return m.client.RPop(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.StringCmd)
}

func (m *ClientMock) RPopLPush(ctx context.Context, source, destination string) *redis.StringCmd {
	if !m.hasStub("RPopLPush") {
		return m.client.RPopLPush(ctx, source, destination)
	}

	return m.Called(ctx, source, destination).Get(0).(*redis.StringCmd)
}

func (m *ClientMock) RPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	if !m.hasStub("RPush") {
		return m.client.RPush(ctx, key, values...)
	}

	return m.Called(ctx, key, values).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) RPushX(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	if !m.hasStub("RPushX") {
		return m.client.RPushX(ctx, key, values)
	}

	return m.Called(ctx, key, values).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) PFAdd(ctx context.Context, key string, els ...interface{}) *redis.IntCmd {
	if !m.hasStub("PFAdd") {
		return m.client.PFAdd(ctx, key, els...)
	}

	return m.Called(ctx, key, els).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) PFCount(ctx context.Context, keys ...string) *redis.IntCmd {
	if !m.hasStub("PFCount") {
		return m.client.PFCount(ctx, keys...)
	}

	return m.Called(ctx, keys).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) PFMerge(ctx context.Context, dest string, keys ...string) *redis.StatusCmd {
	if !m.hasStub("PFMerge") {
		return m.client.PFMerge(ctx, dest, keys...)
	}

	return m.Called(ctx, dest, keys).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) BgRewriteAOF(ctx context.Context) *redis.StatusCmd {
	if !m.hasStub("BgRewriteAOF") {
		return m.client.BgRewriteAOF(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) BgSave(ctx context.Context) *redis.StatusCmd {
	if !m.hasStub("BgSave") {
		return m.client.BgSave(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClientKill(ctx context.Context, ipPort string) *redis.StatusCmd {
	if !m.hasStub("ClientKill") {
		return m.client.ClientKill(ctx, ipPort)
	}

	return m.Called(ctx, ipPort).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClientList(ctx context.Context) *redis.StringCmd {
	if !m.hasStub("ClientList") {
		return m.client.ClientList(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StringCmd)
}

func (m *ClientMock) ClientPause(ctx context.Context, dur time.Duration) *redis.BoolCmd {
	if !m.hasStub("ClientPause") {
		return m.client.ClientPause(ctx, dur)
	}

	return m.Called(ctx, dur).Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) ConfigGet(ctx context.Context, parameter string) *redis.SliceCmd {
	if !m.hasStub("ConfigGet") {
		return m.client.ConfigGet(ctx, parameter)
	}

	return m.Called(ctx, parameter).Get(0).(*redis.SliceCmd)
}

func (m *ClientMock) ConfigResetStat(ctx context.Context) *redis.StatusCmd {
	if !m.hasStub("ConfigResetStat") {
		return m.client.ConfigResetStat(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ConfigSet(ctx context.Context, parameter, value string) *redis.StatusCmd {
	if !m.hasStub("ConfigSet") {
		return m.client.ConfigSet(ctx, parameter, value)
	}

	return m.Called(ctx, parameter, value).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ConfigRewrite(ctx context.Context) *redis.StatusCmd {
	if !m.hasStub("ConfigRewrite") {
		return m.client.ConfigRewrite(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) DBSize(ctx context.Context) *redis.IntCmd {
	if !m.hasStub("DBSize") {
		return m.client.DBSize(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) FlushAll(ctx context.Context) *redis.StatusCmd {
	if !m.hasStub("FlushAll") {
		return m.client.FlushAll(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) FlushAllAsync(ctx context.Context) *redis.StatusCmd {
	if !m.hasStub("FlushAllAsync") {
		return m.client.FlushAllAsync(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) FlushDB(ctx context.Context) *redis.StatusCmd {
	if !m.hasStub("FlushDB") {
		return m.client.FlushDB(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) FlushDBAsync(ctx context.Context) *redis.StatusCmd {
	if !m.hasStub("FlushDBAsync") {
		return m.client.FlushDBAsync(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) Info(ctx context.Context, section ...string) *redis.StringCmd {
	if !m.hasStub("Info") {
		return m.client.Info(ctx, section...)
	}

	return m.Called(ctx, section).Get(0).(*redis.StringCmd)
}

func (m *ClientMock) LastSave(ctx context.Context) *redis.IntCmd {
	if !m.hasStub("LastSave") {
		return m.client.LastSave(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) Save(ctx context.Context) *redis.StatusCmd {
	if !m.hasStub("Save") {
		return m.client.Save(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) Shutdown(ctx context.Context) *redis.StatusCmd {
	if !m.hasStub("Shutdown") {
		return m.client.Shutdown(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ShutdownSave(ctx context.Context) *redis.StatusCmd {
	if !m.hasStub("ShutdownSave") {
		return m.client.ShutdownSave(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ShutdownNoSave(ctx context.Context) *redis.StatusCmd {
	if !m.hasStub("ShutdownNoSave") {
		return m.client.ShutdownNoSave(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) SlaveOf(ctx context.Context, host, port string) *redis.StatusCmd {
	if !m.hasStub("SlaveOf") {
		return m.client.SlaveOf(ctx, host, port)
	}

	return m.Called(ctx, host, port).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) Time(ctx context.Context) *redis.TimeCmd {
	if !m.hasStub("Time") {
		return m.client.Time(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.TimeCmd)
}

func (m *ClientMock) Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	if !m.hasStub("Eval") {
		return m.client.Eval(ctx, script, keys, args...)
	}

	return m.Called(ctx, script, keys, args).Get(0).(*redis.Cmd)
}

func (m *ClientMock) EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	if !m.hasStub("EvalSha") {
		return m.client.EvalSha(ctx, sha1, keys, args...)
	}

	return m.Called(ctx, sha1, keys, args).Get(0).(*redis.Cmd)
}

func (m *ClientMock) ScriptExists(ctx context.Context, hashes ...string) *redis.BoolSliceCmd {
	if !m.hasStub("ScriptExists") {
		return m.client.ScriptExists(ctx, hashes...)
	}

	return m.Called(ctx, hashes).Get(0).(*redis.BoolSliceCmd)
}

func (m *ClientMock) ScriptFlush(ctx context.Context) *redis.StatusCmd {
	if !m.hasStub("ScriptFlush") {
		return m.client.ScriptFlush(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ScriptKill(ctx context.Context) *redis.StatusCmd {
	if !m.hasStub("ScriptKill") {
		return m.client.ScriptKill(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ScriptLoad(ctx context.Context, script string) *redis.StringCmd {
	if !m.hasStub("ScriptLoad") {
		return m.client.ScriptLoad(ctx, script)
	}

	return m.Called(ctx, script).Get(0).(*redis.StringCmd)
}

func (m *ClientMock) DebugObject(ctx context.Context, key string) *redis.StringCmd {
	if !m.hasStub("DebugObject") {
		return m.client.DebugObject(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.StringCmd)
}

func (m *ClientMock) Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd {
	if !m.hasStub("Publish") {
		return m.client.Publish(ctx, channel, message)
	}

	return m.Called(ctx, channel, message).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) PubSubChannels(ctx context.Context, pattern string) *redis.StringSliceCmd {
	if !m.hasStub("PubSubChannels") {
		return m.client.PubSubChannels(ctx, pattern)
	}

	return m.Called(ctx, pattern).Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) PubSubNumSub(ctx context.Context, channels ...string) *redis.StringIntMapCmd {
	if !m.hasStub("PubSubNumSub") {
		return m.client.PubSubNumSub(ctx, channels...)
	}

	return m.Called(ctx, channels).Get(0).(*redis.StringIntMapCmd)
}

func (m *ClientMock) PubSubNumPat(ctx context.Context) *redis.IntCmd {
	if !m.hasStub("PubSubNumPat") {
		return m.client.PubSubNumPat(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ClusterSlots(ctx context.Context) *redis.ClusterSlotsCmd {
	if !m.hasStub("ClusterSlots") {
		return m.client.ClusterSlots(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.ClusterSlotsCmd)
}

func (m *ClientMock) ClusterNodes(ctx context.Context) *redis.StringCmd {
	if !m.hasStub("ClusterNodes") {
		return m.client.ClusterNodes(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StringCmd)
}

func (m *ClientMock) ClusterMeet(ctx context.Context, host, port string) *redis.StatusCmd {
	if !m.hasStub("ClusterMeet") {
		return m.client.ClusterMeet(ctx, host, port)
	}

	return m.Called(ctx, host, port).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterForget(ctx context.Context, nodeID string) *redis.StatusCmd {
	if !m.hasStub("ClusterForget") {
		return m.client.ClusterForget(ctx, nodeID)
	}

	return m.Called(ctx, nodeID).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterReplicate(ctx context.Context, nodeID string) *redis.StatusCmd {
	if !m.hasStub("ClusterReplicate") {
		return m.client.ClusterReplicate(ctx, nodeID)
	}

	return m.Called(ctx, nodeID).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterResetSoft(ctx context.Context) *redis.StatusCmd {
	if !m.hasStub("ClusterResetSoft") {
		return m.client.ClusterResetSoft(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterResetHard(ctx context.Context) *redis.StatusCmd {
	if !m.hasStub("ClusterResetHard") {
		return m.client.ClusterResetHard(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterInfo(ctx context.Context) *redis.StringCmd {
	if !m.hasStub("ClusterInfo") {
		return m.client.ClusterInfo(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StringCmd)
}

func (m *ClientMock) ClusterKeySlot(ctx context.Context, key string) *redis.IntCmd {
	if !m.hasStub("ClusterKeySlot") {
		return m.client.ClusterKeySlot(ctx, key)
	}

	return m.Called(ctx, key).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ClusterGetKeysInSlot(ctx context.Context, slot int, count int) *redis.StringSliceCmd {
	if !m.hasStub("ClusterGetKeysInSlot") {
		return m.client.ClusterGetKeysInSlot(ctx, slot, count)
	}

	return m.Called(ctx, slot, count).Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) ClusterCountFailureReports(ctx context.Context, nodeID string) *redis.IntCmd {
	if !m.hasStub("ClusterCountFailureReports") {
		return m.client.ClusterCountFailureReports(ctx, nodeID)
	}

	return m.Called(ctx, nodeID).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ClusterCountKeysInSlot(ctx context.Context, slot int) *redis.IntCmd {
	if !m.hasStub("ClusterCountKeysInSlot") {
		return m.client.ClusterCountKeysInSlot(ctx, slot)
	}

	return m.Called(ctx, slot).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ClusterDelSlots(ctx context.Context, slots ...int) *redis.StatusCmd {
	if !m.hasStub("ClusterDelSlots") {
		return m.client.ClusterDelSlots(ctx, slots...)
	}

	return m.Called(ctx, slots).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterDelSlotsRange(ctx context.Context, min, max int) *redis.StatusCmd {
	if !m.hasStub("ClusterDelSlotsRange") {
		return m.client.ClusterDelSlotsRange(ctx, min, max)
	}

	return m.Called(ctx, min, max).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterSaveConfig(ctx context.Context) *redis.StatusCmd {
	if !m.hasStub("ClusterSaveConfig") {
		return m.client.ClusterSaveConfig(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterSlaves(ctx context.Context, nodeID string) *redis.StringSliceCmd {
	if !m.hasStub("ClusterSlaves") {
		return m.client.ClusterSlaves(ctx, nodeID)
	}

	return m.Called(ctx, nodeID).Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) ClusterFailover(ctx context.Context) *redis.StatusCmd {
	if !m.hasStub("ClusterFailover") {
		return m.client.ClusterFailover(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterAddSlots(ctx context.Context, slots ...int) *redis.StatusCmd {
	if !m.hasStub("ClusterAddSlots") {
		return m.client.ClusterAddSlots(ctx, slots...)
	}

	return m.Called(ctx, slots).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterAddSlotsRange(ctx context.Context, min, max int) *redis.StatusCmd {
	if !m.hasStub("ClusterAddSlotsRange") {
		return m.client.ClusterAddSlotsRange(ctx, min, max)
	}

	return m.Called(ctx, min, max).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) GeoAdd(ctx context.Context, key string, geoLocation ...*redis.GeoLocation) *redis.IntCmd {
	if !m.hasStub("GeoAdd") {
		return m.client.GeoAdd(ctx, key, geoLocation...)
	}

	return m.Called(ctx, key, geoLocation).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) GeoPos(ctx context.Context, key string, members ...string) *redis.GeoPosCmd {
	if !m.hasStub("GeoPos") {
		return m.client.GeoPos(ctx, key, members...)
	}

	return m.Called(ctx, key, members).Get(0).(*redis.GeoPosCmd)
}

func (m *ClientMock) GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	if !m.hasStub("GeoRadius") {
		return m.client.GeoRadius(ctx, key, longitude, latitude, query)
	}

	return m.Called(ctx, key, longitude, latitude, query).Get(0).(*redis.GeoLocationCmd)
}

func (m *ClientMock) GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.IntCmd {
	if !m.hasStub("GeoRadiusRO") {
		return m.client.GeoRadiusStore(ctx, key, longitude, latitude, query)
	}

	return m.Called(ctx, key, longitude, latitude, query).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) GeoRadiusByMember(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	if !m.hasStub("GeoRadiusByMember") {
		return m.client.GeoRadiusByMember(ctx, key, member, query)
	}

	return m.Called(ctx, key, member, query).Get(0).(*redis.GeoLocationCmd)
}

func (m *ClientMock) GeoRadiusByMemberStore(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *redis.IntCmd {
	if !m.hasStub("GeoRadiusByMemberRO") {
		return m.client.GeoRadiusByMemberStore(ctx, key, member, query)
	}

	return m.Called(ctx, key, member, query).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) GeoDist(ctx context.Context, key string, member1, member2, unit string) *redis.FloatCmd {
	if !m.hasStub("GeoDist") {
		return m.client.GeoDist(ctx, key, member1, member2, unit)
	}

	return m.Called(ctx, key, member1, member2, unit).Get(0).(*redis.FloatCmd)
}

func (m *ClientMock) GeoHash(ctx context.Context, key string, members ...string) *redis.StringSliceCmd {
	if !m.hasStub("GeoHash") {
		return m.client.GeoHash(ctx, key, members...)
	}

	return m.Called(ctx, key, members).Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) Command(ctx context.Context) *redis.CommandsInfoCmd {
	if !m.hasStub("Command") {
		return m.client.Command(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.CommandsInfoCmd)
}

func (m *ClientMock) ClientKillByFilter(ctx context.Context, keys ...string) *redis.IntCmd {
	if !m.hasStub("ClientKillByFilter") {
		return m.client.ClientKillByFilter(ctx, keys...)
	}

	return m.Called(ctx, keys).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) MemoryUsage(ctx context.Context, key string, samples ...int) *redis.IntCmd {
	if !m.hasStub("MemoryUsage") {
		return m.client.MemoryUsage(ctx, key, samples...)
	}

	return m.Called(ctx, key, samples).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ReadOnly(ctx context.Context) *redis.StatusCmd {
	if !m.hasStub("ReadOnly") {
		return m.client.ReadOnly(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ReadWrite(ctx context.Context) *redis.StatusCmd {
	if !m.hasStub("ReadWrite") {
		return m.client.ReadWrite(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClientID(ctx context.Context) *redis.IntCmd {
	if !m.hasStub("ClientID") {
		return m.client.ClientID(ctx)
	}

	return m.Called(ctx).Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) *redis.ScanCmd {
	if !m.hasStub("ScanType") {
		return m.client.ScanType(ctx, cursor, match, count, keyType)
	}

	return m.Called(ctx, cursor, match, count, keyType).Get(0).(*redis.ScanCmd)
}
