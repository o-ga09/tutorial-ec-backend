package redis

import (
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

func NewRedisClient(cfg config.Redis) *redis.Client {
	client = redis.NewClient(&redis.Options{
		Addr: cfg.Host + ":" + cfg.Port ,
		ReadTimeout: 3 * time.Minute,
		WriteTimeout: 3 * time.Minute,
		ContextTimeoutEnabled: true,
	})
	return client
}