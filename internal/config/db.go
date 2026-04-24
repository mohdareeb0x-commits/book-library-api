package config

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("library.db"), &gorm.Config{})
	log.Println(os.Getenv("DB_URL"))
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
