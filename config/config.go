package config

import (
	"fmt"
	"log"
	"os"

	"github.com/aysf/tokogo/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err1 := godotenv.Load("./.env")
	if err1 != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	var err error

	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	InitialMigration()
}

func InitialMigration() {
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Cart{})
}
