package main

import (
	"log"
	"net/http"

	"runemaster/internal/api"
	"runemaster/internal/config"
	"runemaster/internal/db"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	dbPool, err := db.Connect(cfg.DatabaseURL)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbPool.Close()

	router := api.NewRouter(dbPool)

	log.Printf("Starting server in %s mode on port %s", cfg.Environment, cfg.ServerPort)

	err = http.ListenAndServe(":"+cfg.ServerPort, router)

	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
