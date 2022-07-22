package queries

import (
	"context"

	"github.com/joseluis8906/go-standard-layout/internal/domain/post"
)

type (
	// GetAllPosts ...
	GetAllPosts struct {
		Page int
	}

	// GetAllPostHandler ...
	GetAllPostHandler struct {
		PostFinder interface {
			GetAll(context.Context) ([]post.Post, error)
		}
	}
)

// Do ...
func (g GetAllPostHandler) Do(ctx context.Context, query GetAllPosts) ([]post.Post, error) {
	return g.PostFinder.GetAll(ctx)
}
