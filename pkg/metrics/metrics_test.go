package metrics_test

import (
	"context"

	"github.com/joseluis8906/go-standard-layout/pkg/metrics"
)

type (
	CollectorMock func(context.Context, metrics.Metric) error
)

func (cm CollectorMock) Collect(ctx context.Context, m metrics.Metric) error {
	if cm == nil {
		panic("CollectorMock is not implemented")
	}

	return cm(ctx, m)
}
