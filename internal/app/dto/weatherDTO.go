package dto

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}
type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

type Clouds struct {
	All int `json:"all"`
}

type Sys struct {
	Type    int    `json:"type"`
	ID      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type OpenWeatherResponseDTO struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int       `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int64     `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int       `json:"timezone"`
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Cod        int       `json:"cod"`
}

type GetCurrentWeatherResponseDTO struct {
	Coord         Coord     `json:"coord"`
	Weather       []Weather `json:"weather"`
	Base          string    `json:"base"`
	Main          Main      `json:"main"`
	Visibility    int       `json:"visibility"`
	Wind          Wind      `json:"wind"`
	Clouds        Clouds    `json:"clouds"`
	Dt            int64     `json:"dt"`
	Sys           Sys       `json:"sys"`
	Timezone      int       `json:"timezone"`
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Cod           int       `json:"cod"`
	OpenWeatherID int64     `json:"open_weather_id"`
	CreatedBy     uint64    `json:"created_by"`
}

type GetSearchHistoryQueryParamsDTO struct {
	Cursor int `json:"cursor" validate:"number"`
	Limit  int `json:"limit" validate:"required,min=1,max=100"`
}

type GetSearchHistoryResponseDTO struct {
	Data       []*GetCurrentWeatherResponseDTO `json:"data"`
	NextCursor int                             `json:"next_cursor"`
}

type BulkDeleteSearchHistoryQueryParamsDTO struct {
	IDs []uint64 `json:"ids" validate:"required,min=1"`
}
