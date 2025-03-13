package config

import (
	"log"
	"mentorlink/internal/storage/cache"
	postgres "mentorlink/internal/storage/db"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	postgres.Config
	cache.RedisConfig

	Env string `env:"ENV" env-required:"true"`

	AccessTokenTTL  int64 `env:"ACCES_TOKEN_TTL" env-default:"900"`
	RefreshTokenTTL int64 `env:"RESRESH_TOKEN_TTL" env-default:"604800"`

	PrivateKeyPath string `env:"JWT_PRIVATE_KEY_PATH"`
	PublicKeyPath  string `env:"JWT_PUBLIC_KEY_PATH"`
}

func LoadConfig() *Config {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./.env", cfg)
	if err != nil {
		log.Fatal("error reading config: %s", err.Error())
	}
	return cfg
}
