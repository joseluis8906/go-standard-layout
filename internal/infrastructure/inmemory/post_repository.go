package inmemory

import (
	"context"

	"github.com/joseluis8906/go-standard-layout/internal/domain/post"
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

func (r *PostRepository) Persist(ctx context.Context, p post.Post) error {
	r.data = append(r.data, p)

	return nil
}

func (r *PostRepository) GetAll(ctx context.Context) ([]post.Post, error) {
	return r.data, nil
}
