package database

import (
	"log"

	"garrison/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize() {
	dsn := "host=localhost user=docker password=docker dbname=garrison port=5432 sslmode=disable"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	DB.AutoMigrate(new(models.User))
}

func Migrate(models ...interface{}) {
	err := DB.AutoMigrate(models...)

	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}
}
