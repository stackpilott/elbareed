/*this is in internal/adapter/session_Store.go */

package adapter

import (
	"context"
	"errors"
	"sync"
	"webmail-backend/internal/domain"
)

type InMemorySessionStore struct {
	mu       sync.RWMutex
	sessions map[string]*domain.Session
}

func NewInMemorySessionStore() *InMemorySessionStore {
	return &InMemorySessionStore{
		sessions: make(map[string]*domain.Session),
	}
}

// Save securely locks the map, writes the session, and unlocks it.
func (s *InMemorySessionStore) Save(ctx context.Context, session *domain.Session) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.sessions[session.Token] = session
	return nil
}

// Get uses a read-lock, allowing multiple requests to read simultaneously safely.
func (s *InMemorySessionStore) Get(ctx context.Context, token string) (*domain.Session, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	session, exists := s.sessions[token]
	if !exists {
		return nil, errors.New("session not found")
	}
	return session, nil
}
