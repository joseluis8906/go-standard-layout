package queries

import (
	"context"
	stdhttp "net/http"
	"strconv"

	"github.com/joseluis8906/go-standard-layout/internal/domain/post"
	"github.com/joseluis8906/go-standard-layout/pkg/http"
	"github.com/joseluis8906/go-standard-layout/pkg/log"
)

type (
	GetAllPosts struct {
		Page int `json:"page"`
	}

	postResponse struct {
		ID    string `json:"id"`
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	GetAllPostHandler struct {
		PostFinder interface {
			GetAll(context.Context) ([]post.Post, error)
		}
	}
)

func (g GetAllPostHandler) do(ctx context.Context, query GetAllPosts) ([]postResponse, error) {
	data, err := g.PostFinder.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	posts := make([]postResponse, len(data))
	for i, post := range data {
		posts[i] = postResponse{
			ID:    post.ID().String(),
			Title: post.Title(),
			Body:  post.Body(),
		}
	}

	return posts, nil
}

func (g GetAllPostHandler) HandleFunc(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	ctx := r.Context()

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		log.Error(err)
		http.JSON(w, stdhttp.StatusBadRequest, nil, err)
		return
	}

	query := GetAllPosts{
		Page: page,
	}

	posts, err := g.do(ctx, query)
	if err != nil {
		log.Error(err)
		http.JSON(w, stdhttp.StatusBadRequest, nil, err)
		return
	}

	http.JSON(w, stdhttp.StatusOK, posts, nil)
}
