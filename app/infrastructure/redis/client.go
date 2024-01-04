package redis

import (
	"context"
	"time"

	"github.com/o-ga09/tutorial-ec-backend/app/config"
	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
)

func GetRedisClient() *redis.Client {
	return client
}

func SetRedisClient(c *redis.Client) {
	client = c
}

func NewRedisClient(ctx context.Context) *redis.Client {
	cfg := config.GetConfig()
	client = redis.NewClient(&redis.Options{
		Addr: cfg.RedisURL ,
		ReadTimeout: 3 * time.Minute,
		WriteTimeout: 3 * time.Minute,
		ContextTimeoutEnabled: true,
	})
	return client
}