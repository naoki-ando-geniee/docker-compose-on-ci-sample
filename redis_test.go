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

func TestRedisCluster(t *testing.T) {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"localhost:7001", "localhost:7002", "localhost:7003", "localhost:7004", "localhost:7005", "localhost:7006"},
	})
	ctx := context.Background()
	err := client.ForEachShard(ctx, func(ctx context.Context, client *redis.Client) error {
		pong, err := client.Ping(ctx).Result()
		if err != nil {
			return err
		}
		fmt.Println(pong)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	err = client.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		t.Fatal(err)
	}
}
