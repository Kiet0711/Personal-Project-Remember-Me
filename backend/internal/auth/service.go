// Package auth handles registration, login, logout, and password hashing.
//
// Phase 6 — Authentication.
package auth

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

// Service provides authentication business logic.
type Service struct {
	repo          *Repository
	jwtExpiresSec int
	timezone      string
}

// NewService creates a new auth Service.
func NewService(repo *Repository, jwtExpiresSec int) *Service {
	return &Service{
		repo:          repo,
		jwtExpiresSec: jwtExpiresSec,
		timezone:      "Asia/Ho_Chi_Minh",
	}
}

// Register creates a new user account.
// Returns the created user, a JWT token, and nil error on success.
// Returns an error on failure.
func (s *Service) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	// Validate input
	if errs := ValidateRegister(req); len(errs) > 0 {
		return nil, &ValidationError{Errors: errs}
	}

	// Hash password
	hash, err := HashPassword(req.Password)
	if err != nil {
		return nil, &ServiceError{Message: "failed to hash password"}
	}

	// Build user
	user := &User{
		ID:           uuid.New().String(),
		Name:         req.Name,
		Email:        req.Email,
		PhoneNumber:  nullString(req.Phone),
		PasswordHash: hash,
		Timezone:     req.Timezone,
	}

	// Create user in DB
	if err := s.repo.Create(ctx, user); err != nil {
		if errors.Is(err, ErrEmailExists) {
			return nil, &ValidationError{Errors: []string{"email already registered"}}
		}
		return nil, &ServiceError{Message: "failed to create user: " + err.Error()}
	}

	// Generate JWT
	token, err := GenerateToken(user.ID, user.Email, s.jwtExpiresSec)
	if err != nil {
		return nil, &ServiceError{Message: "failed to generate token"}
	}

	return &RegisterResponse{
		User:      user.ToResponse(),
		Token:     token,
		ExpiresIn: formatDuration(s.jwtExpiresSec),
	}, nil
}

// Login authenticates a user with email and password.
// Returns the user data, a JWT token, and nil error on success.
// Returns an error on failure.
func (s *Service) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	// Validate input
	if errs := ValidateLogin(req); len(errs) > 0 {
		return nil, &ValidationError{Errors: errs}
	}

	// Find user by email
	user, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return nil, &ValidationError{Errors: []string{"invalid email or password"}}
		}
		return nil, &ServiceError{Message: "failed to find user: " + err.Error()}
	}

	// Verify password
	if !CheckPassword(req.Password, user.PasswordHash) {
		return nil, &ValidationError{Errors: []string{"invalid email or password"}}
	}

	// Generate JWT
	token, err := GenerateToken(user.ID, user.Email, s.jwtExpiresSec)
	if err != nil {
		return nil, &ServiceError{Message: "failed to generate token"}
	}

	return &LoginResponse{
		User:      user.ToResponse(),
		Token:     token,
		ExpiresIn: formatDuration(s.jwtExpiresSec),
	}, nil
}

// Logout is a no-op for stateless JWT authentication.
// The client simply discards the token.
func (s *Service) Logout(ctx context.Context, userID string) error {
	// Stateless JWT — nothing to invalidate server-side.
	// The client removes the token from storage.
	return nil
}

// ValidationError is returned when input validation fails.
type ValidationError struct {
	Errors []string
}

func (e *ValidationError) Error() string {
	return "validation error: " + joinStrings(e.Errors, "; ")
}

// ServiceError is returned for internal service errors.
type ServiceError struct {
	Message string
}

func (e *ServiceError) Error() string {
	return e.Message
}

func nullString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func formatDuration(seconds int) string {
	hours := seconds / 3600
	if hours >= 24 {
		days := hours / 24
		return formatInt(days) + "d"
	}
	if hours > 0 {
		return formatInt(hours) + "h"
	}
	return formatInt(seconds) + "s"
}

func formatInt(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	for n > 0 {
		result = string(rune('0'+n%10)) + result
		n /= 10
	}
	return result
}

func joinStrings(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	result := strs[0]
	for i := 1; i < len(strs); i++ {
		result += sep + strs[i]
	}
	return result
}
