// user_repository_mock.go
package auth

import (
	"test-server-app/internal/app/domain/entities"
	"test-server-app/internal/app/domain/repositories"
	model "test-server-app/internal/app/infrastructure/database/models"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) CreateUser(usr *entities.User, password string) (*entities.User, error) {
	args := m.Called(usr, password)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *UserRepositoryMock) GetUserByEmail(email string) (*model.UserModel, error) {
	args := m.Called(email)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.UserModel), args.Error(1)
}

func (m *UserRepositoryMock) GetUserByID(id uint64) (*entities.User, error) {
	args := m.Called(id)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), args.Error(1)
}

var _ repositories.UserRepository = &UserRepositoryMock{}
