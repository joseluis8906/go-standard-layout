package metrics

import "context"

type (
	Summary struct {
		metric Metric
	}
)

func NewSummary(opts ...option) *Summary {
	c := &Summary{}
	for _, opt := range opts {
		opt(&c.metric)
	}

	return c
}

func (s Summary) Observe(ctx context.Context, val float64) error {
	s.metric.mType = SummaryType
	s.metric.value = val

	return collector.Collect(ctx, s.metric)
}

func (s Summary) WithTag(key, val string) Summary {
	_, ok := s.metric.tags[key]
	if ok {
		s.metric.tags[key] = val
	}

	return s
}
