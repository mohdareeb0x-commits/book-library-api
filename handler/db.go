package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohdareeb0x-commits/book-library-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateDB() *gorm.DB {
	db, err := gorm.Open((sqlite.Open("library.db")))
	if err != nil {
		panic("Unable to create DB")
	}
	return db
}

func CreateBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		form := models.Books{
			Units: 0,
			Price: 0,
		}

		if err := c.ShouldBind(&form); err != nil {
			Fail(c, http.StatusBadRequest, "FORM_BINDING_ERROR", "unable to get form")
			return
		}

		db.Create(&form)

		OK(c, form, nil)
	}
}

func ListBooks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil {
			Fail(c, http.StatusInternalServerError, "PAGINATION_ERROR", "unable to process page")
			return
		}

		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if err != nil {
			Fail(c, http.StatusInternalServerError, "PAGINATION_ERROR", "unable to process limit")
			return
		}
		
		offset := (page - 1) * limit

		books := []models.Books{}
		if err := db.Limit(limit).Offset(offset).Find(&books).Error; err != nil {
			Fail(c, http.StatusOK, "NO_BOOKS_AVAILABLE", "database is empty")
			return
		}

		meta := models.Meta{
			Page: page,
			Limit: limit,
			Offset: offset,
		}

		OK(c, books, &meta)
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

		if err := db.Where("id = ?", id).First(&book, id).Error; err != nil {
			Fail(c, http.StatusNotFound, "NO_BOOK_AVAILABLE", fmt.Sprintf("no book available with id: %d", id))
			return
		}

		OK(c, book, nil)

	}
}

func UpdateBookByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil || id == 0 {
			Fail(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "unable to process ID")
			return
		}

		book := models.Books{}

		if err := db.Where("id = ?", id).First(&book).Error; err != nil {
			Fail(c, http.StatusNotFound, "NOT_FOUND", fmt.Sprintf("no book available with id: %d", id))
			return
		}

		if err := c.ShouldBind(&book); err != nil {
			Fail(c, http.StatusBadRequest, "FORM_BINDING_ERROR", "unable to get form")
			return
		}

		if err := db.Save(&book).Error; err != nil {
			Fail(c, http.StatusInternalServerError, "DB_ERROR", "unable to update book")
			return
		}

		OK(c, book, nil)

	}
}

func DeleteBookByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil || id == 0 {
			Fail(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "unable to process ID")
			return
		}

		book := models.Books{}
		if err := db.Where("id = ?", id).First(&book).Error; err != nil {
			Fail(c, http.StatusNotFound, "NO_BOOK_AVAILABLE", fmt.Sprintf("no book available with id: %d", id))
			return
		}

		if err := db.Where("id = ?", id).Delete(&book).Error; err != nil {
			Fail(c, http.StatusInternalServerError, "DB_ERROR", "unable to delete book")
			return
		}

		OK(c, book, nil)
	}
}
