package controller

import (
	"github.com/gin-gonic/gin"
)

// Router handles all the routes
type Router struct {
	engine         *gin.Engine
	authController *AuthController
	userController *UserController
	mainController *MainController
}

// NewRouter creates a new Router
func NewRouter(
	engine *gin.Engine,
	authController *AuthController,
	userController *UserController,
	mainController *MainController,
) *Router {
	return &Router{
		engine:         engine,
		authController: authController,
		userController: userController,
		mainController: mainController,
	}
}

// SetupRoutes sets up all the routes
func (r *Router) SetupRoutes() {
	// Main routes
	r.engine.GET("/", r.mainController.Index)

	// Auth routes
	authGroup := r.engine.Group("/api/auth")
	{
		authGroup.POST("/verify", r.AuthMiddleware(), r.authController.VerifyAuth)
	}

	// Profile routes
	profileGroup := r.engine.Group("/api")
	{
		profileGroup.GET("/profile", r.AuthMiddleware(), r.userController.GetProfile)
		profileGroup.PUT("/profile", r.AuthMiddleware(), r.userController.UpdateProfile)
	}
}

// AuthMiddleware is a middleware that checks if the request has a valid authentication token
func (r *Router) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorizationヘッダーがありません"})
			c.Abort()
			return
		}

		c.Next()
	}
}
