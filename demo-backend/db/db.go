package db

import (
	"log"
	"os"

	"github.com/wolewd/database-docker-testing/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB_A *gorm.DB
var DB_B *gorm.DB

func InitDB() {
	var err error

	dsnA := os.Getenv("DATABASE_URL_A")
	dsnB := os.Getenv("DATABASE_URL_B")

	DB_A, err = gorm.Open(postgres.Open(dsnA), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB_A: ", err)
	}

	DB_B, err = gorm.Open(postgres.Open(dsnB), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB_B: ", err)
	}

	if err := DB_A.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("AutoMigrate DB_A failed: ", err)
	}

	if err := DB_B.AutoMigrate(&models.Product{}); err != nil {
		log.Fatal("AutoMigrate DB_B failed: ", err)
	}
}
