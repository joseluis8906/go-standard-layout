package metrics

import "context"

type (
	Histogram struct {
		metric Metric
	}
)

func NewHistogram(opts ...option) *Histogram {
	c := &Histogram{}
	for _, opt := range opts {
		opt(&c.metric)
	}

	return c
}

func (h Histogram) Observe(ctx context.Context, val float64) error {
	h.metric.mType = HistogramType
	h.metric.value = val

	return collector.Collect(ctx, h.metric)
}

func (h Histogram) WithTag(key, val string) Histogram {
	_, ok := h.metric.tags[key]
	if ok {
		h.metric.tags[key] = val
	}

	return h
}
