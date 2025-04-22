package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/baseapp/middleware"
)

// RegisterAuthRoutes registers the auth routes
func RegisterAuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("/verify", middleware.AuthRequired(), VerifyAuthHandler)
	}
}

// VerifyAuthHandler handles the auth verification route
func VerifyAuthHandler(c *gin.Context) {
	// Get the user from the context (set by the AuthRequired middleware)
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "認証されていません"})
		return
	}

	// Extract user info from the token
	token := user.(map[string]interface{})
	
	// Return user info
	c.JSON(http.StatusOK, gin.H{
		"authenticated": true,
		"user": gin.H{
			"uid":            token["uid"],
			"email":          token["email"],
			"email_verified": token["email_verified"],
			"auth_time":      token["auth_time"],
		},
	})
}
