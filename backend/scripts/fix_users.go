//go:build ignore

// Script to drop and recreate the users table.
// Usage: go run scripts/fix_users.go
package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
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

	// Drop existing users table and related objects
	fmt.Println("Dropping users table...")
	_, err = pool.Exec(ctx, "DROP TABLE IF EXISTS users CASCADE")
	if err != nil {
		fmt.Fprintf(os.Stderr, "drop failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("  OK: dropped users table")

	// Drop existing trigger function
	_, err = pool.Exec(ctx, "DROP FUNCTION IF EXISTS update_updated_at_column() CASCADE")
	if err != nil {
		fmt.Fprintf(os.Stderr, "drop function failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("  OK: dropped trigger function")

	// Re-run migration
	content, err := os.ReadFile("../database/migrations/001_create_users.sql")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read 001: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Running 001_create_users.sql...")
	_, err = pool.Exec(ctx, string(content))
	if err != nil {
		fmt.Fprintf(os.Stderr, "create failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("  OK: users table created")

	fmt.Println("\nFix complete.")
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
