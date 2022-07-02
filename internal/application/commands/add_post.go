package commands

import (
	"context"
	"encoding/json"
	stdhttp "net/http"

	"github.com/joseluis8906/go-standard-layout/internal/domain/post"
	"github.com/joseluis8906/go-standard-layout/pkg/http"
	"github.com/joseluis8906/go-standard-layout/pkg/log"
)

type (
	AddPost struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	AddPostHandler struct {
		PostPersistor interface {
			Persist(context.Context, post.Post) error
		}
	}
)

func (a AddPostHandler) do(ctx context.Context, command AddPost) error {
	p, err := post.NewBuilder().
		WithTitle(command.Title).
		WithBody(command.Body).
		Build()

	if err != nil {
		return err
	}

	return a.PostPersistor.Persist(ctx, p)
}

func (g AddPostHandler) HandleFunc(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	ctx := r.Context()

	var command AddPost

	err := json.NewDecoder(r.Body).Decode(&command)
	if err != nil {
		log.Error(err)
		http.JSON(w, stdhttp.StatusBadRequest, nil, err)
		return
	}

	err = g.do(ctx, command)
	if err != nil {
		log.Error(err)
		http.JSON(w, stdhttp.StatusBadRequest, nil, err)
		return
	}

	http.JSON(w, stdhttp.StatusCreated, nil, nil)
}
