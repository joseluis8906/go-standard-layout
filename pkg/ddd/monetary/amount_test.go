package monetary_test

import (
	"testing"

	"github.com/joseluis8906/go-standard-layout/pkg/ddd/monetary"
)

func TestAmount_Add(t *testing.T) {
	a1, err := monetary.NewAmount(1.75, monetary.USD)
	if err != nil {
		t.Error(err)
	}

	inc, err := monetary.NewAmount(1.25, monetary.USD)
	if err != nil {
		t.Error(err)
	}

	want, err := monetary.NewAmount(3.00, monetary.USD)
	if err != nil {
		t.Error(err)
	}

	got, err := a1.Add(inc)
	if err != nil {
		t.Error(err)
	}

	if got != want {
		t.Errorf("Amount are different got %v; want %v", got, want)
	}
}

func TestAmount_Diff(t *testing.T) {
	a1, err := monetary.NewAmount(5.33, monetary.USD)
	if err != nil {
		t.Error(err)
	}

	dec, err := monetary.NewAmount(1.48, monetary.USD)
	if err != nil {
		t.Error(err)
	}

	want, err := monetary.NewAmount(3.85, monetary.USD)
	if err != nil {
		t.Error(err)
	}

	got, err := a1.Diff(dec)
	if err != nil {
		t.Error(err)
	}

	if got != want {
		t.Errorf("Amount are different got %v; want %v", got, want)
	}
}

func TestAmount_Diff_error(t *testing.T) {
	a1, err := monetary.NewAmount(1.48, monetary.USD)
	if err != nil {
		t.Error(err)
	}

	dec, err := monetary.NewAmount(1.49, monetary.USD)
	if err != nil {
		t.Error(err)
	}

	_, err = a1.Diff(dec)

	if err != monetary.ErrInvalidAmount {
		t.Errorf("Amount error is not the expected got %v; want %v", err, monetary.ErrInvalidAmount)
	}
}
