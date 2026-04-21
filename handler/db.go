package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohdareeb0x-commits/book-library-api/models"
	"gorm.io/gorm"
)

func CreateBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		bookName := c.Query("book")
		authorName := c.Query("author")
		published := c.Query("published")
		units, err := strconv.Atoi(c.DefaultQuery("units", "1"))
		if err != nil {
			Fail(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "unable to process Book Units")
		}
		price, err := strconv.Atoi(c.DefaultQuery("price", "0"))
		log.Println(price)
		if err != nil {
			Fail(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "unable to process Book Prices")
		}

		if bookName == "" || authorName == "" || published == "" {
			Fail(c, http.StatusBadRequest, "REQUIRED_QUERY_EMPTY", "book name, author name or publish date is missing")
			return
		}

		book := models.Books{Name: bookName, Author: authorName, DatePublished: published, Units: units, Price: price}
		db.Create(&book)

		OK(c, book)
	}
}

func ListBooks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		books := []models.Books{}
		db.Find(&books)

		if len(books) == 0 {
			Fail(c, http.StatusOK, "NO_BOOKS_AVAILABLE", "database is empty")
			return
		}

		OK(c, books)
	}
}

func ListBooksByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil || id == 0 {
			Fail(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "unable to process ID")
			return
		}

		book := models.Books{}

		db.Where("id = ?", id).Find(&book)

		if book.ID == 0 {
			Fail(c, http.StatusNotFound, "NO_BOOK_AVAILABLE", fmt.Sprintf("no book available with id: %d", id))
			return
		}

		OK(c, book)

	}
}

func UpdateBookByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil || id == 0 {
			Fail(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "unable to process ID")
		}

		bookName := c.Query("book")
		authorName := c.Query("author")
		published := c.Query("published")
		units := c.Query("units")
		price := c.Query("price")

		if bookName == "" && authorName == "" && published == "" && units == "" && price == "" {
			Fail(c, http.StatusBadRequest, "REQUIRED_QUERY_EMPTY", "book name, author name or publish date is missing")
			return
		}

		book := models.Books{}

		db.Where("id = ?", id).Find(&book)

		if book.ID == 0 {
			Fail(c, http.StatusNotFound, "NO_BOOK_AVAILABLE", fmt.Sprintf("no book available with id: %d", id))
			return
		}

		if bookName != "" {
			db.Model(&book).Where("id = ?", id).Update("name", bookName)
		}
		if authorName != "" {
			db.Model(&book).Where("id = ?", id).Update("author", authorName)
		}
		if published != "" {
			db.Model(&book).Where("id = ?", id).Update("date_published", published)
		}
		if units != "" {
			units, err := strconv.Atoi(units)
			if err != nil {
				Fail(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "unable to process Book Units")
			}
			db.Model(&book).Where("id = ?", id).Update("units", units)
		}
		if price != "" {
			price, err := strconv.Atoi(price)
			if err != nil {
				Fail(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "unable to process Book Prices")
			}
			db.Model(&book).Where("id = ?", id).Update("price", price)
		}
		

		OK(c, book)

	}
}

func DeleteBookByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil || id == 0 {
			Fail(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "unable to process ID")
		}

		book := models.Books{}
		db.Where("id = ?", id).Find(&book)

		if book.ID == 0 {
			Fail(c, http.StatusNotFound, "NO_BOOK_AVAILABLE", fmt.Sprintf("no book available with id: %d", id))
			return
		}

		db.Where("id = ?", id).Delete(&book)

		OK(c, book)
	}
}
