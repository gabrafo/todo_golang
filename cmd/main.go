package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gabrafo/todo_golang/internal/env"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	cfg := config{
		port: env.MustEnv("APP_PORT"),
		db: dbConfig{
			dsn: env.MustEnv("DB_URL"),
		},
	}

	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		slog.Error("Unable to connect to database: ", "error", err)
		os.Exit(1)
	}
	slog.Info("Connected to database")
	defer conn.Close(ctx)

	api := api{
		config: cfg,
		logger: logger,
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	if err := api.run(ctx, api.mount()); err != nil {
		slog.Error("Server error: ", "error", err)
		os.Exit(1)
	}
}