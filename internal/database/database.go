package database

import (
	"log"
	"os"

	"github.com/krushalgopale/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnDB() {
	dsn := os.Getenv("DB_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error connecting database: ", err)
	}

	log.Println("Database connected successfully")
  db.AutoMigrate(&models.User{}, &models.Patient{})
	DB = db
}
