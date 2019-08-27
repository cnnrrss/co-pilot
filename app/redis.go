package copilot

import (
	"time"

	"github.com/go-redis/redis"
)

var defaultTTL = time.Duration(time.Hour * 1)

// RedisCache implements the Cache interface
type RedisCache struct {
	*redis.Client
	ttl *time.Duration
}

// NewRedisCache is a constructor for the RedisCache type
func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{
		client,
		&defaultTTL,
	}
}
