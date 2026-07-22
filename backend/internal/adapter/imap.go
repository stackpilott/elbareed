package adapter

import (
	"context"
	"fmt"

	"github.com/emersion/go-imap/client"
)

type IMAPAdapter struct {
	serverHost string
}

func NewIMAPAdapter(host string) *IMAPAdapter {
	return &IMAPAdapter{
		serverHost: host,
	}
}

func (a *IMAPAdapter) connectAndLogin(email, password string) (*client.Client, error) {
	c, err := client.DialTLS(a.serverHost, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to IMAP server: %w", err)
	}

	if err := c.Login(email, password); err != nil {
		return nil, fmt.Errorf("invalid email or app password: %w", err)
	}

	return c, nil
}

func (a *IMAPAdapter) Verify(ctx context.Context, email, password string) error {
	c, err := a.connectAndLogin(email, password)
	if err != nil {
		return err
	}

	defer c.Logout()
	return nil
}
