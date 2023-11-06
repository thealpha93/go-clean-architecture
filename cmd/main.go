package main

import (
	"log"

	"test-server-app/internal/app/config"

	"test-server-app/internal/app/infrastructure/database"
	"test-server-app/internal/app/infrastructure/web"

	"test-server-app/internal/app/usecases/auth"
	"test-server-app/internal/app/usecases/weather"
)

func main() {
	// Need to set up a common response handle structure
	// Since it's only a few Api's, choosing not to do that now

	db, err := config.InitDB(config.AppConfig.DatabaseURL)

	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlDB.Close()

	userRepo := database.NewUserRepoImpl(db)
	authService := auth.NewAuthService(userRepo)

	weatherRepo := database.NewWeatherRepoImpl(db)
	weatherService := weather.NewWeatherService(weatherRepo)

	// This is better moved to config and have a struct for all the services to avoid too much params here
	// But since it's only two, I'm going to let this be for now
	web.StartServer(authService, weatherService)

}
