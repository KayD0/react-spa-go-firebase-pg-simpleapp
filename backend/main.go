package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/baseapp/backend/application/usecase"
	"github.com/baseapp/backend/infrastructure/auth"
	"github.com/baseapp/backend/infrastructure/persistence"
	"github.com/baseapp/backend/interface/controller"
	"github.com/baseapp/backend/interface/repository"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Initialize router
	router := gin.Default()

	// Configure CORS
	corsOrigin := os.Getenv("CORS_ORIGIN")
	if corsOrigin == "" {
		corsOrigin = "http://localhost:3000"
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{corsOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Initialize Firebase Auth
	firebaseAuth, err := auth.NewFirebaseAuth()
	if err != nil {
		log.Printf("Warning: Failed to initialize Firebase Admin SDK: %v", err)
	}

	// Initialize database
	db, err := persistence.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Auto migrate the schema
	if err := db.AutoMigrate(&persistence.UserProfileModel{}); err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}
	log.Println("Database connected and migrated successfully")

	// Initialize repositories
	authRepo := repository.NewAuthRepository(firebaseAuth)
	userRepo := repository.NewUserRepository(db.DB)

	// Initialize use cases
	authUseCase := usecase.NewAuthUseCase(authRepo)
	userUseCase := usecase.NewUserUseCase(userRepo)

	// Initialize controllers
	authController := controller.NewAuthController(authUseCase)
	mainController := controller.NewMainController()
	userController := controller.NewUserController(userUseCase, authController)

	// Initialize router
	routerHandler := controller.NewRouter(router, authController, userController, mainController)
	routerHandler.SetupRoutes()

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// Start server
	fmt.Printf("Server running on port %s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
