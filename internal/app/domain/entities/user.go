package entities

import "time"

type User struct {
	ID        uint64    `json:"id"`
	Email     string    `json:"email" validate:"required,email"`
	Dob       string    `json:"dob" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAT time.Time `json:"deleted_at"`
}
