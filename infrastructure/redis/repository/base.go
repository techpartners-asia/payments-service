package repositoryRedis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepository interface {
	Get(key string, out any) error
	Set(key string, value interface{}, expiration time.Duration) error
	Delete(key string) error
}

type redisRepository struct {
	redis *redis.Client
	ctx   context.Context
}

func NewRedisRepository(redis *redis.Client) RedisRepository {
	return &redisRepository{redis: redis, ctx: context.Background()}
}

func (r *redisRepository) Get(key string, out any) error {
	data, err := r.redis.Get(r.ctx, key).Result()
	if err != nil && err != redis.Nil {
		return err
	}

	if err == redis.Nil {
		return nil
	}

	return json.Unmarshal([]byte(data), &out)
}
func (r *redisRepository) Set(key string, value interface{}, expiration time.Duration) error {
	return r.redis.Set(r.ctx, key, value, expiration).Err()
}

func (r *redisRepository) Delete(key string) error {
	return r.redis.Del(r.ctx, key).Err()
}
