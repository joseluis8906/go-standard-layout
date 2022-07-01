package retail

import (
	"errors"
)

var (
	ErrInvalidLen    = errors.New("retail.SKU invalid length")
	ErrInvalidFormat = errors.New("retail.SKU invalid format")
)

const (
	minLen = 4
)

type (
	SKU struct {
		val string
	}
)

func NoopSKU() SKU {
	return SKU{val: "<nil>"}
}

func NewSKU(sku string) (SKU, error) {
	if len(sku) < minLen {
		return NoopSKU(), ErrInvalidLen
	}

	return SKU{val: sku}, nil
}

func (s SKU) String() string {
	return s.val
}
