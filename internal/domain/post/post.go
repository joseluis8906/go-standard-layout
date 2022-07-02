package post

import "github.com/joseluis8906/go-standard-layout/pkg/ddd"

type (
	Post struct {
		ddd.Aggregate

		title string
		body  string
	}
)

func (p Post) Title() string {
	return p.title
}

func (p Post) Body() string {
	return p.body
}
