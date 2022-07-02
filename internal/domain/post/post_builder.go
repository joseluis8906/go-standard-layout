package post

type (
	Builder struct {
		instance Post
		err      error
	}
)

func NewBuilder() *Builder {
	return &Builder{}
}

func (b Builder) Build() (Post, error) {
	return b.instance, nil
}

func (b *Builder) WithTitle(title string) *Builder {
	b.instance.title = title

	return b
}

func (b *Builder) WithBody(body string) *Builder {
	b.instance.body = body

	return b
}
