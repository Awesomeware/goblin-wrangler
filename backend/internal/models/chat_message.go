package models

import "time"

type ChatMessage struct {
	Username string    `json:"username"`
	Message  string    `json:"message"`
	PostedAt time.Time `json:"posted_at"`
}
