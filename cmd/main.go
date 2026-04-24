package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/mohdareeb0x-commits/book-library-api/internal/config"
	"github.com/mohdareeb0x-commits/book-library-api/internal/models"
	"github.com/mohdareeb0x-commits/book-library-api/internal/routes"
)

func main() {
	
	config.LoadConfig()
	server_config := viper.GetStringMapString("server")
	port := server_config["port"]
	if port == "" {
		port = "8080"
	}

	db := config.InitDB()

	if err := db.AutoMigrate(&models.Book{}); err != nil {
		log.Fatal("Migration failed:", err)
	}

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routes.SetupRoutes(router, db)

	log.Printf("Server running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
