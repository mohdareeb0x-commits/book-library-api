package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohdareeb0x-commits/book-library-api/handler"
	// "github.com/mohdareeb0x-commits/book-library-api/store"
)

func Routes() {
	router := gin.Default()

	router.GET("/books", )
	router.POST("/books", handler.CreateBook)

}