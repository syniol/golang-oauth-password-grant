package cache

import "context"

type Cache interface {
	Persist(ctx context.Context, key, value string) error
	LookUp(ctx context.Context, key string) (string, error)
}

func NewCache() (Cache, error) {
	return newRedisClient(), nil
}
