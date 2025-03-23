package config

import (
	"log"
	"review/internal/storage/cache"
	postgres "review/internal/storage/db"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	postgres.Config
	cache.RedisConfig
	Address       string        `env:"ADDRESS" env-required:"true"`
	Timeout       time.Duration `env:"TIMEOUT" env-default:"4s"`
	IdleTimeout   time.Duration `env:"IDLE_TIMEOUT" env-default:"60s"`
	Env           string        `env:"ENV" env-required:"true"`
	PublicKeyPath string        `env:"JWT_PUBLIC_KEY_PATH"`
}

func LoadConfig() *Config {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./.env", cfg)
	if err != nil {
		log.Fatalf("error reading config: %s", err.Error())
	}
	return cfg
}
