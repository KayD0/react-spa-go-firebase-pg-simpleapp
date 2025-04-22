package middleware

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/baseapp/services"
)

// AuthRequired is a middleware that checks if the request has a valid Firebase ID token
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorizationヘッダーがありません"})
			c.Abort()
			return
		}

		// Extract the token (remove 'Bearer ' prefix if present)
		idToken := authHeader
		if strings.HasPrefix(authHeader, "Bearer ") {
			idToken = strings.TrimPrefix(authHeader, "Bearer ")
		}

		// Verify the token
		token, err := services.VerifyIDToken(idToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "無効な認証トークンです: " + err.Error()})
			c.Abort()
			return
		}

		// Store the token in the context
		c.Set("user", token)
		c.Next()
	}
}

// GetUserIDFromToken extracts the user ID from the token in the context
func GetUserIDFromToken(c *gin.Context) (string, error) {
	// Get the user from the context
	user, exists := c.Get("user")
	if !exists {
		return "", &services.AuthError{Message: "認証されていません。AuthRequiredミドルウェアを使用してください。"}
	}

	// Extract the user ID
	token := user.(map[string]interface{})

	uid, ok := token["user_id"].(string)
	if !ok {
		return "", &services.AuthError{Message: "ユーザーIDが見つかりません"}
	}

	return uid, nil
}
