package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/baseapp/middleware"
	"github.com/baseapp/models"
	"github.com/baseapp/services"
)

// RegisterProfileRoutes registers the profile routes
func RegisterProfileRoutes(router *gin.Engine) {
	profileGroup := router.Group("/api")
	{
		profileGroup.GET("/profile", middleware.AuthRequired(), GetProfileHandler)
		profileGroup.PUT("/profile", middleware.AuthRequired(), UpdateProfileHandler)
	}
}

// GetProfileHandler handles the get profile route
func GetProfileHandler(c *gin.Context) {
	// Get the user ID from the token
	firebaseUID, err := middleware.GetUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Find the user profile
	var profile models.UserProfile
	result := services.DB.Where("firebase_uid = ?", firebaseUID).First(&profile)

	if result.Error != nil {
		// Profile not found, create a new one
		profile = models.UserProfile{
			FirebaseUID: firebaseUID,
		}
		result = services.DB.Create(&profile)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "プロフィールの作成に失敗しました: " + result.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success":  true,
			"profile":  profile.ToMap(),
			"message":  "プロフィールが作成されました",
		})
		return
	}

	// Return the profile
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"profile": profile.ToMap(),
	})
}

// UpdateProfileHandler handles the update profile route
func UpdateProfileHandler(c *gin.Context) {
	// Get the user ID from the token
	firebaseUID, err := middleware.GetUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Parse request body
	var requestBody struct {
		DisplayName string `json:"display_name"`
		Bio         string `json:"bio"`
		Location    string `json:"location"`
		Website     string `json:"website"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "無効なリクエストボディ: " + err.Error(),
		})
		return
	}

	// Find the user profile
	var profile models.UserProfile
	result := services.DB.Where("firebase_uid = ?", firebaseUID).First(&profile)

	if result.Error != nil {
		// Profile not found, create a new one
		profile = models.UserProfile{
			FirebaseUID: firebaseUID,
			DisplayName: requestBody.DisplayName,
			Bio:         requestBody.Bio,
			Location:    requestBody.Location,
			Website:     requestBody.Website,
		}
		result = services.DB.Create(&profile)
	} else {
		// Update the profile
		if requestBody.DisplayName != "" {
			profile.DisplayName = requestBody.DisplayName
		}
		if requestBody.Bio != "" {
			profile.Bio = requestBody.Bio
		}
		if requestBody.Location != "" {
			profile.Location = requestBody.Location
		}
		if requestBody.Website != "" {
			profile.Website = requestBody.Website
		}
		result = services.DB.Save(&profile)
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "プロフィールの更新に失敗しました: " + result.Error.Error(),
		})
		return
	}

	// Return the updated profile
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"profile":  profile.ToMap(),
		"message":  "プロフィールが更新されました",
	})
}
