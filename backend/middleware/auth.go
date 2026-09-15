// Package middleware contains HTTP middlewares used across the Remember Me API.
//
// Phase 6 — Authentication middleware.
package middleware

import (
	"net/http"
	"strings"

	"remember_me/internal/auth"

	"github.com/gin-gonic/gin"
)

// Auth returns a middleware that validates JWT tokens.
// It extracts the token from the Authorization header (Bearer <token>),
// validates it, and sets user_id and user_email in the Gin context.
// Returns 401 if the token is missing, invalid, or expired.
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			return
		}

		tokenString := parts[1]
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		// Set user info in context for handlers to use.
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)

		c.Next()
	}
}
