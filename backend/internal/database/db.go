package database

import (
	"fmt"
	"log"

	"github.com/KeishiIrisa/backend-go-template/internal/config"

	"github.com/KeishiIrisa/backend-go-template/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func AutoMigrate() {
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	// Seed the database with seed data
	if err := UserSeed(DB); err != nil {
		log.Fatalf("Failed to seed data to the database: %v", err)
	}
}

func InitDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.AppConfig.PostgresHost,
		config.AppConfig.PostgresUser,
		config.AppConfig.PostgresPassword,
		config.AppConfig.PostgresDb,
		config.AppConfig.PostgresPort)

	var dbErr error
	DB, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		log.Fatalf("Failed to connect to the database: %v", dbErr)
	}

	log.Println("Database connected successfully!")
}
