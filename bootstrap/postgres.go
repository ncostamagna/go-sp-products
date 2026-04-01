package bootstrap

import (
	"fmt"
	"os"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/ncostamagna/go-sp-products/domain"
)

func InitPostgres() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}


if os.Getenv("DATABASE_DEBUG") == "true" {
	db = db.Debug()
}

if os.Getenv("DATABASE_MIGRATE") == "true" {
	if err := db.AutoMigrate(&domain.Product{}); err != nil {
		return nil, err
	}
}

	return db, nil
}
