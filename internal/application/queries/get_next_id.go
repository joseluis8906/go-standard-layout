package queries

import (
	"context"
	stdhttp "net/http"

	"github.com/google/uuid"
	"github.com/joseluis8906/go-standard-layout/pkg/http"
)

const (
	nilID = "00000000-0000-0000-0000-000000000000"
)

type (
	GetNextIDHandler struct{}
)

func (g GetNextIDHandler) do(ctx context.Context) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nilID, err
	}

	return id.String(), nil
}

func (g GetNextIDHandler) HandleFunc(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	ctx := r.Context()
	data, err := g.do(ctx)
	if err != nil {
		http.JSON(w, stdhttp.StatusInternalServerError, nil, err)
		return
	}

	http.JSON(w, stdhttp.StatusOK, data, nil)
}
