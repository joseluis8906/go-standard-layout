package commands

import (
	"context"

	"github.com/joseluis8906/go-standard-layout/internal/domain/post"
)

type (
	// AddPost ...
	AddPost struct {
		ID    string `json:"id"`
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	// AddPostHandler ...
	AddPostHandler struct {
		PostPersistor interface {
			Persist(context.Context, post.Post) error
		}
	}
)

// Do ...
func (a AddPostHandler) Do(ctx context.Context, command AddPost) error {
	id, err := post.ParsePostID(command.ID)
	if err != nil {
		return err
	}

	p, err := post.NewBuilder().
		WithID(id).
		WithTitle(command.Title).
		WithBody(command.Body).
		Build(true)

	if err != nil {
		return err
	}

	return a.PostPersistor.Persist(ctx, p)
}
