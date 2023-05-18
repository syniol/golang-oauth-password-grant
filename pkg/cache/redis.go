package cache

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisAdapter struct {
	client *redis.Client
	ctx    context.Context
}

func (ra *RedisAdapter) Persist(key, value string) error {
	res := ra.client.Set(ra.ctx, key, value, time.Hour)
	if res.Err() != nil {
		return res.Err()
	}

	return nil
}

func (ra *RedisAdapter) LookUp(key string) (string, error) {
	res := ra.client.Get(ra.ctx, key)
	if res.Err() != nil {
		return "", res.Err()
	}

	return res.Val(), nil
}

func newRedisClient(ctx context.Context) Cache {
	redisClient := redis.NewClient(&redis.Options{
		Addr: func() string {
			if len(os.Getenv("DEBUG")) > 0 {
				return "127.0.0.1:6379"

			}

			return "host.docker.internal:6379"
		}(),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &RedisAdapter{
		client: redisClient,
		ctx:    ctx,
	}
}
