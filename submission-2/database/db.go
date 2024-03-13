package database

import (
	"fmt"
	helpers "submission-2/helper"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
	var (
		host     = helpers.GetEnv("DB_HOST")
		port     = helpers.GetEnv("DB_PORT")
		user     = helpers.GetEnv("DB_USERNAME")
		password = helpers.GetEnv("DB_PASSWORD")
		dbname   = helpers.GetEnv("DB_NAME")
		db       *gorm.DB
		err      error
	)
	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	return db, err
}
