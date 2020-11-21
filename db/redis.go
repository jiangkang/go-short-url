package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
)

var (
	RedisDB *redis.Client
)

func init() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func GetDbCount(ctx context.Context) int {
	RedisDB.Incr(ctx, "counter")
	idString, _ := RedisDB.Get(ctx, "counter").Result()
	id, _ := strconv.Atoi(idString)
	return id
}
