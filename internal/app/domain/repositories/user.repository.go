package repositories

import (
	"test-server-app/internal/app/domain/entities"
	model "test-server-app/internal/app/infrastructure/database/models"
)

type UserRepository interface {
	GetUserByID(id uint64) (*entities.User, error)
	GetUserByEmail(email string) (*model.UserModel, error)
	CreateUser(user *entities.User, password string) (*entities.User, error)
}
