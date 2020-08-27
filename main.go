package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	fmt.Println("redis server ")
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	err := rdb.Set(ctx, "kj", "redis master", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "kj").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("kj: val= ", val)

	val2, err := rdb.Get(ctx, "mj").Result()
	if err != nil {
		fmt.Println("No value")
	} else {
		fmt.Println(val2)
	}

	ListSet(rdb, ctx, "list")
	ListGet(rdb, ctx, "list")

}

func ListSet(clint *redis.Client, ctx context.Context, key string) {
	listVal := []string{"a", "b"}
	err := clint.RPush(ctx, key, listVal).Err()
	if err != nil {
		panic(err)
	}
}

func ListGet(client *redis.Client, ctx context.Context, key string) {
	val, err := client.LRange(ctx, key, 0, 1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
