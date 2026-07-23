package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"webmail-backend/internal/adapter"
	"webmail-backend/internal/domain"

	"github.com/google/uuid"
)

type AuthService struct {
	verifier *adapter.IMAPAdapter
	store    *adapter.InMemorySessionStore
}

func NewAuthService(v *adapter.IMAPAdapter, s *adapter.InMemorySessionStore) *AuthService {
	return &AuthService{verifier: v, store: s}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	if err := s.verifier.Verify(ctx, email, password); err != nil {
		return "", err
	}

	token := uuid.New().String()
	session := &domain.Session{
		Token:     token,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}

	if err := s.store.Save(ctx, session); err != nil {
		return "", errors.New("failed to create session")
	}

	fmt.Printf("Got the token %s", token)
	return token, nil
}

func (s *AuthService) LoginWithGoogle(ctx context.Context, accessToken string) (string, error) {
	req, _ := http.NewRequestWithContext(ctx, "GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return "", errors.New("failed to verify google token")
	}
	defer resp.Body.Close()

	var userInfo struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return "", errors.New("failed to parse google user info")
	}

	if err := s.verifier.VerifyWithOAuth(ctx, userInfo.Email, accessToken); err != nil {
		return "", fmt.Errorf("imap oauth failed: %w", err)
	}

	token := uuid.New().String()
	session := &domain.Session{
		Token:       token,
		Email:       userInfo.Email,
		AccessToken: accessToken,
		CreatedAt:   time.Now(),
	}

	if err := s.store.Save(ctx, session); err != nil {
		return "", errors.New("failed to create session")
	}

	return token, nil
}
