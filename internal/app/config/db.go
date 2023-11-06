package config

import (
	model "test-server-app/internal/app/infrastructure/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(databaseURL string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	Migrate(db)
	return db, nil
}

func Migrate(db *gorm.DB) {

	db.Migrator().AutoMigrate(&model.UserModel{})
	db.Migrator().AutoMigrate(&model.WeatherSearchHistory{})
}
