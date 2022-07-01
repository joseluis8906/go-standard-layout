package inventory

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrInvalidRange = errors.New("inventory.Quantity invalid range")
	ErrInvalidType  = errors.New("inventory.Quantity string is not a quantity representation")
)

type (
	Quantity struct {
		val int
	}
)

func NoopQuanity() Quantity {
	return Quantity{val: -1}
}

func NewQuantity(quantity int) (Quantity, error) {
	if quantity < 0 {
		return NoopQuanity(), ErrInvalidRange
	}

	return Quantity{val: quantity}, nil
}

func (q Quantity) Inc(val int) Quantity {
	return Quantity{val: q.val + val}
}

func (q Quantity) Dec(val int) (Quantity, error) {
	if q.val-val < 0 {
		return NoopQuanity(), ErrInvalidRange
	}

	return Quantity{val: q.val - val}, nil
}

func (q Quantity) String() string {
	return fmt.Sprintf("Quantity:%d", q.val)
}

func (q Quantity) Int() int {
	return q.val
}

func Parse(data string) (Quantity, error) {
	if !strings.HasPrefix(data, "Quantity:") {
		return NoopQuanity(), ErrInvalidType
	}

	parts := strings.Split(data, ":")
	val, err := strconv.Atoi(parts[1])
	if err != nil {
		return NoopQuanity(), err
	}

	return Quantity{val}, nil
}
