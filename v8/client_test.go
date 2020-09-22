package redismock_test

import (
	"context"
	"errors"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock/v8"
	"github.com/go-redis/redis/v8"
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
func RedisIsAvailable(ctx context.Context, client redis.Cmdable) bool {
	return client.Ping(ctx).Err() == nil
}

func RedisDeleteKeys(ctx context.Context, client redis.Cmdable) int64 {
	res, _ := client.Del(ctx, "key1", "key2").Result()
	return res
}

func RedisMGetKeys(ctx context.Context, client redis.Cmdable) []interface{} {
	res, _ := client.MGet(ctx, "key1", "key2").Result()
	return res
}

func RedisScan(ctx context.Context, client redis.Cmdable) []string {
	page, _, _ := client.Scan(ctx, 0, "myredis.hash.*", 1000).Result()
	return page
}

// Test Redis is down.
func TestRedisCannotBePinged(t *testing.T) {
	r := newTestRedis()
	r.On("Ping", context.Background()).
		Return(redis.NewStatusResult("", errors.New("server not available")))

	assert.False(t, RedisIsAvailable(context.Background(), r))
}

func TestRedisDel(t *testing.T) {
	r := newTestRedis()
	r.On("Del", context.Background(), []string{"key1", "key2"}).
		Return(redis.NewIntResult(0, nil))

	assert.Equal(t, int64(0), RedisDeleteKeys(context.Background(), r))
}

func TestRedisMGet(t *testing.T) {
	r := newTestRedis()
	strs := []string{"value1", "value2"}
	values := make([]interface{}, len(strs))
	r.On("Mget", []string{"key1", "key2"}).
		Return(redis.NewSliceResult(values, nil))

	assert.Equal(t, values, RedisMGetKeys(context.Background(), r))
}

func TestRedisScan(t *testing.T) {
	r := newTestRedis()
	res := []string{"result1", "result2", "result3"}
	r.On("Scan", context.Background(), uint64(0), "myredis.hash.*", int64(1000)).
		Return(redis.NewScanCmdResult(res, 0, nil))

	assert.Equal(t, res, RedisScan(context.Background(), r))
}
