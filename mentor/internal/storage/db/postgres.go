package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	UserName string `env:"POSTGRES_USER" env-required:"true"`
	Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
	Host     string `env:"POSTGRES_HOST" env-required:"true"`
	Port     string `env:"POSTGRES_PORT" env-required:"true"`
	DBName   string `env:"POSTGRES_DB" env-required:"true"`
}

type Storage struct {
	db *sqlx.DB
}

func NewStorage(cfg Config) (*Storage, error) {
	dsn := fmt.Sprintf("host=%s port=%s  user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.UserName, cfg.Password, cfg.DBName)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error with connect to database: %w", err)
	}

	return &Storage{db: db}, nil
}
