package main

import (
	"fmt"
	"log"
	"log/slog"
	"mentorlink/internal/config"
	postgres "mentorlink/pkg/db"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	cfg, err := config.New()

	if err != nil {
		log.Fatal("Config error")
	}

	logger := setupLogger(cfg.Env)

	storage, err = postgres.NewStorage(cfg.Config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_ = storage
	_ = logger

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
