package controller

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/baseapp/application/usecase"
)

// AuthController handles authentication-related HTTP requests
type AuthController struct {
    authUseCase *usecase.AuthUseCase
}

// NewAuthController creates a new AuthController
func NewAuthController(authUseCase *usecase.AuthUseCase) *AuthController {
    return &AuthController{
        authUseCase: authUseCase,
    }
}

// VerifyAuth handles the auth verification route
func (c *AuthController) VerifyAuth(ctx *gin.Context) {
    // Get the Authorization header
    authHeader := ctx.GetHeader("Authorization")
    if authHeader == "" {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorizationヘッダーがありません"})
        return
    }

    // Extract the token (remove 'Bearer ' prefix if present)
    idToken := authHeader
    if strings.HasPrefix(authHeader, "Bearer ") {
        idToken = strings.TrimPrefix(authHeader, "Bearer ")
    }

    // Verify the token
    response, err := c.authUseCase.VerifyToken(ctx.Request.Context(), idToken)
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    // Return the response
    ctx.JSON(http.StatusOK, response)
}

// GetUserIDFromToken extracts the user ID from the token in the context
func (c *AuthController) GetUserIDFromToken(ctx *gin.Context) (string, error) {
    // Get the Authorization header
    authHeader := ctx.GetHeader("Authorization")
    if authHeader == "" {
        return "", &AuthError{Message: "Authorizationヘッダーがありません"}
    }

    // Extract the token (remove 'Bearer ' prefix if present)
    idToken := authHeader
    if strings.HasPrefix(authHeader, "Bearer ") {
        idToken = strings.TrimPrefix(authHeader, "Bearer ")
    }

    // Verify the token
    response, err := c.authUseCase.VerifyToken(ctx.Request.Context(), idToken)
    if err != nil {
        return "", &AuthError{Message: err.Error()}
    }

    // Extract the user ID
    uid, err := c.authUseCase.GetUserIDFromClaims(response.User)
    if err != nil {
        return "", &AuthError{Message: err.Error()}
    }

    return uid, nil
}

// AuthError represents an authentication error
type AuthError struct {
    Message string
}

// Error returns the error message
func (e *AuthError) Error() string {
    return e.Message
}
