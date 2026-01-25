package main

import (
	"net/http"
	"time"
	"log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (api *api) mount() http.Handler { // http.Handler is agnostic, doesn't matter if we are using gin, echo, mux, etc.
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

func (api *api) run(h http.Handler) error {
	srv := &http.Server{
		Addr: api.config.port,
		Handler: h,
		ReadTimeout:       5 * time.Second,
    	ReadHeaderTimeout: 2 * time.Second,
    	WriteTimeout:      10 * time.Second,
   		IdleTimeout:       60 * time.Second,
	}

	log.Println("Starting server on port ", api.config.port)

	return srv.ListenAndServe()
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