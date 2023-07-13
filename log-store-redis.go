package main

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisLogStore struct {
	rdb *redis.Client
}

func (rs *redisLogStore) Get(ctx context.Context, id string) (record, error) {
	return record{}, nil
}

func (rs *redisLogStore) Write(ctx context.Context, r record) error {
	return nil
}

func (rs *redisLogStore) All(ctx context.Context) ([]record, error) {
	return nil, nil
}

func (rs *redisLogStore) Recent(ctx context.Context) ([]record, error) {
	return nil, nil
}

func openRedis(redisURL string) (*redis.Client, error) {
	config, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}

	c := redis.NewClient(config)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = c.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return c, nil
}
