package post

import (
	"github.com/joseluis8906/go-standard-layout/pkg/eventbus"
	"github.com/joseluis8906/go-standard-layout/pkg/events"
)

type (
	Builder struct {
		instance Post
		err      error
	}
)

func NewBuilder() *Builder {
	return &Builder{}
}

func (b Builder) Build(isNew bool) (Post, error) {
	if isNew {
		e := events.NewPostCreated()
		e.Attributes.Title = b.instance.title
		e.Attributes.Body = b.instance.body

		b.instance.AddEventID(e.ID)
		eventbus.Add(e.ID, e)
	}

	return b.instance, nil
}

func (b *Builder) WithID(id PostID) *Builder {
	b.instance.id = id

	return b
}

func (b *Builder) WithTitle(title string) *Builder {
	b.instance.title = title

	return b
}

func (b *Builder) WithBody(body string) *Builder {
	b.instance.body = body

	return b
}
