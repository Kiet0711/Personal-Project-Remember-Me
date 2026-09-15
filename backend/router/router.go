// Package router wires the HTTP routes and middleware for the Remember Me API.
//
// Phase 6 — Authentication (auth + user routes wired).
package router

import (
	"net/http"

	"remember_me/internal/auth"
	"remember_me/internal/user"
	"remember_me/middleware"

	"github.com/gin-gonic/gin"
)

// New builds and returns a configured *gin.Engine.
func New(authHandler *auth.Handler, userHandler *user.Handler) *gin.Engine {
	r := gin.New()

	// Global middleware
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())
	r.Use(middleware.CORS())

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"app":    "remember-me-backend",
		})
	})

	// API base prefix
	api := r.Group("/api")
	{
		// ── Auth routes (public) ──────────────────────────────────
		api.POST("/users/register", authHandler.Register)
		api.POST("/users/login", authHandler.Login)

		// ── Protected routes (require JWT) ──────────────────────────
		protected := api.Group("")
		protected.Use(middleware.Auth())
		{
			// Auth
			protected.POST("/users/logout", authHandler.Logout)

			// User profile
			protected.GET("/users/me", userHandler.GetMe)
			protected.PUT("/users/me", userHandler.UpdateProfile)
			protected.PUT("/users/password", userHandler.ChangePassword)
		}
	}

	return r
}
