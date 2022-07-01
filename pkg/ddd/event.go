package ddd

import "time"

type (
	Event struct {
		ID         string    `json:"id"`
		Type       string    `json:"type"`
		OccurredOn time.Time `json:"occurred_on"`
	}
)
