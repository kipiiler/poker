package adapters

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func NewRedisCache() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("HOST") + ":" + os.Getenv("REDIS_PORT_DEVELOPMENT"),
		Password: os.Getenv("REDIS_PASSWORD_DEVELOPMENT"),
		DB:       0,
	})

	return client
}
