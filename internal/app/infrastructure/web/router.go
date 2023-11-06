package web

import (
	"net/http"
	"test-server-app/internal/app/usecases/auth"
	"test-server-app/internal/app/usecases/weather"

	"github.com/gorilla/mux"
)

func setupRouter(authService *auth.AuthService, weatherService *weather.WeatherService) *mux.Router {
	router := mux.NewRouter()

	// headers := handlers.AllowedHeaders([]string{"*"})
	// methods := handlers.AllowedMethods([]string{"*"})
	// origins := handlers.AllowedOrigins([]string{"*"}) // Allowing all origins since it's a test task

	// router.Use(handlers.CORS(headers, methods, origins))

	authHandler := auth.NewAuthHandler(authService)

	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/login", authHandler.LoginHandler).Methods(http.MethodPost)
	authRouter.HandleFunc("/signup", authHandler.CreateUserHandler).Methods(http.MethodPost)

	// Authenticated routes
	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.Use(AuthMiddleware(authService))
	userRouter.HandleFunc("/me", authHandler.GetCurrentUserHandler)

	weatherHandler := weather.NewWeatherHandler(weatherService)

	weatherRouter := router.PathPrefix("/weather").Subrouter()
	weatherRouter.Use(AuthMiddleware(authService))

	weatherRouter.HandleFunc("/current", weatherHandler.GetCurrentWeatherHandler)
	weatherRouter.HandleFunc("/search-history", weatherHandler.GetSearchHistoryHandler)
	// bulk-delete will handle both delete and bulk delete
	weatherRouter.HandleFunc("/bulk-delete", weatherHandler.BulkDeleteSearchHistoryHandler).Methods(http.MethodPost)

	return router
}
