package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config{
		port: ":8080",
		db:   dbConfig{},
	}

	api := api{
		config: cfg,
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	if err := api.run(ctx, api.mount()); err != nil {
		log.Fatalln("Server error: ", err)
	}
}
