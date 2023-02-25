package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestRedis(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background()

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(pong)
}
