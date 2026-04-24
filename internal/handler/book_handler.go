package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohdareeb0x-commits/book-library-api/internal/dto"
	"github.com/mohdareeb0x-commits/book-library-api/internal/response"
	"github.com/mohdareeb0x-commits/book-library-api/internal/service"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var input dto.CreateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Fail(c, http.StatusBadRequest, "JSON_BINDING_ERROR", "unable to get request body")
		return
	}

	book, err := h.service.CreateBook(input)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "unable to create book")
		return
	}

	response.OK(c, book, nil)
}

func (h *BookHandler) ListBooks(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	books, meta, err := h.service.ListBooks(page, limit)
	if err != nil {
		response.Fail(c, http.StatusNotFound, "NOT_FOUND", "unable to get books")
		return
	}

	response.OK(c, books, meta)
}

func (h *BookHandler) ListBooksByID(c *gin.Context) {
	id := c.Param("id")

	book, err := h.service.ListBooksByID(id)
	if err != nil {
		response.Fail(c, http.StatusNotFound, "NOT_FOUND", "unable to get books by id: "+id)
		return
	}

	response.OK(c, book, nil)
}

func (h *BookHandler) UpdateBookByID(c *gin.Context) {
	id := c.Param("id")

	var input dto.UpdateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Fail(c, http.StatusBadRequest, "JSON_BINDING_ERROR", "unable to get request body")
		return
	}

	book, err := h.service.UpdateBookByID(id, input)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "unable to update book with id: "+id)
		return
	}

	response.OK(c, book, nil)
}

func (h *BookHandler) DeleteBookByID(c *gin.Context) {
	id := c.Param("id")

	book, err := h.service.DeleteBookByID(id)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "unable to delete book by id: "+id)
		return
	}

	response.OK(c, book, nil)
}

func (h *BookHandler) SearchBook(c *gin.Context) {
	name := c.Query("name")
	author := c.Query("author")
	limit := c.DefaultQuery("limit", "10")
	page := c.DefaultQuery("page", "1")

	books, meta, err := h.service.SearchBook(name, author, limit, page)
	if err != nil {
		response.Fail(c, http.StatusNotFound, "NOT_FOUND", "unable to find book")
		return
	}

	response.OK(c, books, meta)
}
