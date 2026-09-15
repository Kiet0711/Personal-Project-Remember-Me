// Package main is the entry point of the Remember Me backend server.
//
// Phase 6 — Authentication.
package main

import (
	"log"
	"strconv"

	"remember_me/config"
	"remember_me/database"
	"remember_me/internal/auth"
	"remember_me/internal/user"
	"remember_me/router"
)

func main() {
	// Load .env (if present) and configuration from environment variables.
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Initialize database connection (Phase 5).
	db, err := database.NewConnection(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize JWT secret (Phase 6).
	auth.SetJWTSecret(cfg.JWTSecret)

	// Parse JWT expiry (default 2h).
	expiresIn, _ := strconv.Atoi(cfg.JWTExpiresIn)
	if expiresIn <= 0 {
		expiresIn = 7200
	}

	// Build auth service + handler (Phase 6).
	authRepo := auth.NewRepository(db)
	authSvc := auth.NewService(authRepo, expiresIn)
	authHandler := auth.NewHandler(authSvc)

	// Build user service + handler (Phase 7 stub).
	userSvc := user.NewService(authRepo)
	userHandler := user.NewHandler(userSvc)

	// Build router.
	r := router.New(authHandler, userHandler)

	// Start HTTP server.
	addr := ":" + cfg.ServerPort
	log.Printf("Remember Me backend starting on %s (env=%s)", addr, cfg.AppEnv)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
