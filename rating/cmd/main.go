package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"rating/internal/kafka"
	"rating/internal/repository"
	"syscall"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("starting rating server")

	repo := repository.NewRepository()

	ctx, cancel := context.WithCancel(context.Background())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	consumer, err := kafka.NewConsumer([]string{"localhost:9092"}, "my-group", repo, logger)
	if err != nil {
		logger.Error("failed to create consumer", "error", err)
		os.Exit(1)
	}

	go consumer.Run(ctx, "reviews-topic")

	select {
	case <-sigChan:
		logger.Info("signal caught, shutting down gracefulle")
	case <-ctx.Done():
		logger.Info("context cancelled, shutting down gracefully...")
	}

	cancel()

	closeErr := consumer.Close()
	if closeErr != nil {
		logger.Error("error while closing consumer", "error", closeErr)
	}

	logger.Info("consumer service stopped")
}
