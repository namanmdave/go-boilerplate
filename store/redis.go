package store

import (
	"context"
	"go-boilerplate/config"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis(cfg *config.RedisConfig) (*redis.Client, error) {
	if redisClient != nil {
		return redisClient, nil
	}

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	_, err := client.Ping(context.TODO()).Result()
	if err != nil {
		return nil, err
	}

	redisClient = client
	return redisClient, nil
}
