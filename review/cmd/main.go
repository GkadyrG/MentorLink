package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"review/internal/config"
	"review/internal/handlers/create"
	del "review/internal/handlers/delete"
	"review/internal/handlers/get"
	"review/internal/handlers/update"
	"review/internal/lib/logger/sl"
	mwAuth "review/internal/middleware/auth"
	mwLogger "review/internal/middleware/logger"
	"review/internal/storage/cache"
	"review/internal/storage/db"
	"review/pkg/token"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	cfg := config.LoadConfig()

	log := setupLogger(cfg.Env)

	log.Info(
		"starting review-service",
		slog.String("env", cfg.Env),
	)

	log.Debug("debug messages are enabled")

	storage, err := db.NewStorage(cfg.Config)
	if err != nil {
		log.Error("error created storage", sl.Err(err))
		os.Exit(1)
	}

	redisClient := cache.New(cfg.RedisConfig)

	redisRepository := cache.NewRedisRepository(redisClient)

	tokenMn, err := token.NewTokenManagerRSA(cfg.PublicKeyPath)
	if err != nil {
		log.Error("error created a new token manager", sl.Err(err))
		os.Exit(1)
	}
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.URLFormat)
	router.Use(mwLogger.New(log))

	router.Group(func(r chi.Router) {
		r.Use(mwAuth.AuthMiddleware(tokenMn, log))
		r.Post("/review/create", create.Create(log, storage))
		r.Put("/review/update", update.Update(log, storage))
		r.Delete("/review/delete/{id}", del.Delete(log, storage))

	})

	router.Get("/review/get", get.Get(log, storage, redisRepository))

	log.Info("starting server", slog.String("adsress", cfg.Address))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("failed to start server", sl.Err(err))
		}
	}()

	log.Info("server started")

	<-done
	log.Info("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", sl.Err(err))

		return
	}

	log.Info("server stopped")

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
