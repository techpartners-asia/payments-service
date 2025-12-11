package redisService

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
	configPkg "github.com/techpartners-asia/payments-service/pkg/config"
)

var (
	Redis *redis.Client

	once sync.Once
)

func Init() {
	once.Do(func() {
		redisClient := redis.NewClient(
			&redis.Options{
				Addr:     fmt.Sprintf("%s:%s", configPkg.Env.Redis.Host, configPkg.Env.Redis.Port),
				Password: configPkg.Env.Redis.Password,
				DB:       configPkg.Env.Redis.DB,
			},
		)

		if err := redisClient.Ping(context.Background()).Err(); err != nil {
			log.Fatalf("Failed to connect to Redis: %v", err)
			panic(err)
		}

		Redis = redisClient
	})
}
