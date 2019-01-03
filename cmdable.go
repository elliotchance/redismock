package redismock

import (
	"github.com/go-redis/redis"
	"time"
)

func (m *ClientMock) Pipeline() redis.Pipeliner {
	if !m.hasStub("Pipeline") {
		return m.client.Pipeline()
	}

	return m.Called().Get(0).(redis.Pipeliner)
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

	return m.Called().Get(0).(redis.Pipeliner)
}

func (m *ClientMock) ClientGetName() *redis.StringCmd {
	if !m.hasStub("ClientGetName") {
		return m.client.ClientGetName()
	}

	return m.Called().Get(0).(*redis.StringCmd)
}

func (m *ClientMock) Echo(message interface{}) *redis.StringCmd {
	if !m.hasStub("Echo") {
		return m.client.Echo(message)
	}

	return m.Called().Get(0).(*redis.StringCmd)
}

func (m *ClientMock) Ping() *redis.StatusCmd {
	if !m.hasStub("Ping") {
		return m.client.Ping()
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) Quit() *redis.StatusCmd {
	if !m.hasStub("Quit") {
		return m.client.Quit()
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) Del(keys ...string) *redis.IntCmd {
	if !m.hasStub("Del") {
		return m.client.Del(keys...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) Unlink(keys ...string) *redis.IntCmd {
	if !m.hasStub("Unlink") {
		return m.client.Unlink(keys...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) Dump(key string) *redis.StringCmd {
	if !m.hasStub("Dump") {
		return m.client.Dump(key)
	}

	return m.Called().Get(0).(*redis.StringCmd)
}

func (m *ClientMock) Exists(keys ...string) *redis.IntCmd {
	if !m.hasStub("Exists") {
		return m.client.Exists(keys...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) Expire(key string, expiration time.Duration) *redis.BoolCmd {
	if !m.hasStub("Expire") {
		return m.client.Expire(key, expiration)
	}

	return m.Called().Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) ExpireAt(key string, tm time.Time) *redis.BoolCmd {
	if !m.hasStub("ExpireAt") {
		return m.client.ExpireAt(key, tm)
	}

	return m.Called().Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) Keys(pattern string) *redis.StringSliceCmd {
	if !m.hasStub("Keys") {
		return m.client.Keys(pattern)
	}

	return m.Called().Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) Migrate(host, port, key string, db int64, timeout time.Duration) *redis.StatusCmd {
	if !m.hasStub("Migrate") {
		return m.client.Migrate(host, port, key, db, timeout)
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) Move(key string, db int64) *redis.BoolCmd {
	if !m.hasStub("Move") {
		return m.client.Move(key, db)
	}

	return m.Called().Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) ObjectRefCount(key string) *redis.IntCmd {
	if !m.hasStub("ObjectRefCount") {
		return m.client.ObjectRefCount(key)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ObjectEncoding(key string) *redis.StringCmd {
	if !m.hasStub("ObjectEncoding") {
		return m.client.ObjectEncoding(key)
	}

	return m.Called().Get(0).(*redis.StringCmd)
}

func (m *ClientMock) ObjectIdleTime(key string) *redis.DurationCmd {
	if !m.hasStub("ObjectIdleTime") {
		return m.client.ObjectIdleTime(key)
	}

	return m.Called().Get(0).(*redis.DurationCmd)
}

func (m *ClientMock) Persist(key string) *redis.BoolCmd {
	if !m.hasStub("Persist") {
		return m.client.Persist(key)
	}

	return m.Called().Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) PExpire(key string, expiration time.Duration) *redis.BoolCmd {
	if !m.hasStub("PExpire") {
		return m.client.PExpire(key, expiration)
	}

	return m.Called().Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) PExpireAt(key string, tm time.Time) *redis.BoolCmd {
	if !m.hasStub("PExpireAt") {
		return m.client.PExpireAt(key, tm)
	}

	return m.Called().Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) PTTL(key string) *redis.DurationCmd {
	if !m.hasStub("PTTL") {
		return m.client.PTTL(key)
	}

	return m.Called().Get(0).(*redis.DurationCmd)
}

func (m *ClientMock) RandomKey() *redis.StringCmd {
	if !m.hasStub("RandomKey") {
		return m.client.RandomKey()
	}

	return m.Called().Get(0).(*redis.StringCmd)
}

func (m *ClientMock) Rename(key, newkey string) *redis.StatusCmd {
	if !m.hasStub("Rename") {
		return m.client.Rename(key, newkey)
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) RenameNX(key, newkey string) *redis.BoolCmd {
	if !m.hasStub("RenameNX") {
		return m.client.RenameNX(key, newkey)
	}

	return m.Called().Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) Restore(key string, ttl time.Duration, value string) *redis.StatusCmd {
	if !m.hasStub("Restore") {
		return m.client.Restore(key, ttl, value)
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) RestoreReplace(key string, ttl time.Duration, value string) *redis.StatusCmd {
	if !m.hasStub("RestoreReplace") {
		return m.client.RestoreReplace(key, ttl, value)
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) Sort(key string, sort *redis.Sort) *redis.StringSliceCmd {
	if !m.hasStub("Sort") {
		return m.client.Sort(key, sort)
	}

	return m.Called().Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) SortStore(key, store string, sort *redis.Sort) *redis.IntCmd {
	if !m.hasStub("SortStore") {
		return m.client.SortStore(key, store, sort)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) SortInterfaces(key string, sort *redis.Sort) *redis.SliceCmd {
	if !m.hasStub("SortInterfaces") {
		return m.client.SortInterfaces(key, sort)
	}

	return m.Called().Get(0).(*redis.SliceCmd)
}

func (m *ClientMock) Touch(keys ...string) *redis.IntCmd {
	if !m.hasStub("Touch") {
		return m.client.Touch()
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) TTL(key string) *redis.DurationCmd {
	if !m.hasStub("TTL") {
		return m.client.TTL(key)
	}

	return m.Called().Get(0).(*redis.DurationCmd)
}

func (m *ClientMock) Type(key string) *redis.StatusCmd {
	if !m.hasStub("Type") {
		return m.client.Type(key)
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) Scan(cursor uint64, match string, count int64) *redis.ScanCmd {
	if !m.hasStub("Scan") {
		return m.client.Scan(cursor, match, count)
	}

	return m.Called().Get(0).(*redis.ScanCmd)
}

func (m *ClientMock) SScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	if !m.hasStub("SScan") {
		return m.client.SScan(key, cursor, match, count)
	}

	return m.Called().Get(0).(*redis.ScanCmd)
}

func (m *ClientMock) Append(key, value string) *redis.IntCmd {
	if !m.hasStub("Append") {
		return m.client.Append(key, value)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) BitCount(key string, bitCount *redis.BitCount) *redis.IntCmd {
	if !m.hasStub("BitCount") {
		return m.client.BitCount(key, bitCount)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) BitOpAnd(destKey string, keys ...string) *redis.IntCmd {
	if !m.hasStub("BitOpAnd") {
		return m.client.BitOpAnd(destKey, keys...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) BitOpOr(destKey string, keys ...string) *redis.IntCmd {
	if !m.hasStub("BitOpOr") {
		return m.client.BitOpOr(destKey, keys...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) BitOpXor(destKey string, keys ...string) *redis.IntCmd {
	if !m.hasStub("BitOpXor") {
		return m.client.BitOpXor(destKey, keys...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) BitOpNot(destKey string, key string) *redis.IntCmd {
	if !m.hasStub("BitOpNot") {
		return m.client.BitOpNot(destKey, key)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) BitPos(key string, bit int64, pos ...int64) *redis.IntCmd {
	if !m.hasStub("BitPos") {
		return m.client.BitPos(key, bit, pos...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) Decr(key string) *redis.IntCmd {
	if !m.hasStub("Decr") {
		return m.client.Decr(key)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) DecrBy(key string, decrement int64) *redis.IntCmd {
	if !m.hasStub("DecrBy") {
		return m.client.DecrBy(key, decrement)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) Get(key string) *redis.StringCmd {
	if !m.hasStub("Get") {
		return m.client.Get(key)
	}

	return m.Called().Get(0).(*redis.StringCmd)
}

func (m *ClientMock) GetBit(key string, offset int64) *redis.IntCmd {
	if !m.hasStub("GetBit") {
		return m.client.GetBit(key, offset)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) GetRange(key string, start, end int64) *redis.StringCmd {
	if !m.hasStub("GetRange") {
		return m.client.GetRange(key, start, end)
	}

	return m.Called().Get(0).(*redis.StringCmd)
}

func (m *ClientMock) GetSet(key string, value interface{}) *redis.StringCmd {
	if !m.hasStub("GetSet") {
		return m.client.GetSet(key, value)
	}

	return m.Called().Get(0).(*redis.StringCmd)
}

func (m *ClientMock) Incr(key string) *redis.IntCmd {
	if !m.hasStub("Incr") {
		return m.client.Incr(key)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) IncrBy(key string, value int64) *redis.IntCmd {
	if !m.hasStub("IncrBy") {
		return m.client.IncrBy(key, value)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) IncrByFloat(key string, value float64) *redis.FloatCmd {
	if !m.hasStub("IncrByFloat") {
		return m.client.IncrByFloat(key, value)
	}

	return m.Called().Get(0).(*redis.FloatCmd)
}

func (m *ClientMock) MGet(keys ...string) *redis.SliceCmd {
	if !m.hasStub("MGet") {
		return m.client.MGet()
	}

	return m.Called().Get(0).(*redis.SliceCmd)
}

func (m *ClientMock) MSet(pairs ...interface{}) *redis.StatusCmd {
	if !m.hasStub("MSet") {
		return m.client.MSet()
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) MSetNX(pairs ...interface{}) *redis.BoolCmd {
	if !m.hasStub("MSetNX") {
		return m.client.MSetNX()
	}

	return m.Called().Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	if !m.hasStub("Set") {
		return m.client.Set(key, value, expiration)
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) SetBit(key string, offset int64, value int) *redis.IntCmd {
	if !m.hasStub("SetBit") {
		return m.client.SetBit(key, offset, value)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) SetNX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	if !m.hasStub("SetNX") {
		return m.client.SetNX(key, value, expiration)
	}

	return m.Called().Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) SetXX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	if !m.hasStub("SetXX") {
		return m.client.SetXX(key, value, expiration)
	}

	return m.Called().Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) SetRange(key string, offset int64, value string) *redis.IntCmd {
	if !m.hasStub("SetRange") {
		return m.client.SetRange(key, offset, value)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) StrLen(key string) *redis.IntCmd {
	if !m.hasStub("StrLen") {
		return m.client.StrLen(key)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) BLPop(timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	if !m.hasStub("BLPop") {
		return m.client.BLPop(timeout, keys...)
	}

	return m.Called().Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) BRPop(timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	if !m.hasStub("BRPop") {
		return m.client.BRPop(timeout, keys...)
	}

	return m.Called().Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) BRPopLPush(source, destination string, timeout time.Duration) *redis.StringCmd {
	if !m.hasStub("BRPopLPush") {
		return m.client.BRPopLPush(source, destination, timeout)
	}

	return m.Called().Get(0).(*redis.StringCmd)
}

func (m *ClientMock) LIndex(key string, index int64) *redis.StringCmd {
	if !m.hasStub("LIndex") {
		return m.client.LIndex(key, index)
	}

	return m.Called().Get(0).(*redis.StringCmd)
}

func (m *ClientMock) LInsert(key, op string, pivot, value interface{}) *redis.IntCmd {
	if !m.hasStub("LInsert") {
		return m.client.LInsert(key, op, pivot, value)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) LInsertBefore(key string, pivot, value interface{}) *redis.IntCmd {
	if !m.hasStub("LInsertBefore") {
		return m.client.LInsertBefore(key, pivot, value)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) LInsertAfter(key string, pivot, value interface{}) *redis.IntCmd {
	if !m.hasStub("LInsertAfter") {
		return m.client.LInsertAfter(key, pivot, value)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) LLen(key string) *redis.IntCmd {
	if !m.hasStub("LLen") {
		return m.client.LLen(key)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) LPop(key string) *redis.StringCmd {
	if !m.hasStub("LPop") {
		return m.client.LPop(key)
	}

	return m.Called().Get(0).(*redis.StringCmd)
}

func (m *ClientMock) LPush(key string, values ...interface{}) *redis.IntCmd {
	if !m.hasStub("LPush") {
		return m.client.LPush(key, values...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) LPushX(key string, value interface{}) *redis.IntCmd {
	if !m.hasStub("LPushX") {
		return m.client.LPushX(key, value)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) LRange(key string, start, stop int64) *redis.StringSliceCmd {
	if !m.hasStub("LRange") {
		return m.client.LRange(key, start, stop)
	}

	return m.Called().Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) LRem(key string, count int64, value interface{}) *redis.IntCmd {
	if !m.hasStub("LRem") {
		return m.client.LRem(key, count, value)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) LSet(key string, index int64, value interface{}) *redis.StatusCmd {
	if !m.hasStub("LSet") {
		return m.client.LSet(key, index, value)
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) LTrim(key string, start, stop int64) *redis.StatusCmd {
	if !m.hasStub("LTrim") {
		return m.client.LTrim(key, start, stop)
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) RPop(key string) *redis.StringCmd {
	if !m.hasStub("RPop") {
		return m.client.RPop(key)
	}

	return m.Called().Get(0).(*redis.StringCmd)
}

func (m *ClientMock) RPopLPush(source, destination string) *redis.StringCmd {
	if !m.hasStub("RPopLPush") {
		return m.client.RPopLPush(source, destination)
	}

	return m.Called().Get(0).(*redis.StringCmd)
}

func (m *ClientMock) RPush(key string, values ...interface{}) *redis.IntCmd {
	if !m.hasStub("RPush") {
		return m.client.RPush(key, values...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) RPushX(key string, value interface{}) *redis.IntCmd {
	if !m.hasStub("RPushX") {
		return m.client.RPushX(key, value)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) PFAdd(key string, els ...interface{}) *redis.IntCmd {
	if !m.hasStub("PFAdd") {
		return m.client.PFAdd(key, els...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) PFCount(keys ...string) *redis.IntCmd {
	if !m.hasStub("PFCount") {
		return m.client.PFCount()
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) PFMerge(dest string, keys ...string) *redis.StatusCmd {
	if !m.hasStub("PFMerge") {
		return m.client.PFMerge(dest, keys...)
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) BgRewriteAOF() *redis.StatusCmd {
	if !m.hasStub("BgRewriteAOF") {
		return m.client.BgRewriteAOF()
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) BgSave() *redis.StatusCmd {
	if !m.hasStub("BgSave") {
		return m.client.BgSave()
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClientKill(ipPort string) *redis.StatusCmd {
	if !m.hasStub("ClientKill") {
		return m.client.ClientKill(ipPort)
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClientList() *redis.StringCmd {
	if !m.hasStub("ClientList") {
		return m.client.ClientList()
	}

	return m.Called().Get(0).(*redis.StringCmd)
}

func (m *ClientMock) ClientPause(dur time.Duration) *redis.BoolCmd {
	if !m.hasStub("ClientPause") {
		return m.client.ClientPause(dur)
	}

	return m.Called().Get(0).(*redis.BoolCmd)
}

func (m *ClientMock) ConfigGet(parameter string) *redis.SliceCmd {
	if !m.hasStub("ConfigGet") {
		return m.client.ConfigGet(parameter)
	}

	return m.Called().Get(0).(*redis.SliceCmd)
}

func (m *ClientMock) ConfigResetStat() *redis.StatusCmd {
	if !m.hasStub("ConfigResetStat") {
		return m.client.ConfigResetStat()
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ConfigSet(parameter, value string) *redis.StatusCmd {
	if !m.hasStub("ConfigSet") {
		return m.client.ConfigSet(parameter, value)
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ConfigRewrite() *redis.StatusCmd {
	if !m.hasStub("ConfigRewrite") {
		return m.client.ConfigRewrite()
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) DBSize() *redis.IntCmd {
	if !m.hasStub("DBSize") {
		return m.client.DBSize()
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) FlushAll() *redis.StatusCmd {
	if !m.hasStub("FlushAll") {
		return m.client.FlushAll()
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) FlushAllAsync() *redis.StatusCmd {
	if !m.hasStub("FlushAllAsync") {
		return m.client.FlushAllAsync()
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) FlushDB() *redis.StatusCmd {
	if !m.hasStub("FlushDB") {
		return m.client.FlushDB()
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) FlushDBAsync() *redis.StatusCmd {
	if !m.hasStub("FlushDBAsync") {
		return m.client.FlushDBAsync()
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) Info(section ...string) *redis.StringCmd {
	if !m.hasStub("Info") {
		return m.client.Info()
	}

	return m.Called().Get(0).(*redis.StringCmd)
}

func (m *ClientMock) LastSave() *redis.IntCmd {
	if !m.hasStub("LastSave") {
		return m.client.LastSave()
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) Save() *redis.StatusCmd {
	if !m.hasStub("Save") {
		return m.client.Save()
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) Shutdown() *redis.StatusCmd {
	if !m.hasStub("Shutdown") {
		return m.client.Shutdown()
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ShutdownSave() *redis.StatusCmd {
	if !m.hasStub("ShutdownSave") {
		return m.client.ShutdownSave()
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ShutdownNoSave() *redis.StatusCmd {
	if !m.hasStub("ShutdownNoSave") {
		return m.client.ShutdownNoSave()
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) SlaveOf(host, port string) *redis.StatusCmd {
	if !m.hasStub("SlaveOf") {
		return m.client.SlaveOf(host, port)
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) Time() *redis.TimeCmd {
	if !m.hasStub("Time") {
		return m.client.Time()
	}

	return m.Called().Get(0).(*redis.TimeCmd)
}

func (m *ClientMock) Eval(script string, keys []string, args ...interface{}) *redis.Cmd {
	if !m.hasStub("Eval") {
		return m.client.Eval(script, keys, args...)
	}

	return m.Called().Get(0).(*redis.Cmd)
}

func (m *ClientMock) EvalSha(sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	if !m.hasStub("EvalSha") {
		return m.client.EvalSha(sha1, keys, args...)
	}

	return m.Called().Get(0).(*redis.Cmd)
}

func (m *ClientMock) ScriptExists(hashes ...string) *redis.BoolSliceCmd {
	if !m.hasStub("ScriptExists") {
		return m.client.ScriptExists()
	}

	return m.Called().Get(0).(*redis.BoolSliceCmd)
}

func (m *ClientMock) ScriptFlush() *redis.StatusCmd {
	if !m.hasStub("ScriptFlush") {
		return m.client.ScriptFlush()
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ScriptKill() *redis.StatusCmd {
	if !m.hasStub("ScriptKill") {
		return m.client.ScriptKill()
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ScriptLoad(script string) *redis.StringCmd {
	if !m.hasStub("ScriptLoad") {
		return m.client.ScriptLoad(script)
	}

	return m.Called().Get(0).(*redis.StringCmd)
}

func (m *ClientMock) DebugObject(key string) *redis.StringCmd {
	if !m.hasStub("DebugObject") {
		return m.client.DebugObject(key)
	}

	return m.Called().Get(0).(*redis.StringCmd)
}

func (m *ClientMock) Publish(channel string, message interface{}) *redis.IntCmd {
	if !m.hasStub("Publish") {
		return m.client.Publish(channel, message)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) PubSubChannels(pattern string) *redis.StringSliceCmd {
	if !m.hasStub("PubSubChannels") {
		return m.client.PubSubChannels(pattern)
	}

	return m.Called().Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) PubSubNumSub(channels ...string) *redis.StringIntMapCmd {
	if !m.hasStub("PubSubNumSub") {
		return m.client.PubSubNumSub()
	}

	return m.Called().Get(0).(*redis.StringIntMapCmd)
}

func (m *ClientMock) PubSubNumPat() *redis.IntCmd {
	if !m.hasStub("PubSubNumPat") {
		return m.client.PubSubNumPat()
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ClusterSlots() *redis.ClusterSlotsCmd {
	if !m.hasStub("ClusterSlots") {
		return m.client.ClusterSlots()
	}

	return m.Called().Get(0).(*redis.ClusterSlotsCmd)
}

func (m *ClientMock) ClusterNodes() *redis.StringCmd {
	if !m.hasStub("ClusterNodes") {
		return m.client.ClusterNodes()
	}

	return m.Called().Get(0).(*redis.StringCmd)
}

func (m *ClientMock) ClusterMeet(host, port string) *redis.StatusCmd {
	if !m.hasStub("ClusterMeet") {
		return m.client.ClusterMeet(host, port)
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterForget(nodeID string) *redis.StatusCmd {
	if !m.hasStub("ClusterForget") {
		return m.client.ClusterForget(nodeID)
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterReplicate(nodeID string) *redis.StatusCmd {
	if !m.hasStub("ClusterReplicate") {
		return m.client.ClusterReplicate(nodeID)
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterResetSoft() *redis.StatusCmd {
	if !m.hasStub("ClusterResetSoft") {
		return m.client.ClusterResetSoft()
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterResetHard() *redis.StatusCmd {
	if !m.hasStub("ClusterResetHard") {
		return m.client.ClusterResetHard()
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterInfo() *redis.StringCmd {
	if !m.hasStub("ClusterInfo") {
		return m.client.ClusterInfo()
	}

	return m.Called().Get(0).(*redis.StringCmd)
}

func (m *ClientMock) ClusterKeySlot(key string) *redis.IntCmd {
	if !m.hasStub("ClusterKeySlot") {
		return m.client.ClusterKeySlot(key)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ClusterCountFailureReports(nodeID string) *redis.IntCmd {
	if !m.hasStub("ClusterCountFailureReports") {
		return m.client.ClusterCountFailureReports(nodeID)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ClusterCountKeysInSlot(slot int) *redis.IntCmd {
	if !m.hasStub("ClusterCountKeysInSlot") {
		return m.client.ClusterCountKeysInSlot(slot)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ClusterDelSlots(slots ...int) *redis.StatusCmd {
	if !m.hasStub("ClusterDelSlots") {
		return m.client.ClusterDelSlots(slots...)
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterDelSlotsRange(min, max int) *redis.StatusCmd {
	if !m.hasStub("ClusterDelSlotsRange") {
		return m.client.ClusterDelSlotsRange(min, max)
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterSaveConfig() *redis.StatusCmd {
	if !m.hasStub("ClusterSaveConfig") {
		return m.client.ClusterSaveConfig()
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterSlaves(nodeID string) *redis.StringSliceCmd {
	if !m.hasStub("ClusterSlaves") {
		return m.client.ClusterSlaves(nodeID)
	}

	return m.Called().Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) ClusterFailover() *redis.StatusCmd {
	if !m.hasStub("ClusterFailover") {
		return m.client.ClusterFailover()
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterAddSlots(slots ...int) *redis.StatusCmd {
	if !m.hasStub("ClusterAddSlots") {
		return m.client.ClusterAddSlots()
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClusterAddSlotsRange(min, max int) *redis.StatusCmd {
	if !m.hasStub("ClusterAddSlotsRange") {
		return m.client.ClusterAddSlotsRange(min, max)
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) GeoAdd(key string, geoLocation ...*redis.GeoLocation) *redis.IntCmd {
	if !m.hasStub("GeoAdd") {
		return m.client.GeoAdd(key, geoLocation...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) GeoPos(key string, members ...string) *redis.GeoPosCmd {
	if !m.hasStub("GeoPos") {
		return m.client.GeoPos(key, members...)
	}

	return m.Called().Get(0).(*redis.GeoPosCmd)
}

func (m *ClientMock) GeoRadius(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	if !m.hasStub("GeoRadius") {
		return m.client.GeoRadius(key, longitude, latitude, query)
	}

	return m.Called().Get(0).(*redis.GeoLocationCmd)
}

func (m *ClientMock) GeoRadiusRO(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	if !m.hasStub("GeoRadiusRO") {
		return m.client.GeoRadiusRO(key, longitude, latitude, query)
	}

	return m.Called().Get(0).(*redis.GeoLocationCmd)
}

func (m *ClientMock) GeoRadiusByMember(key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	if !m.hasStub("GeoRadiusByMember") {
		return m.client.GeoRadiusByMember(key, member, query)
	}

	return m.Called().Get(0).(*redis.GeoLocationCmd)
}

func (m *ClientMock) GeoRadiusByMemberRO(key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	if !m.hasStub("GeoRadiusByMemberRO") {
		return m.client.GeoRadiusByMemberRO(key, member, query)
	}

	return m.Called().Get(0).(*redis.GeoLocationCmd)
}

func (m *ClientMock) GeoDist(key string, member1, member2, unit string) *redis.FloatCmd {
	if !m.hasStub("GeoDist") {
		return m.client.GeoDist(key, member1, member2, unit)
	}

	return m.Called().Get(0).(*redis.FloatCmd)
}

func (m *ClientMock) GeoHash(key string, members ...string) *redis.StringSliceCmd {
	if !m.hasStub("GeoHash") {
		return m.client.GeoHash(key, members...)
	}

	return m.Called().Get(0).(*redis.StringSliceCmd)
}

func (m *ClientMock) Command() *redis.CommandsInfoCmd {
	if !m.hasStub("Command") {
		return m.client.Command()
	}

	return m.Called().Get(0).(*redis.CommandsInfoCmd)
}

func (m *ClientMock) ClientKillByFilter(keys ...string) *redis.IntCmd {
	if !m.hasStub("ClientKillByFilter") {
		return m.client.ClientKillByFilter(keys...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) MemoryUsage(key string, samples ...int) *redis.IntCmd {
	if !m.hasStub("MemoryUsage") {
		return m.client.MemoryUsage(key, samples...)
	}

	return m.Called().Get(0).(*redis.IntCmd)
}

func (m *ClientMock) ReadOnly() *redis.StatusCmd {
	if !m.hasStub("ReadOnly") {
		return m.client.ReadOnly()
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ReadWrite() *redis.StatusCmd {
	if !m.hasStub("ReadWrite") {
		return m.client.ReadWrite()
	}

	return m.Called().Get(0).(*redis.StatusCmd)
}

func (m *ClientMock) ClientID() *redis.IntCmd {
	if !m.hasStub("ClientID") {
		return m.client.ClientID()
	}

	return m.Called().Get(0).(*redis.IntCmd)
}
