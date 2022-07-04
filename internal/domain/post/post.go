package post

import (
	"github.com/joseluis8906/go-standard-layout/pkg/ddd"
)

type (
	Post struct {
		ddd.Aggregate

		id    PostID
		title string
		body  string
	}
)

func NoopPost() Post {
	return Post{id: NoopPostID()}
}

func (p Post) IsZero() bool {
	return p.id.IsZero()
}

func (p Post) ID() PostID {
	return p.id
}

func (p Post) Title() string {
	return p.title
}

func (p Post) Body() string {
	return p.body
}
