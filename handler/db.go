package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohdareeb0x-commits/book-library-api/models"
	"gorm.io/gorm"
)

func CreateBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		bookName := c.Query("book_name")
		authorName := c.Query("author_name")
		published := c.Query("published")

		book := models.Books{Name: bookName, Author: authorName, DatePublished: published}

		db.Create(&book)

		if bookName == "" || authorName == "" || published == "" {
			Fail(c, http.StatusBadRequest, "REQUIRED_QUERY_EMPTY", "book name, author name or publish date is missing")
			return
		}
		OK(c, book)
	}
}

func ListBooks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		book := []models.Books{}

		db.Find(&book)

		OK(c, book)
	}
}
