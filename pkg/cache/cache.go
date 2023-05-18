package cache

import "context"

type Cache interface {
	Persist(key, value string) error
	LookUp(key string) (string, error)
}

func NewCache(ctx context.Context) (Cache, error) {
	if ctx == nil {
		return newRedisClient(context.Background()), nil
	}

	return newRedisClient(ctx), nil
}
