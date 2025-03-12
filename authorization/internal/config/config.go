package config

import (
	"fmt"
	"mentorlink/pkg/cache"
	postgres "mentorlink/pkg/db"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	postgres.Config
	cache.RedisConfig
	Env string `env:"ENV" env-required:"true"`
}

func New() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("../.env", cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}
	return cfg, nil
}
