package http

import (
	"encoding/json"
	"net/http"
	"webmail-backend/internal/service"
)

// AuthHandler wraps our service so web requests can use it
type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: s,
	}
}

// LoginRequest is exactly what the React frontend will send us
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login is the actual HTTP endpoint function
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// 1. Decode the JSON from the frontend
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}

	// 2. Call the core service (the exact same function you just tested!)
	session, err := h.authService.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		// If credentials fail, return a 401 Unauthorized
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// 3. Send the secure token back to the frontend as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": session.Token,
	})
}
