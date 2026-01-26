package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	cfg := config{
		port: ":8080",
		db:   dbConfig{},
	}

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
