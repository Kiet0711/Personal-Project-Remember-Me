// Package auth handles registration, login, logout, and password hashing.
//
// Phase 6 — Authentication.
package auth

import "time"

// RegisterRequest is the request body for POST /api/users/register.
type RegisterRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=255"`
	Email    string `json:"email" binding:"required,email,max=255"`
	Phone    string `json:"phone_number" binding:"omitempty,max=50"`
	Password string `json:"password" binding:"required,min=8,max=128"`
	Timezone string `json:"timezone"`
}

// RegisterResponse is the response body for POST /api/users/register.
type RegisterResponse struct {
	User      UserResponse `json:"user"`
	Token     string      `json:"token"`
	ExpiresIn string      `json:"expires_in"`
}

// LoginRequest is the request body for POST /api/users/login.
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse is the response body for POST /api/users/login.
type LoginResponse struct {
	User      UserResponse `json:"user"`
	Token     string      `json:"token"`
	ExpiresIn string      `json:"expires_in"`
}

// UserResponse is the public-facing user data (no password).
type UserResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	Timezone    string  `json:"timezone"`
	CreatedAt   string  `json:"created_at"`
}

// MessageResponse is used for simple message responses (e.g., logout).
type MessageResponse struct {
	Message string `json:"message"`
}

// ErrorResponse is used for all error responses.
type ErrorResponse struct {
	Error string `json:"error"`
}

// ToResponse converts a User to a UserResponse.
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		Timezone:    u.Timezone,
		CreatedAt:   u.CreatedAt.Format(time.RFC3339),
	}
}
