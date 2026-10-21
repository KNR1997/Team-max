// Package main runs the todo API server.
package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"todo/internal/bootstrap"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(log)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg := bootstrap.Config{
		Addr:        env("HTTP_ADDR", ":8080"),
		DatabaseURL: env("DATABASE_URL", "postgres://todo:todo@localhost:5433/todo?sslmode=disable"),
	}

	srv := bootstrap.New(cfg, log)
	if err := srv.Run(ctx); err != nil {
		log.Error("fatal", "error", err)
		os.Exit(1)
	}
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
