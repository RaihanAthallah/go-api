package io

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/hayvee-website-development/go-api-hayvee/config"
)

var sentryLog *SentryLog
var ctx = context.Background()

func initCache() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			config.C.Redis.Host,
			config.C.Redis.Port),
		Username: config.C.Redis.User,
		Password: config.C.Redis.Password,
		DB:       0,
	})

	if err := client.Ping(context.TODO()).Err(); err != nil {
		sentryLog.CaptureError(err)
		return nil, err
	}

	return client, nil
}

func GetCache(key string, obj interface{}) error {
	conn, err := initCache()
	if err != nil {
		return err
	}

	data, err := conn.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(data), &obj)
	if err != nil {
		return err
	}
	return nil
}

func SetCache(key string, exp int, value interface{}) error {
	conn, err := initCache()
	if err != nil {
		return err
	}

	_, err = conn.Set(ctx, key, value, time.Duration(exp)*time.Second).Result()
	if err != nil {
		return err
	}
	return nil
}
