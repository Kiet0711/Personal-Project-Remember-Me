// Package auth handles registration, login, logout, and password hashing.
//
// Phase 6 — Authentication.
package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims represents the claims stored in a JWT token.
type JWTClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// JWTSecret is the secret key used to sign JWT tokens.
// It is set once at startup from config.
var JWTSecret string

// SetJWTSecret initializes the JWT secret from config.
func SetJWTSecret(secret string) {
	JWTSecret = secret
}

// GenerateToken creates a new JWT token for a user.
// The token expires after expiresInSeconds seconds.
func GenerateToken(userID, email string, expiresInSeconds int) (string, error) {
	if JWTSecret == "" {
		return "", errors.New("JWT secret not configured")
	}

	expirationTime := time.Now().Add(time.Duration(expiresInSeconds) * time.Second)

	claims := &JWTClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "remember-me",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWTSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

// ValidateToken parses and validates a JWT token.
// Returns the claims if valid, or an error if invalid.
func ValidateToken(tokenString string) (*JWTClaims, error) {
	if JWTSecret == "" {
		return nil, errors.New("JWT secret not configured")
	}

	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWTSecret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
