package database

import (
	"fmt"
	"log"
	"os"

	"real-time-chat/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var User_DB *gorm.DB

func ConnectDB() {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	User_DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
}

func Init() {
	ConnectDB()
	err := User_DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database")
	}
	log.Println("Database connected")
}

func GetDB() *gorm.DB {
	return User_DB
}

func CloseDB() {
	db, err := User_DB.DB()
	if err != nil {
		log.Fatal("Failed to close the database")
	}
	db.Close()
}
