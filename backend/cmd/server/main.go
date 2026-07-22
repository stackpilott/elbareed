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
	// 1. Configure IMAP Host (Defaulting to Gmail)
	imapHost := os.Getenv("IMAP_HOST")
	if imapHost == "" {
		imapHost = "imap.gmail.com:993"
	}

	// 2. Initialize Adapters (The Real Connections)
	fmt.Printf("Initializing real IMAP connection to %s...\n", imapHost)
	store := adapter.NewInMemorySessionStore()
	verifier := adapter.NewIMAPAdapter(imapHost) // Now using the REAL adapter!

	// 3. Initialize Services (The Business Logic)
	authService := service.NewAuthService(verifier, store)

	// 4. Initialize Handlers (The Web Translators)
	authHandler := transportHttp.NewAuthHandler(authService)

	// 5. Set up the Router
	r := chi.NewRouter()

	// 6. Add Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	// 7. Map URLs to Handlers
	r.Post("/api/login", authHandler.Login)

	// 8. Start the Server
	port := ":8080"
	fmt.Printf("Web server is running on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
