package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func LoadConfig() {
	viper.AddConfigPath("./internal/config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func InitDB() *gorm.DB {
	db_params := viper.GetStringMapString("db_params")
	fmt.Println(db_params["db_directory"])
	err := os.MkdirAll(db_params["db_directory"], 0755)
	if err != nil {
		panic("failed to create database")
	}
	db, err := gorm.Open(sqlite.Open(db_params["db_directory"]+"/"+db_params["db_name"]), &gorm.Config{})
	// log.Println(os.Getenv("DB_URL"))
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
