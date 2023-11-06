package auth

import (
	"errors"
	"test-server-app/internal/app/domain/entities"
	"test-server-app/internal/app/dto"
	model "test-server-app/internal/app/infrastructure/database/models"
	"test-server-app/internal/app/usecases/auth"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestAuthService_Authenticate(t *testing.T) {

	t.Run("should authenticate user", func(t *testing.T) {
		userRepoMock := new(UserRepositoryMock)
		authService := auth.NewAuthService(userRepoMock)

		expectedUserModel := &model.UserModel{
			Email:    "test@test.com",
			Password: "$2a$10$maC/9Cbybm57d2d6xGYjE.FvNEc/qvww8sqYiSoUbUD/qyDZeFrNq", // hashed password for "Test123!@#"
		}

		expectedUser := &entities.User{
			Email: "test@test.com",
			Dob:   "2000-01-01",
		}

		userRepoMock.On("GetUserByEmail", "test@test.com").Return(expectedUserModel, nil)
		response, err := authService.Authenticate("test@test.com", "Test123!@#")

		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, response.AccessToken)
		assert.NotNil(t, response.RefreshToken)
		assert.Equal(t, expectedUser.Email, response.UserInfo.Email)

		userRepoMock.AssertExpectations(t)
	})

	t.Run("should return error for invalid user", func(t *testing.T) {
		userRepoMock := new(UserRepositoryMock)
		authService := auth.NewAuthService(userRepoMock)

		userRepoMock.On("GetUserByEmail", "test@test.com").Return(nil, errors.New("user not found")).Once()
		response, err := authService.Authenticate("test@test.com", "Test123!@#")

		assert.NotNil(t, err)
		assert.Nil(t, response)
		assert.Equal(t, "user not found", err.Error())

		userRepoMock.AssertExpectations(t)
	})

	t.Run("Should return error for invalid password", func(t *testing.T) {
		userRepoMock := new(UserRepositoryMock)
		authService := auth.NewAuthService(userRepoMock)

		expectedUserModel := &model.UserModel{
			Email:    "test@test.com",
			Password: "$2a$10$maC/9Cbybm57d2d6xGYjE.FvNEc/qvww8sqYiSoUbUD/qyDZeFrNq", // hashed password for "Test123!@#"
		}

		userRepoMock.On("GetUserByEmail", expectedUserModel.Email).Return(expectedUserModel, nil).Once()
		// Supply wrong password
		response, err := authService.Authenticate(expectedUserModel.Email, "WrongPassword123")

		assert.Nil(t, response)
		assert.NotNil(t, err)
		assert.Equal(t, "Invalid Credentials", err.Error())

	})
}

func TestAuthService_CreateUser(t *testing.T) {
	userRepoMock := new(UserRepositoryMock)
	authService := auth.NewAuthService(userRepoMock)

	t.Run("should create user successfully", func(t *testing.T) {
		userDTO := dto.UserRegistrationDTO{
			Email:    "test@test.com",
			Password: "password123",
			Dob:      "2000-01-01",
		}

		userEntity := &entities.User{
			Email: userDTO.Email,
			Dob:   userDTO.Dob,
		}

		userRepoMock.On("CreateUser", mock.AnythingOfType("*entities.User"), mock.AnythingOfType("string")).Return(&entities.User{ID: 1, Email: "test@test.com", Dob: "2000-01-01"}, nil).Once()

		user, err := authService.CreateUser(userDTO)

		assert.Nil(t, err)
		if user != nil {
			assert.Equal(t, uint64(1), user.ID)
			assert.Equal(t, userEntity.Email, user.Email)

		}
		userRepoMock.AssertExpectations(t)
	})

	t.Run("Should not create user if the the email already exists", func(t *testing.T) {
		userDTO := dto.UserRegistrationDTO{
			Email:    "test@test.com",
			Password: "password123",
			Dob:      "2000-01-01",
		}

		userRepoMock.On("CreateUser", mock.AnythingOfType("*entities.User"), mock.AnythingOfType("string")).Return(nil, gorm.ErrDuplicatedKey).Once()

		user, err := authService.CreateUser(userDTO)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.Equal(t, gorm.ErrDuplicatedKey.Error(), err.Error())

		userRepoMock.AssertExpectations(t)
	})
}

func TestAuthService_GetUserById(t *testing.T) {
	userRepoMock := new(UserRepositoryMock)
	authService := auth.NewAuthService(userRepoMock)

	userEntity := &entities.User{
		ID:    1,
		Email: "test@test.com",
		Dob:   "2000-01-01",
	}

	t.Run("Should return the user successfully", func(t *testing.T) {
		userRepoMock.On("GetUserByID", uint64(1)).Return(userEntity, nil).Once()

		usr, err := authService.GetUserByID(1)

		assert.Nil(t, err)
		assert.NotNil(t, usr)
		assert.Equal(t, userEntity, usr)

		userRepoMock.AssertExpectations(t)
	})

	t.Run("Should return error if user is not found", func(t *testing.T) {
		userRepoMock.On("GetUserByID", uint64(1)).Return(nil, gorm.ErrRecordNotFound).Once()

		usr, err := authService.GetUserByID(1)

		assert.Nil(t, usr)
		assert.NotNil(t, err)
		assert.Equal(t, gorm.ErrRecordNotFound.Error(), err.Error())

	})
}
