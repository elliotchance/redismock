package redismock

import (
	"github.com/go-redis/redis"
	"time"
)

func (m *ClientMock) Pipeline() redis.Pipeliner {
	if !m.hasStub("Pipeline") {
		return m.client.Pipeline()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(redis.Pipeliner)
}

func (m *ClientMock) Pipelined(fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	if !m.hasStub("Pipelined") {
		return m.client.Pipelined(fn)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).([]redis.Cmder), mockArgs.Get(1).(error)
}

func (m *ClientMock) TxPipelined(fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	if !m.hasStub("TxPipelined") {
		return m.client.TxPipelined(fn)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).([]redis.Cmder), mockArgs.Get(1).(error)
}

func (m *ClientMock) TxPipeline() redis.Pipeliner {
	if !m.hasStub("TxPipeline") {
		return m.client.TxPipeline()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(redis.Pipeliner)
}

func (m *ClientMock) ClientGetName() *redis.StringCmd {
	if !m.hasStub("ClientGetName") {
		return m.client.ClientGetName()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringCmd)
}

func (m *ClientMock) Echo(message interface{}) *redis.StringCmd {
	if !m.hasStub("Echo") {
		return m.client.Echo(message)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringCmd)
}

func (m *ClientMock) Ping() *redis.StatusCmd {
	if !m.hasStub("Ping") {
		return m.client.Ping()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) Quit() *redis.StatusCmd {
	if !m.hasStub("Quit") {
		return m.client.Quit()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) Del(keys ...string) *redis.IntCmd {
	if !m.hasStub("Del") {
		return m.client.Del(keys...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) Unlink(keys ...string) *redis.IntCmd {
	if !m.hasStub("Unlink") {
		return m.client.Unlink(keys...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) Dump(key string) *redis.StringCmd {
	if !m.hasStub("Dump") {
		return m.client.Dump(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringCmd)
}

func (m *ClientMock) Exists(keys ...string) *redis.IntCmd {
	if !m.hasStub("Exists") {
		return m.client.Exists(keys...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) Expire(key string, expiration time.Duration) *redis.BoolCmd {
	if !m.hasStub("Expire") {
		return m.client.Expire(key, expiration)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) ExpireAt(key string, tm time.Time) *redis.BoolCmd {
	if !m.hasStub("ExpireAt") {
		return m.client.ExpireAt(key, tm)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) Keys(pattern string) *redis.StringSliceCmd {
	if !m.hasStub("Keys") {
		return m.client.Keys(pattern)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) Migrate(host, port, key string, db int64, timeout time.Duration) *redis.StatusCmd {
	if !m.hasStub("Migrate") {
		return m.client.Migrate(host, port, key, db, timeout)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) Move(key string, db int64) *redis.BoolCmd {
	if !m.hasStub("Move") {
		return m.client.Move(key, db)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) ObjectRefCount(key string) *redis.IntCmd {
	if !m.hasStub("ObjectRefCount") {
		return m.client.ObjectRefCount(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ObjectEncoding(key string) *redis.StringCmd {
	if !m.hasStub("ObjectEncoding") {
		return m.client.ObjectEncoding(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringCmd)
}

func (m *ClientMock) ObjectIdleTime(key string) *redis.DurationCmd {
	if !m.hasStub("ObjectIdleTime") {
		return m.client.ObjectIdleTime(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.DurationCmd)
}

func (m *ClientMock) Persist(key string) *redis.BoolCmd {
	if !m.hasStub("Persist") {
		return m.client.Persist(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) PExpire(key string, expiration time.Duration) *redis.BoolCmd {
	if !m.hasStub("PExpire") {
		return m.client.PExpire(key, expiration)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) PExpireAt(key string, tm time.Time) *redis.BoolCmd {
	if !m.hasStub("PExpireAt") {
		return m.client.PExpireAt(key, tm)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) PTTL(key string) *redis.DurationCmd {
	if !m.hasStub("PTTL") {
		return m.client.PTTL(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.DurationCmd)
}

func (m *ClientMock) RandomKey() *redis.StringCmd {
	if !m.hasStub("RandomKey") {
		return m.client.RandomKey()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringCmd)
}

func (m *ClientMock) Rename(key, newkey string) *redis.StatusCmd {
	if !m.hasStub("Rename") {
		return m.client.Rename(key, newkey)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) RenameNX(key, newkey string) *redis.BoolCmd {
	if !m.hasStub("RenameNX") {
		return m.client.RenameNX(key, newkey)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) Restore(key string, ttl time.Duration, value string) *redis.StatusCmd {
	if !m.hasStub("Restore") {
		return m.client.Restore(key, ttl, value)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) RestoreReplace(key string, ttl time.Duration, value string) *redis.StatusCmd {
	if !m.hasStub("RestoreReplace") {
		return m.client.RestoreReplace(key, ttl, value)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) Sort(key string, sort *redis.Sort) *redis.StringSliceCmd {
	if !m.hasStub("Sort") {
		return m.client.Sort(key, sort)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) SortStore(key, store string, sort *redis.Sort) *redis.IntCmd {
	if !m.hasStub("SortStore") {
		return m.client.SortStore(key, store, sort)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) SortInterfaces(key string, sort *redis.Sort) *redis.SliceCmd {
	if !m.hasStub("SortInterfaces") {
		return m.client.SortInterfaces(key, sort)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.SliceCmd)
}

func (m *ClientMock) Touch(keys ...string) *redis.IntCmd {
	if !m.hasStub("Touch") {
		return m.client.Touch()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) TTL(key string) *redis.DurationCmd {
	if !m.hasStub("TTL") {
		return m.client.TTL(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.DurationCmd)
}

func (m *ClientMock) Type(key string) *redis.StatusCmd {
	if !m.hasStub("Type") {
		return m.client.Type(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) Scan(cursor uint64, match string, count int64) *redis.ScanCmd {
	if !m.hasStub("Scan") {
		return m.client.Scan(cursor, match, count)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.ScanCmd)
}

func (m *ClientMock) SScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	if !m.hasStub("SScan") {
		return m.client.SScan(key, cursor, match, count)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.ScanCmd)
}

func (m *ClientMock) HScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	if !m.hasStub("HScan") {
		return m.client.HScan(key, cursor, match, count)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.ScanCmd)
}

func (m *ClientMock) ZScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	if !m.hasStub("ZScan") {
		return m.client.ZScan(key, cursor, match, count)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.ScanCmd)
}

func (m *ClientMock) Append(key, value string) *redis.IntCmd {
	if !m.hasStub("Append") {
		return m.client.Append(key, value)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) BitCount(key string, bitCount *redis.BitCount) *redis.IntCmd {
	if !m.hasStub("BitCount") {
		return m.client.BitCount(key, bitCount)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) BitOpAnd(destKey string, keys ...string) *redis.IntCmd {
	if !m.hasStub("BitOpAnd") {
		return m.client.BitOpAnd(destKey, keys...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) BitOpOr(destKey string, keys ...string) *redis.IntCmd {
	if !m.hasStub("BitOpOr") {
		return m.client.BitOpOr(destKey, keys...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) BitOpXor(destKey string, keys ...string) *redis.IntCmd {
	if !m.hasStub("BitOpXor") {
		return m.client.BitOpXor(destKey, keys...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) BitOpNot(destKey string, key string) *redis.IntCmd {
	if !m.hasStub("BitOpNot") {
		return m.client.BitOpNot(destKey, key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) BitPos(key string, bit int64, pos ...int64) *redis.IntCmd {
	if !m.hasStub("BitPos") {
		return m.client.BitPos(key, bit, pos...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) Decr(key string) *redis.IntCmd {
	if !m.hasStub("Decr") {
		return m.client.Decr(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) DecrBy(key string, decrement int64) *redis.IntCmd {
	if !m.hasStub("DecrBy") {
		return m.client.DecrBy(key, decrement)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) Get(key string) *redis.StringCmd {
	if !m.hasStub("Get") {
		return m.client.Get(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringCmd)
}

func (m *ClientMock) GetBit(key string, offset int64) *redis.IntCmd {
	if !m.hasStub("GetBit") {
		return m.client.GetBit(key, offset)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) GetRange(key string, start, end int64) *redis.StringCmd {
	if !m.hasStub("GetRange") {
		return m.client.GetRange(key, start, end)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringCmd)
}

func (m *ClientMock) GetSet(key string, value interface{}) *redis.StringCmd {
	if !m.hasStub("GetSet") {
		return m.client.GetSet(key, value)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringCmd)
}

func (m *ClientMock) Incr(key string) *redis.IntCmd {
	if !m.hasStub("Incr") {
		return m.client.Incr(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) IncrBy(key string, value int64) *redis.IntCmd {
	if !m.hasStub("IncrBy") {
		return m.client.IncrBy(key, value)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) IncrByFloat(key string, value float64) *redis.FloatCmd {
	if !m.hasStub("IncrByFloat") {
		return m.client.IncrByFloat(key, value)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.FloatCmd)
}

func (m *ClientMock) MGet(keys ...string) *redis.SliceCmd {
	if !m.hasStub("MGet") {
		return m.client.MGet()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.SliceCmd)
}

func (m *ClientMock) MSet(pairs ...interface{}) *redis.StatusCmd {
	if !m.hasStub("MSet") {
		return m.client.MSet()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) MSetNX(pairs ...interface{}) *redis.BoolCmd {
	if !m.hasStub("MSetNX") {
		return m.client.MSetNX()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	if !m.hasStub("Set") {
		return m.client.Set(key, value, expiration)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) SetBit(key string, offset int64, value int) *redis.IntCmd {
	if !m.hasStub("SetBit") {
		return m.client.SetBit(key, offset, value)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) SetNX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	if !m.hasStub("SetNX") {
		return m.client.SetNX(key, value, expiration)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) SetXX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	if !m.hasStub("SetXX") {
		return m.client.SetXX(key, value, expiration)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) SetRange(key string, offset int64, value string) *redis.IntCmd {
	if !m.hasStub("SetRange") {
		return m.client.SetRange(key, offset, value)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) StrLen(key string) *redis.IntCmd {
	if !m.hasStub("StrLen") {
		return m.client.StrLen(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) HDel(key string, fields ...string) *redis.IntCmd {
	if !m.hasStub("HDel") {
		return m.client.HDel(key, fields...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) HExists(key, field string) *redis.BoolCmd {
	if !m.hasStub("HExists") {
		return m.client.HExists(key, field)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) HGet(key, field string) *redis.StringCmd {
	if !m.hasStub("HGet") {
		return m.client.HGet(key, field)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringCmd)
}

func (m *ClientMock) HGetAll(key string) *redis.StringStringMapCmd {
	if !m.hasStub("HGetAll") {
		return m.client.HGetAll(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringStringMapCmd)
}

func (m *ClientMock) HIncrBy(key, field string, incr int64) *redis.IntCmd {
	if !m.hasStub("HIncrBy") {
		return m.client.HIncrBy(key, field, incr)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) HIncrByFloat(key, field string, incr float64) *redis.FloatCmd {
	if !m.hasStub("HIncrByFloat") {
		return m.client.HIncrByFloat(key, field, incr)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.FloatCmd)
}

func (m *ClientMock) HKeys(key string) *redis.StringSliceCmd {
	if !m.hasStub("HKeys") {
		return m.client.HKeys(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) HLen(key string) *redis.IntCmd {
	if !m.hasStub("HLen") {
		return m.client.HLen(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) HMGet(key string, fields ...string) *redis.SliceCmd {
	if !m.hasStub("HMGet") {
		return m.client.HMGet(key, fields...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.SliceCmd)
}

func (m *ClientMock) HMSet(key string, fields map[string]interface{}) *redis.StatusCmd {
	if !m.hasStub("HMSet") {
		return m.client.HMSet(key, fields)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) HSet(key, field string, value interface{}) *redis.BoolCmd {
	if !m.hasStub("HSet") {
		return m.client.HSet(key, field, value)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) HSetNX(key, field string, value interface{}) *redis.BoolCmd {
	if !m.hasStub("HSetNX") {
		return m.client.HSetNX(key, field, value)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) HVals(key string) *redis.StringSliceCmd {
	if !m.hasStub("HVals") {
		return m.client.HVals(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) BLPop(timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	if !m.hasStub("BLPop") {
		return m.client.BLPop(timeout, keys...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) BRPop(timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	if !m.hasStub("BRPop") {
		return m.client.BRPop(timeout, keys...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) BRPopLPush(source, destination string, timeout time.Duration) *redis.StringCmd {
	if !m.hasStub("BRPopLPush") {
		return m.client.BRPopLPush(source, destination, timeout)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringCmd)
}

func (m *ClientMock) LIndex(key string, index int64) *redis.StringCmd {
	if !m.hasStub("LIndex") {
		return m.client.LIndex(key, index)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringCmd)
}

func (m *ClientMock) LInsert(key, op string, pivot, value interface{}) *redis.IntCmd {
	if !m.hasStub("LInsert") {
		return m.client.LInsert(key, op, pivot, value)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) LInsertBefore(key string, pivot, value interface{}) *redis.IntCmd {
	if !m.hasStub("LInsertBefore") {
		return m.client.LInsertBefore(key, pivot, value)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) LInsertAfter(key string, pivot, value interface{}) *redis.IntCmd {
	if !m.hasStub("LInsertAfter") {
		return m.client.LInsertAfter(key, pivot, value)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) LLen(key string) *redis.IntCmd {
	if !m.hasStub("LLen") {
		return m.client.LLen(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) LPop(key string) *redis.StringCmd {
	if !m.hasStub("LPop") {
		return m.client.LPop(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringCmd)
}

func (m *ClientMock) LPush(key string, values ...interface{}) *redis.IntCmd {
	if !m.hasStub("LPush") {
		return m.client.LPush(key, values...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) LPushX(key string, value interface{}) *redis.IntCmd {
	if !m.hasStub("LPushX") {
		return m.client.LPushX(key, value)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) LRange(key string, start, stop int64) *redis.StringSliceCmd {
	if !m.hasStub("LRange") {
		return m.client.LRange(key, start, stop)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) LRem(key string, count int64, value interface{}) *redis.IntCmd {
	if !m.hasStub("LRem") {
		return m.client.LRem(key, count, value)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) LSet(key string, index int64, value interface{}) *redis.StatusCmd {
	if !m.hasStub("LSet") {
		return m.client.LSet(key, index, value)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) LTrim(key string, start, stop int64) *redis.StatusCmd {
	if !m.hasStub("LTrim") {
		return m.client.LTrim(key, start, stop)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) RPop(key string) *redis.StringCmd {
	if !m.hasStub("RPop") {
		return m.client.RPop(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringCmd)
}

func (m *ClientMock) RPopLPush(source, destination string) *redis.StringCmd {
	if !m.hasStub("RPopLPush") {
		return m.client.RPopLPush(source, destination)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringCmd)
}

func (m *ClientMock) RPush(key string, values ...interface{}) *redis.IntCmd {
	if !m.hasStub("RPush") {
		return m.client.RPush(key, values...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) RPushX(key string, value interface{}) *redis.IntCmd {
	if !m.hasStub("RPushX") {
		return m.client.RPushX(key, value)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) SAdd(key string, members ...interface{}) *redis.IntCmd {
	if !m.hasStub("SAdd") {
		return m.client.SAdd(key, members...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) SCard(key string) *redis.IntCmd {
	if !m.hasStub("SCard") {
		return m.client.SCard(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) SDiff(keys ...string) *redis.StringSliceCmd {
	if !m.hasStub("SDiff") {
		return m.client.SDiff()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) SDiffStore(destination string, keys ...string) *redis.IntCmd {
	if !m.hasStub("SDiffStore") {
		return m.client.SDiffStore(destination, keys...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) SInter(keys ...string) *redis.StringSliceCmd {
	if !m.hasStub("SInter") {
		return m.client.SInter()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) SInterStore(destination string, keys ...string) *redis.IntCmd {
	if !m.hasStub("SInterStore") {
		return m.client.SInterStore(destination, keys...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) SIsMember(key string, member interface{}) *redis.BoolCmd {
	if !m.hasStub("SIsMember") {
		return m.client.SIsMember(key, member)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) SMembers(key string) *redis.StringSliceCmd {
	if !m.hasStub("SMembers") {
		return m.client.SMembers(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) SMembersMap(key string) *redis.StringStructMapCmd {
	if !m.hasStub("SMembersMap") {
		return m.client.SMembersMap(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringStructMapCmd)
}

func (m *ClientMock) SMove(source, destination string, member interface{}) *redis.BoolCmd {
	if !m.hasStub("SMove") {
		return m.client.SMove(source, destination, member)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) SPop(key string) *redis.StringCmd {
	if !m.hasStub("SPop") {
		return m.client.SPop(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringCmd)
}

func (m *ClientMock) SPopN(key string, count int64) *redis.StringSliceCmd {
	if !m.hasStub("SPopN") {
		return m.client.SPopN(key, count)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) SRandMember(key string) *redis.StringCmd {
	if !m.hasStub("SRandMember") {
		return m.client.SRandMember(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringCmd)
}

func (m *ClientMock) SRandMemberN(key string, count int64) *redis.StringSliceCmd {
	if !m.hasStub("SRandMemberN") {
		return m.client.SRandMemberN(key, count)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) SRem(key string, members ...interface{}) *redis.IntCmd {
	if !m.hasStub("SRem") {
		return m.client.SRem(key, members...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) SUnion(keys ...string) *redis.StringSliceCmd {
	if !m.hasStub("SUnion") {
		return m.client.SUnion(keys...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) SUnionStore(destination string, keys ...string) *redis.IntCmd {
	if !m.hasStub("SUnionStore") {
		return m.client.SUnionStore(destination, keys...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZAdd(key string, members ...redis.Z) *redis.IntCmd {
	if !m.hasStub("ZAdd") {
		return m.client.ZAdd(key, members...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZAddNX(key string, members ...redis.Z) *redis.IntCmd {
	if !m.hasStub("ZAddNX") {
		return m.client.ZAddNX(key, members...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZAddXX(key string, members ...redis.Z) *redis.IntCmd {
	if !m.hasStub("ZAddXX") {
		return m.client.ZAddXX(key, members...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZAddCh(key string, members ...redis.Z) *redis.IntCmd {
	if !m.hasStub("ZAddCh") {
		return m.client.ZAddCh(key, members...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZAddNXCh(key string, members ...redis.Z) *redis.IntCmd {
	if !m.hasStub("ZAddNXCh") {
		return m.client.ZAddNXCh(key, members...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZAddXXCh(key string, members ...redis.Z) *redis.IntCmd {
	if !m.hasStub("ZAddXXCh") {
		return m.client.ZAddXXCh(key, members...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZIncr(key string, member redis.Z) *redis.FloatCmd {
	if !m.hasStub("ZIncr") {
		return m.client.ZIncr(key, member)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.FloatCmd)
}

func (m *ClientMock) ZIncrNX(key string, member redis.Z) *redis.FloatCmd {
	if !m.hasStub("ZIncrNX") {
		return m.client.ZIncrNX(key, member)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.FloatCmd)
}

func (m *ClientMock) ZIncrXX(key string, member redis.Z) *redis.FloatCmd {
	if !m.hasStub("ZIncrXX") {
		return m.client.ZIncrXX(key, member)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.FloatCmd)
}

func (m *ClientMock) ZCard(key string) *redis.IntCmd {
	if !m.hasStub("ZCard") {
		return m.client.ZCard(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZCount(key, min, max string) *redis.IntCmd {
	if !m.hasStub("ZCount") {
		return m.client.ZCount(key, min, max)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZLexCount(key, min, max string) *redis.IntCmd {
	if !m.hasStub("ZLexCount") {
		return m.client.ZLexCount(key, min, max)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZIncrBy(key string, increment float64, member string) *redis.FloatCmd {
	if !m.hasStub("ZIncrBy") {
		return m.client.ZIncrBy(key, increment, member)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.FloatCmd)
}

func (m *ClientMock) ZInterStore(destination string, store redis.ZStore, keys ...string) *redis.IntCmd {
	if !m.hasStub("ZInterStore") {
		return m.client.ZInterStore(destination, store, keys...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZRange(key string, start, stop int64) *redis.StringSliceCmd {
	if !m.hasStub("ZRange") {
		return m.client.ZRange(key, start, stop)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) ZRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd {
	if !m.hasStub("ZRangeWithScores") {
		return m.client.ZRangeWithScores(key, start, stop)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.ZSliceCmd)
}

func (m *ClientMock) ZRangeByScore(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	if !m.hasStub("ZRangeByScore") {
		return m.client.ZRangeByScore(key, opt)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) ZRangeByLex(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	if !m.hasStub("ZRangeByLex") {
		return m.client.ZRangeByLex(key, opt)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) ZRangeByScoreWithScores(key string, opt redis.ZRangeBy) *redis.ZSliceCmd {
	if !m.hasStub("ZRangeByScoreWithScores") {
		return m.client.ZRangeByScoreWithScores(key, opt)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.ZSliceCmd)
}

func (m *ClientMock) ZRank(key, member string) *redis.IntCmd {
	if !m.hasStub("ZRank") {
		return m.client.ZRank(key, member)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZRem(key string, members ...interface{}) *redis.IntCmd {
	if !m.hasStub("ZRem") {
		return m.client.ZRem(key, members...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZRemRangeByRank(key string, start, stop int64) *redis.IntCmd {
	if !m.hasStub("ZRemRangeByRank") {
		return m.client.ZRemRangeByRank(key, start, stop)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZRemRangeByScore(key, min, max string) *redis.IntCmd {
	if !m.hasStub("ZRemRangeByScore") {
		return m.client.ZRemRangeByScore(key, min, max)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZRemRangeByLex(key, min, max string) *redis.IntCmd {
	if !m.hasStub("ZRemRangeByLex") {
		return m.client.ZRemRangeByLex(key, min, max)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZRevRange(key string, start, stop int64) *redis.StringSliceCmd {
	if !m.hasStub("ZRevRange") {
		return m.client.ZRevRange(key, start, stop)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) ZRevRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd {
	if !m.hasStub("ZRevRangeWithScores") {
		return m.client.ZRevRangeWithScores(key, start, stop)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.ZSliceCmd)
}

func (m *ClientMock) ZRevRangeByScore(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	if !m.hasStub("ZRevRangeByScore") {
		return m.client.ZRevRangeByScore(key, opt)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) ZRevRangeByLex(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	if !m.hasStub("ZRevRangeByLex") {
		return m.client.ZRevRangeByLex(key, opt)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) ZRevRangeByScoreWithScores(key string, opt redis.ZRangeBy) *redis.ZSliceCmd {
	if !m.hasStub("ZRevRangeByScoreWithScores") {
		return m.client.ZRevRangeByScoreWithScores(key, opt)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.ZSliceCmd)
}

func (m *ClientMock) ZRevRank(key, member string) *redis.IntCmd {
	if !m.hasStub("ZRevRank") {
		return m.client.ZRevRank(key, member)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ZScore(key, member string) *redis.FloatCmd {
	if !m.hasStub("ZScore") {
		return m.client.ZScore(key, member)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.FloatCmd)
}

func (m *ClientMock) ZUnionStore(dest string, store redis.ZStore, keys ...string) *redis.IntCmd {
	if !m.hasStub("ZUnionStore") {
		return m.client.ZUnionStore(dest, store, keys...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) PFAdd(key string, els ...interface{}) *redis.IntCmd {
	if !m.hasStub("PFAdd") {
		return m.client.PFAdd(key, els...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) PFCount(keys ...string) *redis.IntCmd {
	if !m.hasStub("PFCount") {
		return m.client.PFCount()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) PFMerge(dest string, keys ...string) *redis.StatusCmd {
	if !m.hasStub("PFMerge") {
		return m.client.PFMerge(dest, keys...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) BgRewriteAOF() *redis.StatusCmd {
	if !m.hasStub("BgRewriteAOF") {
		return m.client.BgRewriteAOF()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) BgSave() *redis.StatusCmd {
	if !m.hasStub("BgSave") {
		return m.client.BgSave()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClientKill(ipPort string) *redis.StatusCmd {
	if !m.hasStub("ClientKill") {
		return m.client.ClientKill(ipPort)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClientList() *redis.StringCmd {
	if !m.hasStub("ClientList") {
		return m.client.ClientList()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringCmd)
}

func (m *ClientMock) ClientPause(dur time.Duration) *redis.BoolCmd {
	if !m.hasStub("ClientPause") {
		return m.client.ClientPause(dur)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) ConfigGet(parameter string) *redis.SliceCmd {
	if !m.hasStub("ConfigGet") {
		return m.client.ConfigGet(parameter)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.SliceCmd)
}

func (m *ClientMock) ConfigResetStat() *redis.StatusCmd {
	if !m.hasStub("ConfigResetStat") {
		return m.client.ConfigResetStat()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ConfigSet(parameter, value string) *redis.StatusCmd {
	if !m.hasStub("ConfigSet") {
		return m.client.ConfigSet(parameter, value)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ConfigRewrite() *redis.StatusCmd {
	if !m.hasStub("ConfigRewrite") {
		return m.client.ConfigRewrite()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) DBSize() *redis.IntCmd {
	if !m.hasStub("DBSize") {
		return m.client.DBSize()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) FlushAll() *redis.StatusCmd {
	if !m.hasStub("FlushAll") {
		return m.client.FlushAll()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) FlushAllAsync() *redis.StatusCmd {
	if !m.hasStub("FlushAllAsync") {
		return m.client.FlushAllAsync()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) FlushDB() *redis.StatusCmd {
	if !m.hasStub("FlushDB") {
		return m.client.FlushDB()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) FlushDBAsync() *redis.StatusCmd {
	if !m.hasStub("FlushDBAsync") {
		return m.client.FlushDBAsync()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) Info(section ...string) *redis.StringCmd {
	if !m.hasStub("Info") {
		return m.client.Info()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringCmd)
}

func (m *ClientMock) LastSave() *redis.IntCmd {
	if !m.hasStub("LastSave") {
		return m.client.LastSave()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) Save() *redis.StatusCmd {
	if !m.hasStub("Save") {
		return m.client.Save()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) Shutdown() *redis.StatusCmd {
	if !m.hasStub("Shutdown") {
		return m.client.Shutdown()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ShutdownSave() *redis.StatusCmd {
	if !m.hasStub("ShutdownSave") {
		return m.client.ShutdownSave()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ShutdownNoSave() *redis.StatusCmd {
	if !m.hasStub("ShutdownNoSave") {
		return m.client.ShutdownNoSave()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) SlaveOf(host, port string) *redis.StatusCmd {
	if !m.hasStub("SlaveOf") {
		return m.client.SlaveOf(host, port)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) Time() *redis.TimeCmd {
	if !m.hasStub("Time") {
		return m.client.Time()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.TimeCmd)
}

func (m *ClientMock) Eval(script string, keys []string, args ...interface{}) *redis.Cmd {
	if !m.hasStub("Eval") {
		return m.client.Eval(script, keys, args...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.Cmd)
}

func (m *ClientMock) EvalSha(sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	if !m.hasStub("EvalSha") {
		return m.client.EvalSha(sha1, keys, args...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.Cmd)
}

func (m *ClientMock) ScriptExists(hashes ...string) *redis.BoolSliceCmd {
	if !m.hasStub("ScriptExists") {
		return m.client.ScriptExists()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.BoolSliceCmd)
}

func (m *ClientMock) ScriptFlush() *redis.StatusCmd {
	if !m.hasStub("ScriptFlush") {
		return m.client.ScriptFlush()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ScriptKill() *redis.StatusCmd {
	if !m.hasStub("ScriptKill") {
		return m.client.ScriptKill()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ScriptLoad(script string) *redis.StringCmd {
	if !m.hasStub("ScriptLoad") {
		return m.client.ScriptLoad(script)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringCmd)
}

func (m *ClientMock) DebugObject(key string) *redis.StringCmd {
	if !m.hasStub("DebugObject") {
		return m.client.DebugObject(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringCmd)
}

func (m *ClientMock) Publish(channel string, message interface{}) *redis.IntCmd {
	if !m.hasStub("Publish") {
		return m.client.Publish(channel, message)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) PubSubChannels(pattern string) *redis.StringSliceCmd {
	if !m.hasStub("PubSubChannels") {
		return m.client.PubSubChannels(pattern)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) PubSubNumSub(channels ...string) *redis.StringIntMapCmd {
	if !m.hasStub("PubSubNumSub") {
		return m.client.PubSubNumSub()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringIntMapCmd)
}

func (m *ClientMock) PubSubNumPat() *redis.IntCmd {
	if !m.hasStub("PubSubNumPat") {
		return m.client.PubSubNumPat()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ClusterSlots() *redis.ClusterSlotsCmd {
	if !m.hasStub("ClusterSlots") {
		return m.client.ClusterSlots()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.ClusterSlotsCmd)
}

func (m *ClientMock) ClusterNodes() *redis.StringCmd {
	if !m.hasStub("ClusterNodes") {
		return m.client.ClusterNodes()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringCmd)
}

func (m *ClientMock) ClusterMeet(host, port string) *redis.StatusCmd {
	if !m.hasStub("ClusterMeet") {
		return m.client.ClusterMeet(host, port)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterForget(nodeID string) *redis.StatusCmd {
	if !m.hasStub("ClusterForget") {
		return m.client.ClusterForget(nodeID)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterReplicate(nodeID string) *redis.StatusCmd {
	if !m.hasStub("ClusterReplicate") {
		return m.client.ClusterReplicate(nodeID)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterResetSoft() *redis.StatusCmd {
	if !m.hasStub("ClusterResetSoft") {
		return m.client.ClusterResetSoft()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterResetHard() *redis.StatusCmd {
	if !m.hasStub("ClusterResetHard") {
		return m.client.ClusterResetHard()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterInfo() *redis.StringCmd {
	if !m.hasStub("ClusterInfo") {
		return m.client.ClusterInfo()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringCmd)
}

func (m *ClientMock) ClusterKeySlot(key string) *redis.IntCmd {
	if !m.hasStub("ClusterKeySlot") {
		return m.client.ClusterKeySlot(key)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ClusterCountFailureReports(nodeID string) *redis.IntCmd {
	if !m.hasStub("ClusterCountFailureReports") {
		return m.client.ClusterCountFailureReports(nodeID)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ClusterCountKeysInSlot(slot int) *redis.IntCmd {
	if !m.hasStub("ClusterCountKeysInSlot") {
		return m.client.ClusterCountKeysInSlot(slot)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ClusterDelSlots(slots ...int) *redis.StatusCmd {
	if !m.hasStub("ClusterDelSlots") {
		return m.client.ClusterDelSlots(slots...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterDelSlotsRange(min, max int) *redis.StatusCmd {
	if !m.hasStub("ClusterDelSlotsRange") {
		return m.client.ClusterDelSlotsRange(min, max)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterSaveConfig() *redis.StatusCmd {
	if !m.hasStub("ClusterSaveConfig") {
		return m.client.ClusterSaveConfig()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterSlaves(nodeID string) *redis.StringSliceCmd {
	if !m.hasStub("ClusterSlaves") {
		return m.client.ClusterSlaves(nodeID)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) ClusterFailover() *redis.StatusCmd {
	if !m.hasStub("ClusterFailover") {
		return m.client.ClusterFailover()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterAddSlots(slots ...int) *redis.StatusCmd {
	if !m.hasStub("ClusterAddSlots") {
		return m.client.ClusterAddSlots()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterAddSlotsRange(min, max int) *redis.StatusCmd {
	if !m.hasStub("ClusterAddSlotsRange") {
		return m.client.ClusterAddSlotsRange(min, max)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) GeoAdd(key string, geoLocation ...*redis.GeoLocation) *redis.IntCmd {
	if !m.hasStub("GeoAdd") {
		return m.client.GeoAdd(key, geoLocation...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.IntCmd)
}

func (m *ClientMock) GeoPos(key string, members ...string) *redis.GeoPosCmd {
	if !m.hasStub("GeoPos") {
		return m.client.GeoPos(key, members...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.GeoPosCmd)
}

func (m *ClientMock) GeoRadius(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	if !m.hasStub("GeoRadius") {
		return m.client.GeoRadius(key, longitude, latitude, query)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.GeoLocationCmd)
}

func (m *ClientMock) GeoRadiusRO(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	if !m.hasStub("GeoRadiusRO") {
		return m.client.GeoRadiusRO(key, longitude, latitude, query)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.GeoLocationCmd)
}

func (m *ClientMock) GeoRadiusByMember(key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	if !m.hasStub("GeoRadiusByMember") {
		return m.client.GeoRadiusByMember(key, member, query)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.GeoLocationCmd)
}

func (m *ClientMock) GeoRadiusByMemberRO(key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	if !m.hasStub("GeoRadiusByMemberRO") {
		return m.client.GeoRadiusByMemberRO(key, member, query)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.GeoLocationCmd)
}

func (m *ClientMock) GeoDist(key string, member1, member2, unit string) *redis.FloatCmd {
	if !m.hasStub("GeoDist") {
		return m.client.GeoDist(key, member1, member2, unit)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.FloatCmd)
}

func (m *ClientMock) GeoHash(key string, members ...string) *redis.StringSliceCmd {
	if !m.hasStub("GeoHash") {
		return m.client.GeoHash(key, members...)
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) Command() *redis.CommandsInfoCmd {
	if !m.hasStub("Command") {
		return m.client.Command()
	}

	mockArgs := m.Called()

	return mockArgs.Get(0).(*redis.CommandsInfoCmd)
}
