package events

import "github.com/joseluis8906/go-standard-layout/pkg/ddd"

const (
	PostCreatedType = "[POST_CREATED]"
)

type (
	PostCreated struct {
		ddd.Event
		Attributes struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		} `json:"attributes"`
	}
)

func NewPostCreated() PostCreated {
	e := PostCreated{}
	e.Event = ddd.NewEvent()
	e.Event.Type = PostCreatedType

	return e
}
