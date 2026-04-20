package store

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateDB() (*gorm.DB) {
	db, err := gorm.Open((sqlite.Open("library.db")))
	if err != nil {
		panic("Unable to create DB")
	}
	return db
}