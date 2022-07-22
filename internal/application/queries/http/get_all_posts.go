package http

import (
	stdhttp "net/http"
	"strconv"

	"github.com/joseluis8906/go-standard-layout/internal/application/queries"
	"github.com/joseluis8906/go-standard-layout/pkg/http"
	"github.com/joseluis8906/go-standard-layout/pkg/log"
)

// GetAllPost ...
type (
	GetAllPost struct {
		Handler queries.GetAllPostHandler
	}

	// getAllPostResponse ...
	getAllPostResponse struct {
		ID    string `json:"id"`
		Title string `json:"title"`
		Body  string `json:"body"`
	}
)

// HandleFunc ...
func (gap GetAllPost) HandleFunc(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	ctx := r.Context()

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		log.Errorf("error trying to handle GetAllPosts, casting page query param to int: %v", err)
		http.JSON(w, stdhttp.StatusBadRequest, nil, err)
		return
	}

	query := queries.GetAllPosts{
		Page: page,
	}

	posts, err := gap.Handler.Do(ctx, query)
	if err != nil {
		log.Errorf("error trying to handle GetAllPosts command: %v", err)
		http.JSON(w, stdhttp.StatusBadRequest, nil, err)
		return
	}

	res := make([]getAllPostResponse, len(posts))
	for i, post := range posts {
		res[i] = getAllPostResponse{
			ID:    post.ID().String(),
			Title: post.Title(),
			Body:  post.Body(),
		}
	}

	http.JSON(w, stdhttp.StatusOK, res, nil)
}
