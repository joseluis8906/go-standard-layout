package inmemory

import (
	"context"

	"github.com/joseluis8906/go-standard-layout/internal/domain/post"
	"github.com/joseluis8906/go-standard-layout/pkg/eventbus"
)

type (
	PostRepository struct {
		data map[post.PostID]post.Post
	}
)

func NewPostRepository() *PostRepository {
	return &PostRepository{
		data: map[post.PostID]post.Post{},
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

	r.data[p.ID()] = p
	return nil
}

func (r *PostRepository) GetAll(ctx context.Context) ([]post.Post, error) {
	posts := []post.Post{}
	for _, p := range r.data {
		posts = append(posts, p)
	}

	return posts, nil
}

func (r *PostRepository) GetByID(ctx context.Context, id post.PostID) (post.Post, error) {
	return r.data[id], nil
}
