package web

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"test-server-app/internal/app/usecases/auth"

	"github.com/gorilla/mux"
)

func AuthMiddleware(authService *auth.AuthService) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract token from Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header is required", http.StatusUnauthorized)
				return
			}

			bearerToken := strings.Split(authHeader, " ")
			if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
				http.Error(w, "Invalid token format", http.StatusUnauthorized)
				return
			}

			tokenString := bearerToken[1]
			userID, err := authService.ParseToken(tokenString)
			if err != nil {
				http.Error(w, fmt.Sprintf("Unauthorized: %v", err), http.StatusUnauthorized)
				return
			}

			// Store the user ID in the request context
			ctx := context.WithValue(r.Context(), "userID", userID)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
