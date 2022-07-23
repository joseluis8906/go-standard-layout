package http

import (
	"encoding/json"
	stdhttp "net/http"

	"github.com/joseluis8906/go-standard-layout/internal/application/commands"
	"github.com/joseluis8906/go-standard-layout/pkg/http"
	"github.com/joseluis8906/go-standard-layout/pkg/log"
)

// AddPost ...
type (
	addPostRequest struct {
		ID    string `json:"id"`
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	AddPost struct {
		Handler commands.AddPostHandler
	}
)

// HandleFunc ...
func (a AddPost) HandleFunc(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	ctx := r.Context()

	var req addPostRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Logger().Errorf("error trying to handle AddPost command, casting req json: %v", err)
		http.JSON(w, stdhttp.StatusBadRequest, nil, err)
		return
	}

	command := commands.AddPost{
		ID:    req.ID,
		Title: req.Title,
		Body:  req.Body,
	}

	err = a.Handler.Do(ctx, command)
	if err != nil {
		log.Logger().Errorf("error trying to handle AddPost command: %v", err)
		http.JSON(w, stdhttp.StatusBadRequest, nil, err)
		return
	}

	http.JSON(w, stdhttp.StatusCreated, nil, nil)
}
