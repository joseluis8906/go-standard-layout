package post

import "github.com/google/uuid"

const (
	nilPostIDVal = "00000000-0000-0000-0000-000000000000"
)

type (
	PostID struct {
		val uuid.UUID
	}
)

func NewPostID() (PostID, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return NoopPostID(), err
	}

	return PostID{val: id}, nil
}

func ParsePostID(val string) (PostID, error) {
	id, err := uuid.Parse(val)
	if err != nil {
		return NoopPostID(), err
	}

	return PostID{val: id}, nil
}

func NoopPostID() PostID {
	return PostID{val: uuid.UUID{}}
}

func (pi PostID) IsZero() bool {
	return pi.val.String() == nilPostIDVal
}

func (pi PostID) String() string {
	return pi.val.String()
}
