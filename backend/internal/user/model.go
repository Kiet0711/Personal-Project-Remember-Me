// Package user handles user profile operations.
//
// Phase 6 — User stub (full implementation in Phase 7).
package user

import (
	"time"

	"remember_me/internal/auth"
)

// UserResponse mirrors auth.UserResponse for the public-facing user data.
type UserResponse = auth.UserResponse

// UpdateProfileRequest is the request body for PUT /api/users/me.
type UpdateProfileRequest struct {
	Name      string `json:"name" binding:"omitempty,min=2,max=255"`
	Phone     string `json:"phone_number" binding:"omitempty,max=50"`
	Timezone  string `json:"timezone"`
}

// ChangePasswordRequest is the request body for PUT /api/users/password.
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required,min=8,max=128"`
}

// ToResponse converts an auth.User to a UserResponse.
func ToResponse(u *auth.User) UserResponse {
	return UserResponse{
		ID:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		Timezone:    u.Timezone,
		CreatedAt:   u.CreatedAt.Format(time.RFC3339),
	}
}
