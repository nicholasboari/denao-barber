package configs

import (
	"fmt"

	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	var err error
	db, err = InitPostgres()
	if err != nil {
		return fmt.Errorf("error initializing postgres: %v", err)
	}
	return nil
}

func GetPostgres() *gorm.DB {
	return db
}
