package customer

import "github.com/google/uuid"

type (
	CustomerID struct {
		val uuid.UUID
	}
)

func NewCustomerID() (CustomerID, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return NoopCustomerID(), err
	}

	return CustomerID{val: uid}, nil
}

func ParseCurtomerID(raw string) (CustomerID, error) {
	val, err := uuid.Parse(raw)
	if err != nil {
		return NoopCustomerID(), err
	}

	return CustomerID{val}, nil
}

func NoopCustomerID() CustomerID {
	return CustomerID{val: uuid.UUID{}}
}

func (u CustomerID) String() string {
	return u.val.String()
}
