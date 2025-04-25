package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/baseapp/application/dto"
	"github.com/baseapp/application/usecase"
)

// UserController handles user-related HTTP requests
type UserController struct {
	userUseCase *usecase.UserUseCase
	authController *AuthController
}

// NewUserController creates a new UserController
func NewUserController(userUseCase *usecase.UserUseCase, authController *AuthController) *UserController {
	return &UserController{
		userUseCase: userUseCase,
		authController: authController,
	}
}

// GetProfile handles the get profile route
func (c *UserController) GetProfile(ctx *gin.Context) {
	// Get the user ID from the token
	firebaseUID, err := c.authController.GetUserIDFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Get the profile
	profile, err := c.userUseCase.GetProfile(ctx.Request.Context(), firebaseUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Return the profile
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"profile": profile,
	})
}

// UpdateProfile handles the update profile route
func (c *UserController) UpdateProfile(ctx *gin.Context) {
	// Get the user ID from the token
	firebaseUID, err := c.authController.GetUserIDFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Parse request body
	var requestBody dto.UserUpdateRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "無効なリクエストボディ: " + err.Error(),
		})
		return
	}

	// Update the profile
	profile, err := c.userUseCase.UpdateProfile(ctx.Request.Context(), firebaseUID, &requestBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Return the updated profile
	ctx.JSON(http.StatusOK, gin.H{
		"success":  true,
		"profile":  profile,
		"message":  "プロフィールが更新されました",
	})
}
