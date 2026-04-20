package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohdareeb0x-commits/book-library-api/handler"
	"github.com/mohdareeb0x-commits/book-library-api/models"
	"github.com/mohdareeb0x-commits/book-library-api/store"
)

func Routes() {
	router := gin.Default()
	db := store.CreateDB()

	db.AutoMigrate(&models.Books{})

	router.GET("/books", handler.ListBooks(db))
	router.GET("/books/:id", handler.ListBooks(db))
	router.POST("/books", handler.CreateBook(db))

	router.Run()
}
