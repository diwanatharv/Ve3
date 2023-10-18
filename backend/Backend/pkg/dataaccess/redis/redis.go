package redis

import (
	"awesomeProject/pkg/config"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type rediss struct {
	client *redis.Client
}

func Redismanager() *rediss {
	return &rediss{
		client: config.Makeredisserever(),
	}
}

type RedisMethod interface {
	Setredis(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Getredis(ctx context.Context, key string) *redis.StringCmd
	Deletekey(ctx context.Context, keys ...string) *redis.IntCmd
}

func (r *rediss) Setredis(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.client.Set(ctx, key, value, expiration)
}
func (r *rediss) Getredis(ctx context.Context, key string) *redis.StringCmd {
	return r.client.Get(ctx, key)
}
func (r *rediss) Deletekey(ctx context.Context, keys ...string) *redis.IntCmd {
	return r.client.Del(ctx, keys...)
}
