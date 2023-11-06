package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"test-server-app/internal/app/config"
	"test-server-app/internal/app/domain/repositories"
	"test-server-app/internal/app/dto"
	model "test-server-app/internal/app/infrastructure/database/models"
)

type WeatherService struct {
	weatherRepo repositories.WeatherRepository
}

func NewWeatherService(wr repositories.WeatherRepository) *WeatherService {
	return &WeatherService{weatherRepo: wr}
}

func (ws *WeatherService) GetCurrentWeather(city string, userId uint64) (*dto.GetCurrentWeatherResponseDTO, error) {
	// Call the external api to get the lat and long of the city
	// This 3rd party api should ideally be another service and should have been injected into this service
	// Letting this be for now because of time constraint and because this is a test task
	// TODO: Inject OpenWeatherService into weather service
	weatherUrl := fmt.Sprintf("%s?q=%s&appId=%s&units=metric", config.AppConfig.OPEN_WEATHER_BASE_URL, city, config.AppConfig.OPEN_WEATHER_API_KEY)
	resp, err := http.Get(weatherUrl)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Weather api failed with status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var currentWeather dto.OpenWeatherResponseDTO
	if err := json.Unmarshal(body, &currentWeather); err != nil {
		return nil, err
	}

	// Write to the search history
	weather, err := ws.weatherRepo.CreateSearchWeatherHistory(currentWeather, userId)

	// Transform to the required DTO
	weatherSearchHistory := transformWeatherSearchModelToDTO(weather)

	// For now choosing to return the whole data to front end.
	// Mainly because I haven't decided what to show in the front end.
	// Not ideal
	// Send back the response
	return weatherSearchHistory, nil
}

func (ws *WeatherService) GetSearchHistory(q dto.GetSearchHistoryQueryParamsDTO, userId uint64) (*dto.GetSearchHistoryResponseDTO, error) {
	// panic("not implemented")
	searchHistory, err := ws.weatherRepo.GetSearchHistory(q, userId)
	if err != nil {
		return nil, err
	}

	searchHistoryList := make([]*dto.GetCurrentWeatherResponseDTO, len(searchHistory))

	for i, sh := range searchHistory {
		searchHistoryList[i] = transformWeatherSearchModelToDTO(&sh)
	}
	var nextCursor int
	if len(searchHistoryList) == q.Limit {
		nextCursor = searchHistoryList[len(searchHistoryList)-1].ID
	}
	response := &dto.GetSearchHistoryResponseDTO{
		Data:       searchHistoryList,
		NextCursor: nextCursor,
	}
	return response, nil
}

func (ws *WeatherService) BulkDeleteSearchHistory(ids []uint64, userId uint64) error {
	return ws.weatherRepo.BulkDeleteSearchHistory(ids, userId)
}

func transformWeatherSearchModelToDTO(weather *model.WeatherSearchHistory) *dto.GetCurrentWeatherResponseDTO {
	var dtoWeather []dto.Weather
	err := json.Unmarshal(weather.Weather, &dtoWeather)

	if err != nil {
		println("error occured while unmarshelling data")
	}
	weatherSearchHistory := &dto.GetCurrentWeatherResponseDTO{
		Coord: dto.Coord{
			Lon: weather.Coord.Lon,
			Lat: weather.Coord.Lat,
		},
		Weather:    dtoWeather,
		Base:       weather.Base,
		Main:       dto.Main(weather.Main),
		Visibility: weather.Visibility,
		Wind: dto.Wind{
			Speed: weather.Wind.Speed,
			Deg:   weather.Wind.Deg,
		},
		Clouds: dto.Clouds{
			All: weather.Clouds.All,
		},
		Dt: weather.Dt,
		Sys: dto.Sys{
			Type:    weather.Sys.Type,
			ID:      weather.Sys.ID,
			Country: weather.Sys.Country,
			Sunrise: weather.Sys.Sunrise,
			Sunset:  weather.Sys.Sunset,
		},
		Timezone:      weather.Timezone,
		OpenWeatherID: int64(weather.OpenWeatherID),
		Name:          weather.Name,
		Cod:           weather.Cod,
		CreatedBy:     weather.CreatedBy,
		ID:            int(weather.ID),
	}
	return weatherSearchHistory
}
