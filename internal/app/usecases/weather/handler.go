package weather

import (
	"encoding/json"
	"net/http"
	"strconv"
	"test-server-app/internal/app/dto"

	"github.com/go-playground/validator/v10"
)

type WeatherHandler struct {
	WeatherService *WeatherService
}

func NewWeatherHandler(weatherService *WeatherService) *WeatherHandler {
	return &WeatherHandler{
		WeatherService: weatherService,
	}
}

var validate = validator.New()

func (h *WeatherHandler) GetCurrentWeatherHandler(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value("userID").(uint64)

	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "Please provide a city name", http.StatusBadRequest)
		return
	}

	result, err := h.WeatherService.GetCurrentWeather(city, userID)
	if err != nil {
		http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *WeatherHandler) GetSearchHistoryHandler(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value("userID").(uint64)

	limitInt := 10

	limit := r.URL.Query().Get("limit")
	cursor := r.URL.Query().Get("cursor")

	if limit != "" {
		var err error
		limitInt, err = strconv.Atoi(limit)
		if err != nil {
			http.Error(w, "Invalid limit", http.StatusBadRequest)
			return
		}
	}

	cursorInt, err := strconv.Atoi(cursor)

	if err != nil {
		// handle the error if the conversion fails
	}

	qParams := dto.GetSearchHistoryQueryParamsDTO{
		Limit:  limitInt,
		Cursor: cursorInt,
	}

	err = validate.Struct(qParams)
	if err != nil {
		// handle the error if the conversion fails
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.WeatherService.GetSearchHistory(qParams, userID)
	if err != nil {
		http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *WeatherHandler) BulkDeleteSearchHistoryHandler(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value("userID").(uint64)
	var req dto.BulkDeleteSearchHistoryQueryParamsDTO

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.WeatherService.BulkDeleteSearchHistory(req.IDs, userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
