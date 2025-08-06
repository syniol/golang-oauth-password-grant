package cache

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"log"
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

			return "cache:6379"
		}(),
		DB: 0, // use default DB
		TLSConfig: (func() *tls.Config {
			caCert, err := os.ReadFile(os.Getenv("REDIS_TLS_CA_FILE"))
			if err != nil {
				log.Fatal("error loading redis TLS CA Certificate", err)
			}

			caCertPool := x509.NewCertPool()
			caCertPool.AppendCertsFromPEM(caCert)

			cert, err := tls.LoadX509KeyPair(
				os.Getenv("REDIS_TLS_CERT_FILE"),
				os.Getenv("REDIS_TLS_KEY_FILE"),
			)
			if err != nil {
				log.Fatal("error loading redis TLS Key and Cert Pair Certificate", err)
			}

			return &tls.Config{
				RootCAs:            caCertPool,
				Certificates:       []tls.Certificate{cert},
				InsecureSkipVerify: true,
			}
		})(),
	})

	return &RedisAdapter{
		client: redisClient,
		ctx:    ctx,
	}
}
