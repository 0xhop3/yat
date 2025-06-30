package utils

import (
	"github.com/0xhop3/yat/backend/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetUserFromContext extracts user from gin context
func GetUserFromContext(c *gin.Context) (*models.User, bool) {
	user, exists := c.Get("user")
	if !exists {
		return nil, false
	}

	userModel, ok := user.(*models.User)
	return userModel, ok
}

// GetUserIDFromContext extracts user ID from gin context
func GetUserIDFromContext(c *gin.Context) (uuid.UUID, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		return uuid.Nil, false
	}

	id, ok := userID.(uuid.UUID)
	return id, ok
}

// GetAuth0IDFromContext extracts Auth0 ID from gin context
func GetAuth0IDFromContext(c *gin.Context) (string, bool) {
	auth0ID, exists := c.Get("auth0_id")
	if !exists {
		return "", false
	}

	id, ok := auth0ID.(string)
	return id, ok
}
