package domain

import "time"

// Suggestion is a user-submitted Promise that requires moderation
type Suggestion struct {
	ID           int       `json:"id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	PoliticianID string    `json:"politician" db:"politician_id"`
	Promise      string    `json:"promise"`
	Status       string    `json:"status"`
	StatusDetail string    `json:"status_detail" db:"status_detail"`
	Category     string    `json:"category"`
	SourceName   string    `json:"source_name" db:"source_name"`
	SourceLink   string    `json:"source_link" db:"source_link"`
}
