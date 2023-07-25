package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisLogStore struct {
	rdb          *redis.Client
	logHashKey   string
	logIDListKey string
}

func newRedisLogStore(c *redis.Client) *redisLogStore {
	return &redisLogStore{c, "logs", "logids"}
}

func (rs *redisLogStore) GetLogByID(ctx context.Context, id string) (*logRecord, error) {
	log, err := rs.rdb.HGet(ctx, rs.logHashKey, id).Result()
	if err != nil {
		return nil, err
	}

	r := &logRecord{}
	err = json.Unmarshal([]byte(log), r)
	if err != nil {
		return nil, err
	}

	return r, nil

}

func (rs *redisLogStore) SaveLog(ctx context.Context, l *logRecord) error {
	logID := fmt.Sprintf("log:%s", l.ID.String())

	log, err := json.Marshal(l)
	if err != nil {
		return err
	}

	err = rs.rdb.HSet(ctx, rs.logHashKey, logID, log).Err()
	if err != nil {
		return err
	}

	err = rs.rdb.LPush(ctx, rs.logIDListKey, logID).Err()
	if err != nil {
		return err
	}

	return nil
}

func (rs *redisLogStore) FindLogs(ctx context.Context, filter logFilter) ([]*logRecord, error) {
	result := make([]*logRecord, 0, filter.limit)

	end := filter.offset + filter.limit - 1
	ids, err := rs.rdb.LRange(ctx, rs.logIDListKey, int64(filter.offset), int64(end)).Result()
	if err != nil {
		return nil, err
	}

	logs, err := rs.rdb.HMGet(ctx, rs.logHashKey, ids...).Result()
	if err != nil {
		return nil, err
	}

	for i, v := range logs {
		r := &logRecord{}
		logString := v.(string)
		if err := json.Unmarshal([]byte(logString), r); err != nil {
			return nil, fmt.Errorf("error: failed unmarshal at index %d", i)
		}
		result = append(result, r)
	}

	return result, nil
}

// rediss://default:b53dd65dcb914e96969ab06c5ca625d1@upward-tahr-33126.upstash.io:33126
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
