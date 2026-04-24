package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohdareeb0x-commits/book-library-api/internal/handler"
	"github.com/mohdareeb0x-commits/book-library-api/internal/repository"
	"github.com/mohdareeb0x-commits/book-library-api/internal/service"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	repo := repository.NewBookRepository(db)
	service := service.NewBookService(repo)
	handler := handler.NewBookHandler(service)

	books := r.Group("/books")
	{
		books.GET("/", handler.ListBooks)
		books.GET("/:id", handler.ListBooksByID)
		books.GET("/search", handler.SearchBook)
		books.POST("/", handler.CreateBook)
		books.PATCH("/:id", handler.UpdateBookByID)
		books.DELETE("/:id", handler.DeleteBookByID)
	}
}
