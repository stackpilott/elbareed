package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"webmail-backend/internal/adapter"
	"webmail-backend/internal/service"
	transportHttp "webmail-backend/internal/transport/http" // Alias to avoid conflict with the standard "http" package
)

// We keep our MockVerifier so we can test the API immediately
type MockVerifier struct{}

func (m *MockVerifier) Verify(ctx context.Context, email, password string) error {
	if email == "test@example.com" && password == "secret" {
		return nil
	}
	return fmt.Errorf("invalid credentials")
}

func main() {
	// 1. Initialize Adapters (The Database/External connections)
	store := adapter.NewInMemorySessionStore()
	verifier := &MockVerifier{} // We will swap this out for the real IMAP adapter later!

	// 2. Initialize Services (The Business Logic)
	authService := service.NewAuthService(verifier, store)

	// 3. Initialize Handlers (The Web Translators)
	authHandler := transportHttp.NewAuthHandler(authService)

	// 4. Set up the Router
	r := chi.NewRouter()

	// 5. Add Middleware
	r.Use(middleware.Logger)    // Logs every request to the terminal
	r.Use(middleware.Recoverer) // Prevents the server from crashing if there is a panic
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Allow our Vite React frontend
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	// 6. Map URLs to Handlers
	r.Post("/api/login", authHandler.Login)

	// 7. Start the Server
	port := ":8080"
	fmt.Printf("Web server is running on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
