package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohdareeb0x-commits/book-library-api/handler"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB) {
	router := gin.Default()

	router.GET("/books", handler.ListBooks(db))
	router.GET("/books/:id", handler.ListBooksByID(db))
	router.POST("/books", handler.CreateBook(db))
	router.PUT("/books/:id", handler.UpdateBookByID(db))
	router.DELETE("/books/:id", handler.DeleteBookByID(db))

	router.Run()
}
