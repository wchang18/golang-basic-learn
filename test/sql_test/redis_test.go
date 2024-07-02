package sql

import (
	"context"
	"fmt"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	rds := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:16379",
		Password: "123456",
	})

	ctx := context.Background()
	res := rds.SetEX(ctx, "chang_name", "gary", time.Second*60)
	fmt.Println(res.Val())
}

func TestRedisMock(t *testing.T) {
	ms, _ := miniredis.Run()
	defer ms.Close()
	ms.Set("my_name", "gary")

	rds := redis.NewClient(&redis.Options{
		Addr: ms.Addr(),
	})
	ctx := context.Background()
	val := rds.Get(ctx, "my_name")
	assert.Equal(t, "gary", val.Val())
}
