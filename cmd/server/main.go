package main

import (
	"log/slog"
	"net/http"
	"os"

	"runemaster/internal/api"
	"runemaster/internal/config"
	"runemaster/internal/db"
	"runemaster/internal/logger"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		slog.Error("Failed to load config", slog.String("error", err.Error()))
		os.Exit(1)
	}

	logger.Init(cfg.Environment)

	dbPool, err := db.Connect(cfg.DatabaseURL)

	if err != nil {
		slog.Error("Failed to connect to database", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer dbPool.Close()

	router := api.NewRouter(dbPool)

	slog.Info("Starting server",
		slog.String("mode", cfg.Environment),
		slog.String("port", cfg.ServerPort),
	)

	if err := http.ListenAndServe(":"+cfg.ServerPort, router); err != nil {
		slog.Error("Server error", slog.String("error", err.Error()))
	}
}
