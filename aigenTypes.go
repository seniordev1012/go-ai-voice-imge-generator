package main

import "time"

// Message represents a message in the chat
type Message struct {
	ID        int
	Sender    string
	Content   string
	CreatedAt time.Time
}

type User struct {
	ID       int
	Username string
	Password string
	Token    string
}
