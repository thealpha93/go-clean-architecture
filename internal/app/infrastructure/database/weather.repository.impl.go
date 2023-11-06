package database

import (
	"encoding/json"
	"test-server-app/internal/app/domain/entities"
	"test-server-app/internal/app/domain/repositories"
	"test-server-app/internal/app/dto"
	model "test-server-app/internal/app/infrastructure/database/models"

	"gorm.io/gorm"
)

type WeatherRepoImpl struct {
	db *gorm.DB
}

func NewWeatherRepoImpl(db *gorm.DB) repositories.WeatherRepository {
	return &WeatherRepoImpl{db: db}
}

// CreateSearchWeatherHistory implements repositories.WeatherRepository.
func (wri *WeatherRepoImpl) CreateSearchWeatherHistory(weather dto.OpenWeatherResponseDTO, userID uint64) (*model.WeatherSearchHistory, error) {

	weatherJson, err := json.Marshal(weather.Weather)

	if err != nil {
		println("error occured while marshelling data")
	}

	weatherSearchHistory := &model.WeatherSearchHistory{
		Coord: model.Coord{
			Lon: weather.Coord.Lon,
			Lat: weather.Coord.Lat,
		},
		Weather:    weatherJson,
		Base:       weather.Base,
		Main:       model.Main(weather.Main),
		Visibility: weather.Visibility,
		Wind: model.Wind{
			Speed: weather.Wind.Speed,
			Deg:   weather.Wind.Deg,
		},
		Clouds: model.Clouds{
			All: weather.Clouds.All,
		},
		Dt: weather.Dt,
		Sys: model.Sys{
			Type:    weather.Sys.Type,
			ID:      weather.Sys.ID,
			Country: weather.Sys.Country,
			Sunrise: weather.Sys.Sunrise,
			Sunset:  weather.Sys.Sunset,
		},
		Timezone:      weather.Timezone,
		OpenWeatherID: int64(weather.ID),
		Name:          weather.Name,
		Cod:           weather.Cod,
		CreatedBy:     userID,
	}

	result := wri.db.Create(weatherSearchHistory)
	if result.Error != nil {
		return nil, result.Error
	}

	return weatherSearchHistory, nil
}

// GetSearchHistory implements repositories.WeatherRepository.
func (wri *WeatherRepoImpl) GetSearchHistory(q dto.GetSearchHistoryQueryParamsDTO, userID uint64) ([]model.WeatherSearchHistory, error) {
	var results []model.WeatherSearchHistory

	db := wri.db.Where("created_by = ?", userID).Limit(q.Limit)
	if q.Cursor != 0 {
		db = db.Where("id > ?", q.Cursor)
	}
	result := db.Find(&results)
	if result.Error != nil {
		return nil, result.Error
	}

	return results, wri.db.Error
}

// BulkDeleteSearchHistory implements repositories.WeatherRepository.
func (wri *WeatherRepoImpl) BulkDeleteSearchHistory(ids []uint64, userId uint64) error {
	// implements soft delete
	return wri.db.Where("id in ?", ids).Where("created_by = ?", userId).Unscoped().Delete(&model.WeatherSearchHistory{}).Error
}

// DeleteSearchHistory implements repositories.WeatherRepository.
func (*WeatherRepoImpl) DeleteSearchHistory(id uint64) error {
	panic("unimplemented")
}

// GetCurrentWeather implements repositories.WeatherRepository.
func (*WeatherRepoImpl) GetCurrentWeather(city string) (*entities.Weather, error) {
	panic("unimplemented")
}

// UpdateSearchHistory implements repositories.WeatherRepository.
func (*WeatherRepoImpl) UpdateSearchHistory(id uint64) (*entities.Weather, error) {
	panic("unimplemented")
}
