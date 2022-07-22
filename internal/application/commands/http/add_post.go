package http

import (
	"encoding/json"
	stdhttp "net/http"

	"github.com/joseluis8906/go-standard-layout/internal/application/commands"
	"github.com/joseluis8906/go-standard-layout/pkg/http"
	"github.com/joseluis8906/go-standard-layout/pkg/log"
)

// AddPost ...
type AddPost struct {
	Handler commands.AddPostHandler
}

// HandleFunc ...
func (a AddPost) HandleFunc(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	ctx := r.Context()

	var command commands.AddPost

	err := json.NewDecoder(r.Body).Decode(&command)
	if err != nil {
		log.Errorf("error trying to handle AddPost command, casting command json: %v", err)
		http.JSON(w, stdhttp.StatusBadRequest, nil, err)
		return
	}

	err = a.Handler.Do(ctx, command)
	if err != nil {
		log.Errorf("error trying to handle AddPost command: %v", err)
		http.JSON(w, stdhttp.StatusBadRequest, nil, err)
		return
	}

	http.JSON(w, stdhttp.StatusCreated, nil, nil)
}
