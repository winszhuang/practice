package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

const Key = "userCombo"

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	var (
		userId   = "17345"
		gameName = "poker"
	)

	field := fmt.Sprintf("%s_%s", userId, gameName)
	cmd := rdb.HSet(ctx, Key, field, 0)
	if err := cmd.Err(); err != nil {
		panic(err)
	}

	result, err := rdb.HGet(ctx, Key, field).Result()
	if err != nil {
		panic(err)
	}

	println(result)
}
