// Package router wires the HTTP routes and middleware for the Remember Me API.
//
// Phase 4 — Backend foundation (structure-only).
// Feature routes (tasks, reminders, etc.) will be added in their respective phases.
package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"remember_me/database"
	"remember_me/middleware"
)

// New builds and returns a configured *gin.Engine.
func New(db *database.DB) *gin.Engine {
	if os.Getenv("APP_DEBUG") == "false" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

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
		// Feature routes will be mounted here in later phases.
		// Placeholder so the variable is not unused.
		api.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Remember Me API"})
		})

		// Example (do NOT uncomment yet — Phase 9+):
		// api.POST("/users/register", authHandler.Register)
		// api.GET("/tasks", taskHandler.List)
	}

	return r
}
