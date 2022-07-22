package queries

import (
	"context"

	"github.com/joseluis8906/go-standard-layout/internal/domain/post"
)

type (
	// GetPost ...
	GetPost struct {
		ID string `json:"id"`
	}

	// GetPostHandler ...
	GetPostHandler struct {
		PostFinder interface {
			GetByID(context.Context, post.PostID) (post.Post, error)
		}
	}
)

// Do ...
func (g GetPostHandler) Do(ctx context.Context, query GetPost) (post.Post, error) {
	id, err := post.ParsePostID(query.ID)
	if err != nil {
		return post.NoopPost(), err
	}

	return g.PostFinder.GetByID(ctx, id)
}
