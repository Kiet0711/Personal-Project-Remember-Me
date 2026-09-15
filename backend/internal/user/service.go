// Package user handles user profile operations.
//
// Phase 6 — User stub (full implementation in Phase 7).
package user

import (
	"context"

	"remember_me/internal/auth"
)

// Service provides user profile business logic.
type Service struct {
	authRepo *auth.Repository
}

// NewService creates a new user Service.
func NewService(authRepo *auth.Repository) *Service {
	return &Service{authRepo: authRepo}
}

// GetMe retrieves the current user's profile.
func (s *Service) GetMe(ctx context.Context, userID string) (*UserResponse, error) {
	user, err := s.authRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	resp := ToResponse(user)
	return &resp, nil
}

// UpdateProfile updates the current user's name, phone, and timezone.
func (s *Service) UpdateProfile(ctx context.Context, userID string, req *UpdateProfileRequest) (*UserResponse, error) {
	// TODO: implement in Phase 7
	_ = req
	user, err := s.authRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	resp := ToResponse(user)
	return &resp, nil
}

// ChangePassword changes the current user's password.
// TODO: implement in Phase 7 (needs auth service bcrypt)
func (s *Service) ChangePassword(ctx context.Context, userID string, req *ChangePasswordRequest) error {
	// TODO: implement in Phase 7
	_ = ctx
	_ = userID
	_ = req
	return nil
}
