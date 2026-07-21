package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"webmail-backend/internal/domain"
	"webmail-backend/internal/port"
)

type AuthService struct {
	verifier port.CredentialVerifier
	store    port.SessionStore
}

// NewAuthService injects our port dependencies into the service.
func NewAuthService(v port.CredentialVerifier, s port.SessionStore) *AuthService {
	return &AuthService{
		verifier: v,
		store:    s,
	}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*domain.Session, error) {
	// 1. Verify credentials with the upstream server
	if err := s.verifier.Verify(ctx, email, password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// 2. Generate a secure random session token
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return nil, errors.New("failed to generate token")
	}
	token := hex.EncodeToString(bytes)

	// 3. Create the domain session object
	session := &domain.Session{
		Token:    token,
		Email:    email,
		Password: password,
	}

	// 4. Save to our session store
	if err := s.store.Save(ctx, session); err != nil {
		return nil, errors.New("failed to save session")
	}

	return session, nil
}
