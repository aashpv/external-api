package app

import (
	"external-api/internal/config"
	"external-api/internal/database/postgres"
	"external-api/internal/transport/rest/handlers/info"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
)

func Run() {
	cfg := config.MustLoad()

	fmt.Println("application started")

	db, err := postgres.New(cfg.Database)
	if err != nil {
		fmt.Errorf("failed connected databse:%w\n", err)
		return
	}
	fmt.Println("database connected")

	router := chi.NewRouter()

	// Middlewares
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Get("/info", info.New(db))

	address := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	fmt.Println("starting server", slog.String("address", address))

	srv := &http.Server{
		Addr:         address,
		Handler:      router,
		ReadTimeout:  cfg.Server.Timeout,
		WriteTimeout: cfg.Server.Timeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Errorf("error starting server: %w\n", err)
	}

	fmt.Println("server stopped")
}
