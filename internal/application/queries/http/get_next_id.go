package http

import (
	stdhttp "net/http"

	"github.com/joseluis8906/go-standard-layout/internal/application/queries"
	"github.com/joseluis8906/go-standard-layout/pkg/http"
	"github.com/joseluis8906/go-standard-layout/pkg/log"
)

// GetNextID ...
type GetNextID struct {
	Handler queries.GetNextIDHandler
}

// HandleFunc ...
func (gni GetNextID) HandleFunc(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	ctx := r.Context()
	id, err := gni.Handler.Do(ctx)
	if err != nil {
		log.Errorf("error trying to handle GetNextID command: %v", err)
		http.JSON(w, stdhttp.StatusInternalServerError, nil, err)
		return
	}

	http.JSON(w, stdhttp.StatusOK, id, nil)
}
