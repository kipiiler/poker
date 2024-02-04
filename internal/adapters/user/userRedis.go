package adapters

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type UserCache struct {
	client *redis.Client
}

func NewUserCache(client *redis.Client) *UserCache {
	return &UserCache{client: client}
}

func (u *UserCache) CheckKeyExists(key string) (bool, error) {
	// Use the Exists method to check if a key exists
	exists, err := u.client.Exists(context.Background(), key).Result()
	if err != nil {
		return false, err
	}

	return exists > 0, nil
}

func (u *UserCache) AddKey(key string, expire int64) error {
	// Use the Set method to add a key with expiration time
	expiration := time.Duration(expire) * time.Second
	err := u.client.Set(context.Background(), key, "some_value", expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (u *UserCache) RemoveKey(key string) error {
	// Use the Del method to remove a key
	err := u.client.Del(context.Background(), key).Err()
	if err != nil {
		return err
	}

	return nil
}
