package domain

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

// Suggestion is a user-submitted Promise that requires moderation
type Suggestion struct {
	ID           int            `json:"id"`
	CreatedAt    time.Time      `json:"created_at" db:"created_at"`
	AcceptedAt   pq.NullTime    `json:"-" db:"accepted_at"`
	AcceptedBy   sql.NullString `json:"-" db:"accepted_by"`
	DeletedAt    pq.NullTime    `json:"-" db:"deleted_at"`
	DeletedBy    sql.NullString `json:"-" db:"deleted_by"`
	PoliticianID string         `json:"politician" db:"politician_id"`
	Promise      string         `json:"promise"`
	Status       string         `json:"status"`
	StatusDetail string         `json:"status_detail" db:"status_detail"`
	Category     string         `json:"category"`
	SourceName   string         `json:"source_name" db:"source_name"`
	SourceLink   string         `json:"source_link" db:"source_link"`
}
