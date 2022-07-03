package ddd

import (
	"time"

	"github.com/google/uuid"
)

type (
	Event struct {
		ID         string    `json:"id"`
		Type       string    `json:"type"`
		OccurredOn time.Time `json:"occurred_on"`
	}
)

func NewEvent() Event {
	id, _ := uuid.NewRandom()

	return Event{
		ID:         id.String(),
		OccurredOn: time.Now(),
	}
}
