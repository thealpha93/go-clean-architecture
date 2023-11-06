package auth

import (
	"encoding/json"
	"net/http"
	"test-server-app/internal/app/dto"

	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	AuthService *AuthService
}

func NewAuthHandler(authService *AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

var validate = validator.New()

func (h *AuthHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser dto.UserRegistrationDTO
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := validate.Struct(newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdUser, err := h.AuthService.CreateUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(createdUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginCreds dto.UserLoginRequestDTO
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&loginCreds); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if err := validate.Struct(loginCreds); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.AuthService.Authenticate(loginCreds.Email, loginCreds.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func (h *AuthHandler) GetCurrentUserHandler(w http.ResponseWriter, r *http.Request) {

	// Retrieve the user ID from the request context
	userID, _ := r.Context().Value("userID").(uint64)

	result, err := h.AuthService.GetUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}
