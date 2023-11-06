package auth

import (
	"errors"
	"fmt"
	"strings"
	"test-server-app/internal/app/config"
	"test-server-app/internal/app/domain/entities"
	"test-server-app/internal/app/domain/repositories"
	"test-server-app/internal/app/dto"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo repositories.UserRepository
}

const (
	accessTokenDuration  = 7 * 24 * time.Hour
	refreshTokenDuration = 7 * 24 * time.Hour
)

var (
	JWT_SECRET           = config.AppConfig.JWT_SECRET
	REFRESH_TOKEN_SECRET = config.AppConfig.REFRESH_TOKEN_SECRET
)

var (
	ErrInvalidCredentials = errors.New("invalid username or password")
)

type Claims struct {
	UserId uint64 `json:"user_id"`
	jwt.StandardClaims
}

type RefreshTokenClaims struct {
	UserId uint64 `json:"user_id"`
	jwt.StandardClaims
}

func NewAuthService(ur repositories.UserRepository) *AuthService {
	return &AuthService{userRepo: ur}
}

func (as *AuthService) Authenticate(email, password string) (*dto.UserLoginResponseDTO, error) {
	user, err := as.userRepo.GetUserByEmail(email)

	if err != nil {
		return nil, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("Invalid Credentials")
	}

	// Generate Access token
	accessToken, err := generateAccessToken(user.ID, accessTokenDuration)
	if err != nil {
		return nil, err
	}

	refreshToken, err := generateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	userInfo := &entities.User{
		Email:     user.Email,
		Dob:       user.Dob,
		CreatedAt: user.CreatedAt,
	}
	return &dto.UserLoginResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserInfo:     *userInfo,
	}, nil
}

func (as *AuthService) CreateUser(usr dto.UserRegistrationDTO) (*entities.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	userEntity := &entities.User{
		Email: usr.Email,
		Dob:   usr.Dob,
	}

	user, err := as.userRepo.CreateUser(userEntity, string(hashedPassword))

	if err != nil {
		println("user creation failed %s", err.Error())
		return nil, err
	}

	return user, err
}

func (as *AuthService) GetUserByID(userId uint64) (*entities.User, error) {
	user, err := as.userRepo.GetUserByID(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func generateAccessToken(userId uint64, duration time.Duration) (string, error) {
	claims := Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWT_SECRET))
}

func generateRefreshToken(userId uint64) (string, error) {
	claims := Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(refreshTokenDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(REFRESH_TOKEN_SECRET))
}

var ErrInvalidToken = errors.New("Invalid token")

func (as *AuthService) ParseToken(token string) (uint64, error) {
	token = strings.TrimSpace(token)

	if token == "" {
		return 0, ErrInvalidToken
	}

	_token, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(JWT_SECRET), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := _token.Claims.(*Claims)
	if !ok || !_token.Valid {
		return 0, ErrInvalidToken
	}

	return claims.UserId, nil
}
