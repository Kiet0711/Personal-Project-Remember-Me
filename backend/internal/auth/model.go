// Package auth handles registration, login, logout, and password hashing.
//
// Phase 6 — Authentication.
package auth

import "time"

// User represents a user in the system.
type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PhoneNumber  *string   `json:"phone_number,omitempty"`
	PasswordHash string    `json:"-"` // never expose in JSON
	Timezone     string    `json:"timezone"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
