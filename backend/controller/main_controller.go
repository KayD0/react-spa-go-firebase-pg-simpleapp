package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MainController handles general HTTP requests
type MainController struct {
}

// NewMainController creates a new MainController
func NewMainController() *MainController {
	return &MainController{}
}

// Index handles the index route
func (c *MainController) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ユーザープロフィールAPIが実行中です",
		"endpoints": gin.H{
			"auth_verify": "/api/auth/verify (Authorizationヘッダーを持つPOST)",
			"profile_get": "/api/profile/ (Authorizationヘッダーを持つGET)",
			"profile_update": "/api/profile/ (JSONボディとAuthorizationヘッダーを持つPUT)",
		},
	})
}
