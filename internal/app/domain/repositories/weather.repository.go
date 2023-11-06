package repositories

import (
	"test-server-app/internal/app/domain/entities"
	"test-server-app/internal/app/dto"
	model "test-server-app/internal/app/infrastructure/database/models"
)

type WeatherRepository interface {
	CreateSearchWeatherHistory(weather dto.OpenWeatherResponseDTO, userID uint64) (*model.WeatherSearchHistory, error)
	GetCurrentWeather(city string) (*entities.Weather, error)
	GetSearchHistory(q dto.GetSearchHistoryQueryParamsDTO, userID uint64) ([]model.WeatherSearchHistory, error)
	UpdateSearchHistory(id uint64) (*entities.Weather, error)
	DeleteSearchHistory(id uint64) error
	BulkDeleteSearchHistory(ids []uint64, userID uint64) error
}
