package domain

import "time"

// PromiseStatusName for discrete statuses
type PromiseStatusName string

const (
	// NotStarted - The Promise hasn't been worked on
	NotStarted PromiseStatusName = "not-started"

	// InProgress - The Promise is being implemented now
	InProgress PromiseStatusName = "in-progress"

	// Accomplished - The Promise has been fulfilled
	Accomplished PromiseStatusName = "accomplished"

	// Failed - The Promise has been broken
	Failed PromiseStatusName = "failed"
)

// PromiseStatus is the status of a Promise's implementation at a point in time
type PromiseStatus struct {
	ID        int               `json:"-"`
	PromiseID int               `json:"-" db:"promise_id"`
	Promise   *Promise          `json:"-"`
	Name      PromiseStatusName `json:"name"`
	UpdatedOn time.Time         `json:"date" db:"updated_on"`
	Detail    string            `json:"detail"`
	Sources   []Source          `json:"sources,omitempty"`
}

// ValidStatuses returns the names of the valid statuses for a promise to be in
func ValidStatuses() []PromiseStatusName {
	return []PromiseStatusName{
		NotStarted,
		InProgress,
		Accomplished,
		Failed,
	}
}
