package queries

import (
	"context"

	"github.com/google/uuid"
)

const (
	nilID = "00000000-0000-0000-0000-000000000000"
)

type (
	// GetNextIDHandler ...
	GetNextIDHandler struct{}
)

// Do ...
func (g GetNextIDHandler) Do(ctx context.Context) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nilID, err
	}

	return id.String(), nil
}
