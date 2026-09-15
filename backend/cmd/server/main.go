// Package main is the entry point of the Remember Me backend server.
//
// Phase 4 — Backend foundation (structure-only).
// This file only wires the basic HTTP server, config loader, and router.
// Business features (auth, task CRUD, etc.) are intentionally NOT implemented.
package main

import (
	"log"

	"remember_me/config"
	"remember_me/database"
	"remember_me/router"
)

func main() {
	// Load .env (if present) and configuration from environment variables.
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Initialize database connection (Phase 5 will make this real).
	db, err := database.NewConnection(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Build router.
	r := router.New(db)

	// Start HTTP server.
	addr := ":" + cfg.ServerPort
	log.Printf("Remember Me backend starting on %s (env=%s)", addr, cfg.AppEnv)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
