// Package auth handles registration, login, logout, and password hashing.
//
// Phase 6 — Authentication.
package auth

import (
	"context"
	"errors"
	"fmt"

	"remember_me/database"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ErrUserNotFound is returned when a user does not exist.
var ErrUserNotFound = errors.New("user not found")

// ErrEmailExists is returned when trying to register with an existing email.
var ErrEmailExists = errors.New("email already exists")

// Repository provides database operations for auth.
type Repository struct {
	pool *pgxpool.Pool
}

// NewRepository creates a new auth Repository.
func NewRepository(db *database.DB) *Repository {
	return &Repository{pool: db.Pool}
}

// Create inserts a new user into the database.
// Returns the created user or an error.
func (r *Repository) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (id, name, email, phone_number, password_hash, timezone)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING created_at, updated_at
	`

	err := r.pool.QueryRow(ctx, query,
		user.ID,
		user.Name,
		user.Email,
		user.PhoneNumber,
		user.PasswordHash,
		user.Timezone,
	).Scan(&user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		// Check for unique violation
		if isUniqueViolation(err) {
			return ErrEmailExists
		}
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

// GetByEmail retrieves a user by email address.
// Returns ErrUserNotFound if no user matches.
func (r *Repository) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, name, email, phone_number, password_hash, timezone, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	user := &User{}
	err := r.pool.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PhoneNumber,
		&user.PasswordHash,
		&user.Timezone,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	return user, nil
}

// GetByID retrieves a user by ID.
// Returns ErrUserNotFound if no user matches.
func (r *Repository) GetByID(ctx context.Context, id string) (*User, error) {
	query := `
		SELECT id, name, email, phone_number, password_hash, timezone, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	user := &User{}
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PhoneNumber,
		&user.PasswordHash,
		&user.Timezone,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	return user, nil
}

// isUniqueViolation checks if a PostgreSQL error is a unique constraint violation.
func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == "23505"
	}
	return false
}
