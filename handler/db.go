package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohdareeb0x-commits/book-library-api/models"
	"gorm.io/gorm"
)

func CreateBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		bookName := c.Query("book_name")
		authorName := c.Query("author_name")
		published := c.Query("published")

		if bookName == "" || authorName == "" || published == "" {
			Fail(c, http.StatusBadRequest, "REQUIRED_QUERY_EMPTY", "book name, author name or publish date is missing")
			return
		}

		book := models.Books{Name: bookName, Author: authorName, DatePublished: published}
		db.Create(&book)

		OK(c, book)
	}
}

func ListBooks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		book := []models.Books{}

		db.Find(&book)

		if book[0].ID == 0 {
			Fail(c, http.StatusNoContent, "NO_BOOKS_AVAILABLE", "database is empty")
		}

		OK(c, book)
	}
}

func ListBooksByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.DefaultQuery("id", "1"))
		if err != nil {
			Fail(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "unable to process ID")
		}

		book := models.Books{}

		db.Where("id = ?", id).Find(&book)

		OK(c, book)

	}
}
