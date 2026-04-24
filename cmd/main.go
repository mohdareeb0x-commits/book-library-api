package main

import (
	"os"
	"log"
	"github.com/gin-gonic/gin"

	"github.com/mohdareeb0x-commits/book-library-api/internal/routes"
	"github.com/mohdareeb0x-commits/book-library-api/internal/config"
	"github.com/mohdareeb0x-commits/book-library-api/internal/models"
)

func main() {
	// Load environment variables (optional but recommended)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize database
	db := config.InitDB()

	// Run migrations
	if err := db.AutoMigrate(&models.Book{}); err != nil {
		log.Fatal("Migration failed:", err)
	}

	// Initialize Gin router
	router := gin.New()

	// Attach middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Setup routes
	routes.SetupRoutes(router, db)

	// Start server
	log.Printf("Server running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
