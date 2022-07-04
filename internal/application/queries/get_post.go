package queries

import (
	"context"
	stdhttp "net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joseluis8906/go-standard-layout/internal/domain/post"
	"github.com/joseluis8906/go-standard-layout/pkg/http"
	"github.com/joseluis8906/go-standard-layout/pkg/log"
)

type (
	GetPost struct {
		ID string `json:"id"`
	}

	GetPostHandler struct {
		PostFinder interface {
			GetByID(context.Context, post.PostID) (post.Post, error)
		}
	}
)

func (g GetPostHandler) do(ctx context.Context, query GetPost) (post.Post, error) {
	id, err := post.ParsePostID(query.ID)
	if err != nil {
		return post.NoopPost(), err
	}

	return g.PostFinder.GetByID(ctx, id)
}

func (g GetPostHandler) HandleFunc(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	ctx := r.Context()
	dataID := chi.URLParam(r, "id")

	id, err := post.ParsePostID(dataID)
	log.Info(id)
	if err != nil {
		http.JSON(w, stdhttp.StatusBadRequest, nil, err)
		return
	}

	p, err := g.PostFinder.GetByID(ctx, id)
	if err != nil {
		http.JSON(w, stdhttp.StatusBadRequest, nil, err)
		return
	}

	pr := postResponse{
		ID:    p.ID().String(),
		Title: p.Title(),
		Body:  p.Body(),
	}

	http.JSON(w, stdhttp.StatusOK, pr, nil)
}
