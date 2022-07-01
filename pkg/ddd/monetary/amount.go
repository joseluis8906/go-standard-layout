package monetary

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/joseluis8906/go-standard-layout/pkg/ordering"
)

var (
	ErrInvalidAmount            = errors.New("monetary.Amount value is less than 0")
	ErrCurrencyUnavailable      = errors.New("monetary.Amount currency is unavailable")
	ErrCurrencyAreNotComparable = errors.New("monetary.Amount currencies are not comparable")
)

const (
	USD = "USD"
)

var (
	currencies = []string{USD}
)

type (
	Amount struct {
		value    float64
		currency string
	}
)

func NoopAmount() Amount {
	return Amount{value: -1, currency: "<nil>"}
}

func NewAmount(amount float64, currency string) (Amount, error) {
	if amount < 0 {
		return NoopAmount(), ErrInvalidAmount
	}

	ok := false
	for _, curr := range currencies {
		if curr == currency {
			ok = true
			break
		}
	}

	if !ok {
		return NoopAmount(), ErrCurrencyUnavailable
	}

	return Amount{amount, currency}, nil
}

func Parse(raw string) (Amount, error) {
	parts := strings.Split(raw, " ")
	amount, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return NoopAmount(), err
	}

	currency := parts[1]

	return Amount{value: amount, currency: currency}, nil
}

func (c Amount) check(another Amount) error {
	if c.currency != another.currency {
		return ErrCurrencyAreNotComparable
	}

	return nil
}

func (c Amount) String() string {
	return fmt.Sprintf("%.2f %s", c.value, c.currency)
}

func (c Amount) Add(increment Amount) (Amount, error) {
	err := c.check(increment)
	if err != nil {
		return NoopAmount(), err
	}

	return Amount{value: c.value + increment.value, currency: c.currency}, nil
}

func (c Amount) Diff(another Amount) (Amount, error) {
	cmp, err := c.Cmp(another)
	if err != nil {
		return NoopAmount(), err
	}

	if cmp == ordering.Less {
		return NoopAmount(), ErrInvalidAmount
	}

	return Amount{value: c.value - another.value, currency: c.currency}, nil
}

func (c Amount) Cmp(another Amount) (ordering.Cmp, error) {
	err := c.check(another)
	if err != nil {
		return ordering.Equal, err
	}

	if c.value < another.value {
		return ordering.Less, nil
	}

	if c.value > another.value {
		return ordering.Greater, nil
	}

	return ordering.Equal, nil
}

func (c Amount) AsFloat() float64 {
	return c.value
}
