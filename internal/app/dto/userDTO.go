package dto

import "test-server-app/internal/app/domain/entities"

type UserRegistrationDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Dob      string `json:"dob" validate:"required"`
}

type UserLoginRequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserLoginResponseDTO struct {
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	UserInfo     entities.User `json:"user_info"`
}
