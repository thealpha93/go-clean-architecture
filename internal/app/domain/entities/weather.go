package entities

import "time"

type Weather struct {
	ID          uint64    `json:"id"`
	City        string    `json:"city" validate:"required"`
	Temperature float64   `json:"temperature"`
	Condition   string    `json:"condition"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
