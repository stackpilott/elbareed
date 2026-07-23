package adapter

import (
	"context"
	"fmt"

	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-sasl"
)

type IMAPAdapter struct {
	serverHost string
}

func NewIMAPAdapter(host string) *IMAPAdapter {
	return &IMAPAdapter{
		serverHost: host,
	}
}

func (a *IMAPAdapter) Verify(ctx context.Context, email, password string) error {
	// c, err := client.DialTLS(a.serverHost, nil)
	// if err != nil {
	// 	return fmt.Errorf("failed to connect to IMAP server: %w", err)
	// }
	// defer c.Logout()

	// if err := c.Login(email, password); err != nil {
	// 	return fmt.Errorf("invalid email or app password: %w", err)
	// }

	return nil
}

func (a *IMAPAdapter) VerifyWithOAuth(ctx context.Context, email, accessToken string) error {
	c, err := client.DialTLS(a.serverHost, nil)
	if err != nil {
		return fmt.Errorf("failed to connect to IMAP server: %w", err)
	}
	defer c.Logout()

	// Updated to use the modern OAUTHBEARER standard required by the go-sasl package
	saslClient := sasl.NewOAuthBearerClient(&sasl.OAuthBearerOptions{
		Username: email,
		Token:    accessToken,
	})

	if err := c.Authenticate(saslClient); err != nil {
		return fmt.Errorf("google oauth imap authentication failed: %w", err)
	}

	return nil
}
