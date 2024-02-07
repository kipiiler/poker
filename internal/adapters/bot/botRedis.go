package adapters

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type BotCache struct {
	client *redis.Client
}

func NewBotCache(client *redis.Client) *BotCache {
	return &BotCache{client: client}
}

func (b *BotCache) CheckKeyExists(key string) (bool, error) {
	// Use the Exists method to check if a key exists
	exists, err := b.client.Exists(context.Background(), key).Result()
	if err != nil {
		return false, err
	}

	return exists > 0, nil
}

func (b *BotCache) AddKey(key string, value string) error {
	// Bot value expire after 1 day
	expiration := time.Duration(86400) * time.Second
	err := b.client.Set(context.Background(), key, value, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (b *BotCache) AddKeyWithExpiration(key string, value string, expiration int) error {
	// Bot value expire after 1 day
	expirationTime := time.Duration(expiration) * time.Second
	err := b.client.Set(context.Background(), key, value, expirationTime).Err()
	if err != nil {
		return err
	}

	return nil
}

func (b *BotCache) RemoveKey(key string) error {
	// Use the Del method to remove a key
	err := b.client.Del(context.Background(), key).Err()
	if err != nil {
		return err
	}

	return nil
}

func (b *BotCache) GetKey(key string) (string, error) {
	// Use the Get method to retrieve a key
	value, err := b.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}

	return value, nil
}

func (b *BotCache) GetKeysWithPrefix(prefix string) ([]string, error) {
	// Use the Keys method to retrieve all keys with a given prefix
	keys, err := b.client.Keys(context.Background(), prefix).Result()
	if err != nil {
		return nil, err
	}

	return keys, nil
}
