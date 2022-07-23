package http

import (
	"fmt"
	stdhttp "net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joseluis8906/go-standard-layout/internal/application/queries"
	"github.com/joseluis8906/go-standard-layout/pkg/http"
	"github.com/joseluis8906/go-standard-layout/pkg/log"
)

type (
	// GetPost ...
	GetPost struct {
		Handler queries.GetPostHandler
	}

	getPostResponse struct {
		ID    string `json:"id"`
		Title string `json:"title"`
		Body  string `json:"body"`
	}
)

// HandleFunc ...
func (gp GetPost) HandleFunc(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	ctx := r.Context()

	query := queries.GetPost{
		ID: chi.URLParam(r, "id"),
	}

	p, err := gp.Handler.Do(ctx, query)
	if err != nil {
		log.Logger().Errorf("error trying to handle GetPost command: %v", err)
		http.JSON(w, stdhttp.StatusBadRequest, nil, err)
		return
	}

	pr := getPostResponse{
		ID:    p.ID().String(),
		Title: p.Title(),
		Body:  p.Body(),
	}

	if p.IsZero() {
		log.Logger().Errorf("error trying to handle GetPost command, product is zero: %v", p)
		http.JSON(w, stdhttp.StatusNotFound, nil, fmt.Errorf("product not found"))
		return
	}

	http.JSON(w, stdhttp.StatusOK, pr, nil)
}
