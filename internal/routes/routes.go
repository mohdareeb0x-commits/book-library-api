package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohdareeb0x-commits/book-library-api/internal/handler"
	"github.com/mohdareeb0x-commits/book-library-api/internal/middleware"
	"github.com/mohdareeb0x-commits/book-library-api/internal/repository"
	"github.com/mohdareeb0x-commits/book-library-api/internal/service"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {

	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	books := r.Group("/books")
	books.Use(middleware.AuthMiddleware())
	{
		books.GET("/", bookHandler.ListBooks)
		books.GET("/:id", bookHandler.ListBooksByID)
		books.GET("/search", bookHandler.SearchBook)
	}

	adminRoutes := r.Group("/admin")
	adminRoutes.Use(middleware.AuthMiddleware(), middleware.AdminOnly)
	{
		books := adminRoutes.Group("/books")
		books.POST("/", bookHandler.CreateBook)
		books.PATCH("/:id", bookHandler.UpdateBookByID)
		books.DELETE("/:id", bookHandler.DeleteBookByID)
	}

	users := r.Group("/user")
	{
		users.POST("/register", authHandler.CreateUser)
		users.POST("/login", authHandler.Login)
		users.POST("/logout", authHandler.Logout)
	}
}
