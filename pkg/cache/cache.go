package cache

import (
	"context"
	"time"
)

type Cache interface {
	PersistWithTimeToLive(ctx context.Context, key, value string, ttl time.Duration) error
	Persist(ctx context.Context, key, value string) error
	LookUp(ctx context.Context, key string) (string, error)
}

func NewCache() (Cache, error) {
	return newRedisClient(), nil
}
