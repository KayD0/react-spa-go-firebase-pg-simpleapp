package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterMainRoutes registers the main routes
func RegisterMainRoutes(router *gin.Engine) {
	router.GET("/", IndexHandler)
}

// IndexHandler handles the index route
func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ユーザープロフィールAPIが実行中です",
		"endpoints": gin.H{
			"auth_verify": "/api/auth/verify (Authorizationヘッダーを持つPOST)",
			"profile_get": "/api/profile/ (Authorizationヘッダーを持つGET)",
			"profile_update": "/api/profile/ (JSONボディとAuthorizationヘッダーを持つPUT)",
		},
	})
}
