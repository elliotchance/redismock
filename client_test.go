package redismock_test

import (
	"errors"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis/v7"
	"github.com/stretchr/testify/assert"
)

// newTestRedis returns a redis.Cmdable.
func newTestRedis() *redismock.ClientMock {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	return redismock.NewNiceMock(client)
}

// This would be your production code.
func RedisIsAvailable(client redis.Cmdable) bool {
	return client.Ping().Err() == nil
}

func RedisDeleteKeys(client redis.Cmdable) int64 {
	res, _ := client.Del("key1", "key2").Result()
	return res
}

func RedisMGetKeys(client redis.Cmdable) []interface{} {
	res, _ := client.MGet("key1", "key2").Result()
	return res
}

func RedisScan(client redis.Cmdable) []string {
	page, _, _ := client.Scan(0, "myredis.hash.*", 1000).Result()
	return page
}

// Test Redis is down.
func TestRedisCannotBePinged(t *testing.T) {
	r := newTestRedis()
	r.On("Ping").
		Return(redis.NewStatusResult("", errors.New("server not available")))

	assert.False(t, RedisIsAvailable(r))
}

func TestRedisDel(t *testing.T) {
	r := newTestRedis()
	r.On("Del", []string{"key1", "key2"}).
		Return(redis.NewIntResult(0, nil))

	assert.Equal(t, int64(0), RedisDeleteKeys(r))
}

func TestRedisMGet(t *testing.T) {
	r := newTestRedis()
	strs := []string{"value1", "value2"}
	values := make([]interface{}, len(strs))
	r.On("Mget", []string{"key1", "key2"}).
		Return(redis.NewSliceResult(values, nil))

	assert.Equal(t, values, RedisMGetKeys(r))
}

func TestRedisScan(t *testing.T) {
	r := newTestRedis()
	res := []string{"result1", "result2", "result3"}
	r.On("Scan", uint64(0), "myredis.hash.*", int64(1000)).
		Return(redis.NewScanCmdResult(res, 0, nil))

	assert.Equal(t, res, RedisScan(r))
}
