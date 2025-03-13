package main

import (
	"log/slog"
	"mentorlink/internal/config"
	"mentorlink/internal/handlers"
	"mentorlink/internal/lib/logger/sl"
	"mentorlink/internal/storage/cache"
	"mentorlink/internal/storage/db"
	"mentorlink/internal/token"
	"net/http"
	"os"

	"github.com/go-chi/chi"
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

	redisClient := cache.New(cfg.RedisConfig)

	redisRepository := cache.NewRedisRepository(redisClient)

	storage, err := db.NewStorage(cfg.Config)
	if err != nil {
		log.Error("error creation storage", sl.Err(err))
		os.Exit(1)
	}

	tokemMn, err := token.NewTokenmanagerRSA(cfg.PrivateKeyPath, cfg.PublicKeyPath)
	if err != nil {
		log.Error("error with token manager", sl.Err(err))
		os.Exit(1)
	}

	router := chi.NewRouter()
	router.Post("/auth/register", handlers.Register(log, storage))
	router.Post("/auth/login", handlers.Login(log, storage, tokemMn))
	router.Post("/auth/logout", handlers.Logout(log, redisRepository, tokemMn))

	err = http.ListenAndServe("localhost:8081", router)
	log.Error(err.Error())
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
