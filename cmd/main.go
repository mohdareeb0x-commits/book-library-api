package main

import (
	"github.com/mohdareeb0x-commits/book-library-api/routes"
	"github.com/mohdareeb0x-commits/book-library-api/models"
	"github.com/mohdareeb0x-commits/book-library-api/store"
)

func main() {
	db := store.CreateDB()

	db.AutoMigrate(&models.Books{})

	routes.Routes(db)
}
