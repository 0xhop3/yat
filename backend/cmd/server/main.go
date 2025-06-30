package main

import (
	"github.com/0xhop3/yat/backend/internal/config"
	"github.com/0xhop3/yat/backend/internal/database"
	"github.com/0xhop3/yat/backend/internal/handlers"
	"github.com/0xhop3/yat/backend/internal/middleware"
	"github.com/0xhop3/yat/backend/internal/repositories"
	"github.com/0xhop3/yat/backend/internal/services"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Load configuration
	cfg := config.Load()
	if cfg == nil {
		log.Fatal("Failed to load configuration")
	}

	// Initialize database connections
	db, err := database.NewPostgresConnection(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepo)

	// Initialize middleware
	authenticationMiddleware := middleware.NewAuthenticationMiddleware(cfg, userService)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)

	// Setup router
	router := setupRouter(cfg, authenticationMiddleware, userHandler)

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(router.Run(":" + cfg.Port))
}

func setupRouter(_ *config.Config, authenticationMiddleware *middleware.AuthenticationMiddleware, userHandler *handlers.UserHandler) *gin.Engine {
	router := gin.Default()

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(corsMiddleware())

	// Health check endpoint (no auth required)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API routes
	api := router.Group("/api/v1")
	{

		// Protected routes (authentication required)
		protected := api.Group("")
		protected.Use(authenticationMiddleware.ValidateJWT())
		{
			// User routes
			users := protected.Group("/users")
			{
				users.POST("", userHandler.CreateUser) // Admin only
			}
		}
	}

	return router
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
