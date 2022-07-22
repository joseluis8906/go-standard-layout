package inmemory

import (
	"context"
	"sync"

	"github.com/joseluis8906/go-standard-layout/internal/domain/post"
	"github.com/joseluis8906/go-standard-layout/pkg/errors"
	"github.com/joseluis8906/go-standard-layout/pkg/eventbus"
)

const (
	// ErrNotFound ...
	ErrNotFound = errors.Error("error post not found")
)

type (
	// PostRepository ...
	PostRepository struct {
		data map[post.PostID]post.Post
		mux  sync.RWMutex
	}
)

// NewPostRepository ...
func NewPostRepository() *PostRepository {
	return &PostRepository{
		data: map[post.PostID]post.Post{},
	}
}

// Persist ...
func (r *PostRepository) Persist(ctx context.Context, p post.Post) (err error) {
	defer func() {
		if err != nil {
			eventbus.Discard(ctx, p.EventIDs())
			return
		}

		eventbus.Publish(ctx, p.EventIDs())
	}()

	r.mux.Lock()
	defer r.mux.Unlock()

	r.data[p.ID()] = p
	return nil
}

// GetAll ...
func (r *PostRepository) GetAll(ctx context.Context) ([]post.Post, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	posts := []post.Post{}
	for _, p := range r.data {
		posts = append(posts, p)
	}

	return posts, nil
}

// GetByID ...
func (r *PostRepository) GetByID(ctx context.Context, id post.PostID) (post.Post, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	p, ok := r.data[id]
	if !ok {
		return post.NoopPost(), ErrNotFound
	}

	return p, nil
}
