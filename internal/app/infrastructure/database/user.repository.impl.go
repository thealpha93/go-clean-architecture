package database

import (
	"test-server-app/internal/app/domain/entities"
	"test-server-app/internal/app/domain/repositories"
	model "test-server-app/internal/app/infrastructure/database/models"

	"gorm.io/gorm"
)

type UserRepoImpl struct {
	db *gorm.DB
}

func NewUserRepoImpl(db *gorm.DB) repositories.UserRepository {
	return &UserRepoImpl{db: db}
}

// CreateUser implements repositories.UserRepository.
func (uri *UserRepoImpl) CreateUser(usr *entities.User, password string) (*entities.User, error) {
	user := &model.UserModel{
		Email:    usr.Email,
		Password: password,
		Dob:      usr.Dob,
	}

	result := uri.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	entityUser := entities.User{
		ID:        user.ID,
		Email:     user.Email,
		Dob:       user.Dob,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAT: user.UpdatedAt,
	}
	return &entityUser, nil
}

// GetUserByEmail implements repositories.UserRepository.
func (uri *UserRepoImpl) GetUserByEmail(email string) (*model.UserModel, error) {
	var user model.UserModel

	if err := uri.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByID implements repositories.UserRepository.
func (uri *UserRepoImpl) GetUserByID(id uint64) (*entities.User, error) {
	var user model.UserModel

	if err := uri.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	userEntity := &entities.User{
		ID:    user.ID,
		Email: user.Email,
		Dob:   user.Dob,
	}

	return userEntity, nil
}
