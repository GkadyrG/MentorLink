package config

import (
	"log"
	postgres "review/internal/storage/db"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	postgres.Config
	Env string `env:"ENV" env-required:"true"`
}

func LoadConfig() *Config {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./.env", cfg)
	if err != nil {
		log.Fatalf("error reading config: %s", err.Error())
	}
	return cfg
}
