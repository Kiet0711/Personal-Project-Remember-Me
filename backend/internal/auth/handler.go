// Package auth handles registration, login, logout, and password hashing.
//
// Phase 6 — Authentication.
package auth

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler handles HTTP requests for auth endpoints.
type Handler struct {
	svc *Service
}

// NewHandler creates a new auth Handler.
func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

// Register handles POST /api/users/register.
func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request body: " + err.Error()})
		return
	}

	resp, err := h.svc.Register(c.Request.Context(), &req)
	if err != nil {
		var ve *ValidationError
		if errors.As(err, &ve) {
			c.JSON(http.StatusUnprocessableEntity, ErrorResponse{Error: ve.Error()})
			return
		}
		var se *ServiceError
		if errors.As(err, &se) {
			log.Printf("register error: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: se.Error()})
			return
		}
		log.Printf("register unexpected error: %v", err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// Login handles POST /api/users/login.
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request body: " + err.Error()})
		return
	}

	resp, err := h.svc.Login(c.Request.Context(), &req)
	if err != nil {
		var ve *ValidationError
		if errors.As(err, &ve) {
			c.JSON(http.StatusUnprocessableEntity, ErrorResponse{Error: ve.Error()})
			return
		}
		var se *ServiceError
		if errors.As(err, &se) {
			log.Printf("login error: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: se.Error()})
			return
		}
		log.Printf("login unexpected error: %v", err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "internal server error"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Logout handles POST /api/users/logout.
func (h *Handler) Logout(c *gin.Context) {
	// Get user ID from context (set by auth middleware).
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "unauthorized"})
		return
	}

	if err := h.svc.Logout(c.Request.Context(), userID.(string)); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "failed to logout"})
		return
	}

	c.JSON(http.StatusOK, MessageResponse{Message: "logged out successfully"})
}
