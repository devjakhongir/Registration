package utils

import (
	"errors"
	
	"app/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DB() (*gorm.DB, error) {

	PGConnection := config.DBConfig()

	return gorm.Open(postgres.Open(PGConnection), &gorm.Config{})
}

func IsNotFound(row *gorm.DB) bool {

	return errors.Is(row.Error, gorm.ErrRecordNotFound)
}