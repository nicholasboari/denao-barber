package configs

import (
	"github.com/nicholasboari/denao-barber/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnector() *gorm.DB {
	dsn := "host=localhost user=postgres password=admin port=5432 dbname=denao sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helper.ErrorPanic(err)
	return db
}
