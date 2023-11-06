package model

import (
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

type WeatherSearchHistory struct {
	gorm.Model
	ID            uint64          `gorm:"primaryKey;autoIncrement:true;"`
	Coord         Coord           `gorm:"type:jsonb"`
	Weather       json.RawMessage `gorm:"type:jsonb"`
	Base          string
	Main          Main `gorm:"type:jsonb"`
	Visibility    int
	Wind          Wind   `gorm:"type:jsonb"`
	Clouds        Clouds `gorm:"type:jsonb"`
	Dt            int64
	Sys           Sys `gorm:"type:jsonb"`
	Timezone      int
	OpenWeatherID int64
	Name          string
	Cod           int
	CreatedBy     uint64
	CreatedAt     time.Time
	UpdatedAt     time.Time

	User UserModel `gorm:"ForeignKey:CreatedBy;"`
}

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
	FeelsLike float64 `gorm:"column:feels_like" json:"feels_like"`
	TempMin   float64 `gorm:"column:temp_min" json:"temp_min"`
	TempMax   float64 `gorm:"column:temp_max" json:"temp_max"`
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

// Implement the database/sql.Scanner interface for *model.Coord
func (c *Coord) Scan(value interface{}) error {
	// Perform the necessary type conversion from driver.Value to *model.Coord
	if jsonbData, ok := value.([]byte); ok {
		// Unmarshal the JSONB value into *model.Coord
		err := json.Unmarshal(jsonbData, c)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("unsupported Scan, storing driver.Value type []uint8 into type *model.Coord")
}

func (c *Main) Scan(value interface{}) error {
	// Perform the necessary type conversion from driver.Value to *model.Coord
	if jsonbData, ok := value.([]byte); ok {
		// Unmarshal the JSONB value into *model.Coord
		err := json.Unmarshal(jsonbData, c)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("unsupported Scan, storing driver.Value type []uint8 into type *model.Coord")
}

func (c *Wind) Scan(value interface{}) error {
	// Perform the necessary type conversion from driver.Value to *model.Coord
	if jsonbData, ok := value.([]byte); ok {
		// Unmarshal the JSONB value into *model.Coord
		err := json.Unmarshal(jsonbData, c)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("unsupported Scan, storing driver.Value type []uint8 into type *model.Coord")
}

func (c *Clouds) Scan(value interface{}) error {
	// Perform the necessary type conversion from driver.Value to *model.Coord
	if jsonbData, ok := value.([]byte); ok {
		// Unmarshal the JSONB value into *model.Coord
		err := json.Unmarshal(jsonbData, c)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("unsupported Scan, storing driver.Value type []uint8 into type *model.Coord")
}

func (c *Sys) Scan(value interface{}) error {
	// Perform the necessary type conversion from driver.Value to *model.Coord
	if jsonbData, ok := value.([]byte); ok {
		// Unmarshal the JSONB value into *model.Coord
		err := json.Unmarshal(jsonbData, c)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("unsupported Scan, storing driver.Value type []uint8 into type *model.Coord")
}
