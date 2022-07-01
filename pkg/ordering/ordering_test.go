package ordering_test

import (
	"testing"

	"github.com/joseluis8906/go-standard-layout/pkg/ordering"
)

func TestOrdering(t *testing.T) {
	if ordering.Less > ordering.Equal {
		t.Error("Less is greater than Equal")
	}

	if ordering.Equal > ordering.Greater {
		t.Error("Equal is greater than Greater")
	}
}
