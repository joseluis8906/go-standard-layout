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

	getAllPostResponse struct {
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

func (g GetAllPostHandler) do(ctx context.Context, query GetAllPosts) ([]getAllPostResponse, error) {
	data, err := g.PostFinder.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	posts := make([]getAllPostResponse, len(data))
	for i, post := range data {
		posts[i] = getAllPostResponse{
			ID:    post.ID().String(),
			Title: post.Title(),
			Body:  post.Body(),
		}
	}

	return posts, nil
}

func (g GetAllPostHandler) HandleFunc(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	ctx := r.Context()
	log.Info("testing logs")

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		log.Errorf("error trying to handle GetAllPosts, casting page query param to int: %v", err)
		http.JSON(w, stdhttp.StatusBadRequest, nil, err)
		return
	}

	query := GetAllPosts{
		Page: page,
	}

	posts, err := g.do(ctx, query)
	if err != nil {
		log.Errorf("error trying to handle GetAllPosts command: %v", err)
		http.JSON(w, stdhttp.StatusBadRequest, nil, err)
		return
	}

	http.JSON(w, stdhttp.StatusOK, posts, nil)
}
