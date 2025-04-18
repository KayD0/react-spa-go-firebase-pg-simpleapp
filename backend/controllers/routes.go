package controllers

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all routes for the application
func RegisterRoutes(router *gin.Engine) {
	// Register main routes
	RegisterMainRoutes(router)

	// Register auth routes
	RegisterAuthRoutes(router)

	// Register profile routes
	RegisterProfileRoutes(router)
}
