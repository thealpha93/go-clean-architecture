package web

import (
	"log"
	"net/http"
	"test-server-app/internal/app/usecases/auth"
	"test-server-app/internal/app/usecases/weather"

	"github.com/rs/cors"
)

func StartServer(authService *auth.AuthService, weatherService *weather.WeatherService) {
	r := setupRouter(authService, weatherService)

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders: []string{"*"},
		//ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}).Handler(r)

	log.Fatal(http.ListenAndServe(":6261", handler))
}
