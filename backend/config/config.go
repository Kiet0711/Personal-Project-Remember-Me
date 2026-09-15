// Package config loads application configuration from environment variables.
//
// Phase 4 — Backend foundation. No business logic here.
package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all runtime configuration loaded from environment variables.
type Config struct {
	AppEnv    string
	AppDebug  bool
	AppURL    string

	ServerPort string

	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string

	JWTSecret    string
	JWTExpiresIn string

	SessionLifetime string
}

// Load reads the .env file (if present) and returns a populated Config.
//
// .env loading is best-effort: a missing file is not fatal because the
// process can still be configured via real environment variables in production.
func Load() (*Config, error) {
	// Best-effort: ignore "file not found" errors.
	_ = godotenv.Load()

	cfg := &Config{
		AppEnv:          getEnv("APP_ENV", "development"),
		AppDebug:        getEnvBool("APP_DEBUG", true),
		AppURL:          getEnv("APP_URL", "http://localhost:8080"),
		ServerPort:      getEnv("SERVER_PORT", "8080"),
		DBHost:          getEnv("DB_HOST", "localhost"),
		DBPort:          getEnv("DB_PORT", "5432"),
		DBName:          getEnv("DB_NAME", "remember_me"),
		DBUser:          getEnv("DB_USER", "postgres"),
		DBPassword:      os.Getenv("DB_PASSWORD"),
		JWTSecret:       os.Getenv("JWT_SECRET"),
		JWTExpiresIn:    getEnv("JWT_EXPIRES_IN", "7200"),
		SessionLifetime: getEnv("SESSION_LIFETIME", "7200"),
	}

	if cfg.JWTSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET is required (set it in backend/.env)")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v == "true" || v == "1"
}
