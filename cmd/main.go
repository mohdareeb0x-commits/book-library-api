package main

import (
	"github.com/mohdareeb0x-commits/book-library-api/handler"
	"github.com/mohdareeb0x-commits/book-library-api/models"
	"github.com/mohdareeb0x-commits/book-library-api/routes"
)

func main() {
	db := handler.CreateDB()
	db.AutoMigrate(&models.Books{})

	routes.Routes(db)
}
