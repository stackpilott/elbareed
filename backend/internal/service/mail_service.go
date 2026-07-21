package service

import (
	"context"
	"errors"
	"webmail-backend/internal/domain"
	"webmail-backend/internal/port"
)

type MailService struct {
	reader port.MailReader
	sender port.MailSender
	store  port.SessionStore
}

func NewMailService(r port.MailReader, snd port.MailSender, st port.SessionStore) *MailService {
	return &MailService{
		reader: r,
		sender: snd,
		store:  st,
	}
}

func (s *MailService) getSession(ctx context.Context, token string) (*domain.Session, error) {
	session, err := s.store.Get(ctx, token)
	if err != nil {
		return nil, errors.New("Unauthorized or session expired")
	}
	return session, nil
}

func (s *MailService) ListMailboxes(ctx context.Context, token string) ([]domain.Mailbox, error) {
	session, err := s.getSession(ctx, token)
	if err != nil {
		return nil, err
	}

	return s.reader.ListMailboxes(ctx, session)
}

func (s *MailService) GetMessages(ctx context.Context, token string, mailboxName string) ([]domain.Message, error) {
	session, err := s.getSession(ctx, token)
	if err != nil {
		return nil, err
	}

	return s.reader.GetMessages(ctx, session, mailboxName)
}
