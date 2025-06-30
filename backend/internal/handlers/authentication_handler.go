package handlers

import (
	"net/http"

	"github.com/0xhop3/yat/backend/internal/models"
	"github.com/0xhop3/yat/backend/internal/services"
	"github.com/0xhop3/yat/backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuthenticationHandler struct {
	userService *services.UserService
}

func NewAuthenticationHandler(userService *services.UserService) *AuthenticationHandler {
	return &AuthenticationHandler{
		userService: userService,
	}
}

func (a *AuthenticationHandler) GetProfile(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "User not found in context", nil)
		return
	}

	userModel := user.(*models.User)
	utils.SuccessResponse(c, http.StatusOK, "Profile retrieved successfully", userModel)
}

func (a *AuthenticationHandler) Callback(c *gin.Context) {
	utils.SuccessResponse(c, http.StatusOK, "Authentication callback processed", nil)
}
