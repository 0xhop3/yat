package handlers

import (
	"github.com/0xhop3/yat/backend/internal/models"
	"github.com/0xhop3/yat/backend/internal/services"
	"github.com/0xhop3/yat/backend/internal/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

func NewHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	user, err := u.userService.CreateUser(&req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to create user", err)
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "User created successfully", user)
}
