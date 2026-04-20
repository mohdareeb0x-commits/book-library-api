package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mohdareeb0x-commits/book-library-api/models"
	"github.com/mohdareeb0x-commits/book-library-api/store"
	// "gorm.io/gorm"
)

func CreateBook(c *gin.Context) {
	db := store.CreateDB()

	bookName := c.Query("book_name")
	authorName := c.Query("author_name")
	published := c.Query("published")

	book := models.Books{Name: bookName, Author: authorName, DatePublished: published}

	db.Create(&book)

	

}

func ListBooks(c *gin.Context)
