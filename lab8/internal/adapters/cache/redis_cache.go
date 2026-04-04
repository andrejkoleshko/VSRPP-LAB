package cache

import (
    "context"
    "time"

    "github.com/redis/go-redis/v9"
)

type RedisCache struct {
    client *redis.Client
}

func NewRedisCache(addr string) *RedisCache {
    return &RedisCache{
        client: redis.NewClient(&redis.Options{
            Addr: addr,
        }),
    }
}

func (c *RedisCache) Get(key string) (float32, bool) {
    ctx := context.Background()
    val, err := c.client.Get(ctx, key).Float32()
    if err != nil {
        return 0, false
    }
    return val, true
}

func (c *RedisCache) Set(key string, value float32) error {
    ctx := context.Background()
    return c.client.Set(ctx, key, value, 10*time.Minute).Err()
}