package port

import (
	"context"
	"webmail-backend/internal/domain"
)

// MailReader defines how we fetch mail, decoupled from the actual IMAP protocol.
type MailReader interface {
	ListMailboxes(ctx context.Context, session *domain.Session) ([]domain.Mailbox, error)
	GetMessages(ctx context.Context, session *domain.Session, mailboxName string) ([]domain.Message, error)
}

// MailSender defines how we send mail, decoupled from the actual SMTP protocol.
type MailSender interface {
	Send(ctx context.Context, session *domain.Session, msg *domain.Message) error
}

// CredentialVerifier checks if the user's email/password are valid upstream.
type CredentialVerifier interface {
	Verify(ctx context.Context, email, password string) error
}

// SessionStore defines how we keep users logged in (Memory, Redis, etc.)
type SessionStore interface {
	Save(ctx context.Context, session *domain.Session) error
	Get(ctx context.Context, token string) (*domain.Session, error)
}
