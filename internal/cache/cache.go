package cache

import (
	"context"
	"time"

	"github.com/alielmi98/go-weather-api/internal/config"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, value string, expiration time.Duration) error
}

type redisCache struct {
	client *redis.Client
}

func NewCache(cfg *config.Config) Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		logrus.Fatalf("failed to connect to Redis: %v", err)
	}

	return &redisCache{
		client: client,
	}
}

func (c *redisCache) Get(ctx context.Context, key string) (string, error) {
	value, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (c *redisCache) Set(ctx context.Context, key, value string, expiration time.Duration) error {
	err := c.client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}
