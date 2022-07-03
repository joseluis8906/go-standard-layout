package inmemory

import (
	"context"

	"github.com/joseluis8906/go-standard-layout/internal/domain/post"
	"github.com/joseluis8906/go-standard-layout/pkg/eventbus"
)

type (
	PostRepository struct {
		data []post.Post
	}
)

func NewPostRepository() *PostRepository {
	return &PostRepository{
		data: []post.Post{},
	}
}

func (r *PostRepository) Persist(ctx context.Context, p post.Post) (err error) {
	defer func() {
		if err != nil {
			eventbus.Discard(ctx, p.EventIDs())
			return
		}

		eventbus.Publish(ctx, p.EventIDs())
	}()

	r.data = append(r.data, p)
	return nil
}

func (r *PostRepository) GetAll(ctx context.Context) ([]post.Post, error) {
	return r.data, nil
}
