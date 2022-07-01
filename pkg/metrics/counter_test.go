package metrics_test

import (
	"context"
	"testing"

	"github.com/joseluis8906/go-standard-layout/pkg/metrics"
)

var (
	ctx = context.Background()
)

func TestCounter_Inc(t *testing.T) {
	metrics.SetCollector(CollectorMock(func(context.Context, metrics.Metric) error {
		return nil
	}))

	counter := metrics.NewCounter(
		metrics.Name("test"),
		metrics.Namespace("test"),
		metrics.Subsystem("test"),
		metrics.Help("test metric"),
		metrics.Tags("tag1", "tag2"),
	)

	err := counter.
		WithTag("tag1", "val 1").
		WithTag("tag2", "val 2").
		Inc(ctx)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
