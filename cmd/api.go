package main

import (
	"net/http"
	"time"
	"log"
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (api *api) mount() http.Handler { // mount returns an http.Handler, keeping the server framework-agnostic.

	// Router
	r := chi.NewRouter()

  	r.Use(middleware.RequestID)
  	r.Use(middleware.RealIP)
  	r.Use(middleware.Logger)
  	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
	    w.Write([]byte("OK"))
	})

	return r
}

func (api *api) run(ctx context.Context, h http.Handler) error {
	srv := &http.Server{
		Addr:              api.config.port,
		Handler:           h,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Println("Starting server on port", api.config.port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErrors <- err
		}
	}()

	select {
	case <-ctx.Done(): // Blocks the main goroutine until the application context gets cancelled.
		log.Println("Shutdown signal received!")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second) // Create a context to attempt a graceful 30 second shutdown.
		defer cancel() // Releases context resources once the srv.Shutdown() call is complete.

		return srv.Shutdown(shutdownCtx) // Shutdown stops accepting new connections and waits for in-flight requests to finish.

	case err := <-serverErrors:
		return err
	}
}


type api struct {
	config config
}

type config struct {
	port string
	db dbConfig
}

type dbConfig struct {
	dsn string
}