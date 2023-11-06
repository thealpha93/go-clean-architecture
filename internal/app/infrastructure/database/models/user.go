package model

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ID        uint64 `gorm:"primaryKey;autoIncrement:true"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	Dob       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (UserModel) TableName() string {
	return "users"
}
