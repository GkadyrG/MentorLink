package main

import (
	"log/slog"
	"mentorlink/internal/config"
	"mentorlink/internal/lib/logger/sl"
	"mentorlink/internal/storage/cache"
	"mentorlink/internal/storage/db"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	cfg := config.LoadConfig()

	log := setupLogger(cfg.Env)

	log.Info("starting auth-service", slog.String("env", cfg.Env))

	redis := cache.New(cfg.RedisConfig)

	storage, err := db.NewStorage(cfg.Config)
	if err != nil {
		log.Error("error creation storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage
	_ = redis
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
