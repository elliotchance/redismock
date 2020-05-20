# redismock

Package `github.com/elliotchance/redismock` is useful for unit testing
applications that interact with Redis. It uses the
[`stretchr/testify` mocking system](https://github.com/stretchr/testify).

Unlike using a real or fake Redis (more below), `redismock` provides normal and
nice mocks for more flexibility in control behavior. One example might be to
test the case where Redis takes too long to respond to a command, or returns an
expected error.

### NewMock()

Creates a hollow mock. You will need to stub all commands that you will interact
with.

This is most useful when you want to strictly test all Redis interactions.

### NewNiceMock(client)

Create a mocks that falls back to the real `client` in cases where a command has
not been stubbed.

This is most useful when you want a real Redis instance, but you need to stub
off certain commands or behaviors.

## Compatibility

### github.com/go-redis/redis

`go-redis/redis` is a real Redis client, and one of the most popular. It 
expects to connect to a real Redis server.

It provides a huge interface for all of the Redis commands called `Cmdable`. You
should use `Cmdable` interface throughout your code, instead of the struct type.

### github.com/alicebob/miniredis

`miniredis` creates a fake Redis server that can be used by `go-redis/redis`
which makes it fantastic for running unit tests since you do not need to run a
separate Redis server. It also only retains state with the instance, so each
test gets an empty Redis database each time.

## Example

```go
import (
	"testing"
	"github.com/go-redis/redis"
	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock"
	"errors"
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

// Test Redis is down.
func TestRedisCannotBePinged(t *testing.T) {
	r := newTestRedis()
	r.On("Ping").
		Return(redis.NewStatusResult("", errors.New("server not available")))

	assert.False(t, RedisIsAvailable(r))
}
```
