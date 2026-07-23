package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"webmail-backend/internal/adapter"
	"webmail-backend/internal/service"
	transportHttp "webmail-backend/internal/transport/http"
)

func main() {
	imapHost := os.Getenv("IMAP_HOST")
	if imapHost == "" {
		imapHost = "imap.gmail.com:993"
	}

	fmt.Printf("Initializing real IMAP connection to %s...\n", imapHost)
	store := adapter.NewInMemorySessionStore()
	verifier := adapter.NewIMAPAdapter(imapHost)

	authService := service.NewAuthService(verifier, store)
	authHandler := transportHttp.NewAuthHandler(authService)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	r.Post("/api/login", authHandler.Login)
	r.Post("/api/auth/google", authHandler.GoogleLogin)

	port := ":8080"
	fmt.Printf("Web server is running on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
