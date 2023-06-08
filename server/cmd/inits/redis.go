package inits

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"

	"github.com/bagasunix/kuyngetrip/pkg/env"
)

func InitRedis(ctx context.Context, configs *env.Configs) *redis.Client {
	host := configs.RedisHost
	port := configs.RedisPort
	password := configs.RedisPassword
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprint(host, ":", port),
		Password: password, // no password set
		DB:       0,        // use default DB
	}).WithContext(ctx)
}
