package main

import (
	"context"
	"fmt"
	"log"

	"webmail-backend/internal/adapter"
	"webmail-backend/internal/service"
)

// 1. Create a quick Mock Adapter
type MockVerifier struct{}

func (m *MockVerifier) Verify(ctx context.Context, email, password string) error {
	if email == "test@example.com" && password == "secret" {
		return nil // Success!
	}
	return fmt.Errorf("invalid credentials")
}

func main() {
	ctx := context.Background()

	// 2. Initialize our dependencies (Adapters)
	fmt.Println("Starting up the backend core...")
	store := adapter.NewInMemorySessionStore()
	verifier := &MockVerifier{}

	// 3. Inject them into our Service
	authService := service.NewAuthService(verifier, store)

	// 4. Test a failed login
	fmt.Println("\n--- Testing Failed Login ---")
	_, err := authService.Login(ctx, "wrong@example.com", "badpass")
	if err != nil {
		fmt.Printf("Expected error caught: %v\n", err)
	}

	// 5. Test a successful login
	fmt.Println("\n--- Testing Successful Login ---")
	session, err := authService.Login(ctx, "test@example.com", "secret")
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Success! Generated Secure Token: %s\n", session.Token)

	// 6. Verify the session was actually saved to our store
	fmt.Println("\n--- Verifying Session Storage ---")
	savedSession, err := store.Get(ctx, session.Token)
	if err != nil {
		log.Fatalf("Failed to retrieve session: %v", err)
	}
	fmt.Printf("Retrieved session for: %s\n", savedSession.Email)
}
