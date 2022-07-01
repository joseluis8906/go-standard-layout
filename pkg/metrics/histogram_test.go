package metrics_test

import (
	"context"
	"testing"

	"github.com/joseluis8906/go-standard-layout/pkg/metrics"
)

func TestHistogram_Set(t *testing.T) {
	metrics.SetCollector(CollectorMock(func(context.Context, metrics.Metric) error {
		return nil
	}))

	histogram := metrics.NewHistogram(
		metrics.Name("test"),
		metrics.Namespace("test"),
		metrics.Subsystem("test"),
		metrics.Help("test metric"),
		metrics.Tags("tag1", "tag2"),
	)

	err := histogram.
		WithTag("tag1", "val 1").
		WithTag("tag2", "val 2").
		Observe(ctx, 4)

	if err != nil {
		t.Errorf("Unexpeted error: %v", err)
	}
}
