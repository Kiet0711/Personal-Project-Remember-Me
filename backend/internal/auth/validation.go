// Package auth handles registration, login, logout, and password hashing.
//
// Phase 6 — Authentication.
package auth

import (
	"regexp"
	"strings"
	"time"
)

// ValidateRegister validates a registration request.
// Returns a list of validation errors (empty if valid).
func ValidateRegister(req *RegisterRequest) []string {
	var errs []string

	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	req.Phone = strings.TrimSpace(req.Phone)
	req.Password = strings.TrimSpace(req.Password)
	req.Timezone = strings.TrimSpace(req.Timezone)

	if req.Name == "" {
		errs = append(errs, "name is required")
	} else if len(req.Name) < 2 {
		errs = append(errs, "name must be at least 2 characters")
	} else if len(req.Name) > 255 {
		errs = append(errs, "name must not exceed 255 characters")
	}

	if req.Email == "" {
		errs = append(errs, "email is required")
	} else if !isValidEmail(req.Email) {
		errs = append(errs, "email is invalid")
	}

	if req.Password == "" {
		errs = append(errs, "password is required")
	} else if len(req.Password) < 8 {
		errs = append(errs, "password must be at least 8 characters")
	} else if len(req.Password) > 128 {
		errs = append(errs, "password must not exceed 128 characters")
	}

	// Timezone is optional; default to Asia/Ho_Chi_Minh
	if req.Timezone == "" {
		req.Timezone = "Asia/Ho_Chi_Minh"
	} else if !isValidTimezone(req.Timezone) {
		errs = append(errs, "timezone is invalid")
	}

	return errs
}

// ValidateLogin validates a login request.
// Returns a list of validation errors (empty if valid).
func ValidateLogin(req *LoginRequest) []string {
	var errs []string

	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	req.Password = strings.TrimSpace(req.Password)

	if req.Email == "" {
		errs = append(errs, "email is required")
	} else if !isValidEmail(req.Email) {
		errs = append(errs, "email is invalid")
	}

	if req.Password == "" {
		errs = append(errs, "password is required")
	}

	return errs
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func isValidTimezone(tz string) bool {
	_, err := time.LoadLocation(tz)
	return err == nil
}
