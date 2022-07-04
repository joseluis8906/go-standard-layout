package queries

import (
	"context"
	"fmt"
	stdhttp "net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joseluis8906/go-standard-layout/internal/domain/post"
	"github.com/joseluis8906/go-standard-layout/pkg/http"
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

	query := GetPost{
		ID: chi.URLParam(r, "id"),
	}

	p, err := g.do(ctx, query)
	if err != nil {
		http.JSON(w, stdhttp.StatusBadRequest, nil, err)
		return
	}

	pr := postResponse{
		ID:    p.ID().String(),
		Title: p.Title(),
		Body:  p.Body(),
	}

	if p.IsZero() {
		http.JSON(w, stdhttp.StatusNotFound, nil, fmt.Errorf("product not found"))
		return
	}

	http.JSON(w, stdhttp.StatusOK, pr, nil)
}
