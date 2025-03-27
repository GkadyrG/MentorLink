package cache

import (
	"fmt"

	"github.com/go-redis/redis"
)

type RedisConfig struct {
	Host     string `env:"REDIS_HOST" env-required:"true"`
	Port     string `env:"REDIS_PORT" env-required:"true"`
	Password string `env:"REDIS_PASSWORD" env-required:"true"`
}

func New(cfg RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
	})
	if err := client.Ping().Err(); err != nil {
		panic(fmt.Errorf("failed to connect Redis:%w", err))
	}
	return client
}

type RedisRepository struct {
	Client *redis.Client
}
