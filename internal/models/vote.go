package models

import "time"

type Vote struct {
	PollId      *string    `json:"poll_id,omitempty"`
	UserId      *string    `json:"user_id,omitempty"`
	Entry       Fields     `json:"entry,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}
