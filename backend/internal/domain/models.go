package domain

import "time"

type Mailbox struct {
	Name        string
	UnreadCount int
	TotalCount  int
}

type Message struct {
	ID      string
	Subject string
	From    string
	To      []string
	Date    time.Time
	Body    string
	Flags   []string
}

type Session struct {
	Token       string
	Email       string
	Password    string
	AccessToken string
	CreatedAt   time.Time
}
