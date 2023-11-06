package config

import (
	"github.com/spf13/viper"
)

type AppConfiguration struct {
	SERVER_PORT           int
	DatabaseURL           string
	OPEN_WEATHER_API_KEY  string
	OPEN_WEATHER_BASE_URL string
	JWT_SECRET            string
	REFRESH_TOKEN_SECRET  string
}

func LoadConfig() *AppConfiguration {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/Users/aakash/Documents/learn/go/test-server-repo")
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()
	viper.ReadInConfig()

	config := &AppConfiguration{}

	err := viper.Unmarshal(config)

	if err != nil {
		panic(err)
	}
	return config
}

var AppConfig = LoadConfig()
