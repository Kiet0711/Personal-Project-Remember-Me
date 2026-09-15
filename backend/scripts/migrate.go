//go:build ignore

// Script to run all migration files against the database.
// Usage: go run scripts/migrate.go
package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	dbUser := getEnv("DB_USER", "postgres")
	dbPass := url.QueryEscape(getEnv("DB_PASSWORD", ""))
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbName := getEnv("DB_NAME", "RememberMe")

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		url.QueryEscape(dbUser),
		dbPass,
		dbHost,
		dbPort,
		url.QueryEscape(dbName),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to connect: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "failed to ping database: %v\n", err)
		os.Exit(1)
	}

	// Find all migration files
	migrationDir := filepath.Join("..", "database", "migrations")
	files, err := os.ReadDir(migrationDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read migrations dir: %v\n", err)
		os.Exit(1)
	}

	var sqlFiles []string
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".sql") {
			sqlFiles = append(sqlFiles, f.Name())
		}
	}
	sort.Strings(sqlFiles)

	for _, name := range sqlFiles {
		path := filepath.Join(migrationDir, name)
		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to read %s: %v\n", name, err)
			os.Exit(1)
		}

		fmt.Printf("Running: %s\n", name)
		_, err = pool.Exec(ctx, string(content))
		if err != nil {
			// IF NOT EXISTS suppresses errors for already-applied migrations
			fmt.Fprintf(os.Stderr, "  WARNING: %s: %v\n", name, err)
		} else {
			fmt.Printf("  OK: %s\n", name)
		}
	}

	fmt.Println("\nMigrations complete.")
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
