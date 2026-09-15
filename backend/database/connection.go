// Package database manages the PostgreSQL connection lifecycle.
//
// Phase 5 — Database connection.
// Uses pgxpool for connection pooling.
package database

import (
	"context"
	"fmt"
	"time"

	"remember_me/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DB wraps a pgx connection pool. All database operations go through this type.
type DB struct {
	Pool *pgxpool.Pool
}

// NewConnection establishes a connection pool to PostgreSQL and verifies
// connectivity by pinging the database.
//
// It reads DB_* configuration from cfg and expects a PostgreSQL server
// running at DB_HOST:DB_PORT with credentials DB_USER/DB_PASSWORD.
func NewConnection(cfg *config.Config) (*DB, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database DSN: %w", err)
	}

	// Reasonable pool defaults for a local dev setup.
	poolConfig.MaxConns = 20
	poolConfig.MinConns = 2
	poolConfig.MaxConnLifetime = time.Hour
	poolConfig.MaxConnIdleTime = 30 * time.Minute

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Verify connectivity.
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database %s@%s:%s/%s: %w",
			cfg.DBUser, cfg.DBHost, cfg.DBPort, cfg.DBName, err)
	}

	return &DB{Pool: pool}, nil
}

// Close releases all pool connections.
func (db *DB) Close() {
	if db.Pool != nil {
		db.Pool.Close()
	}
}
