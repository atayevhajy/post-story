package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var redisClient *redis.Client
var ctx = context.TODO()

func initRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error connecting to redis: %v", err)
	}
}
