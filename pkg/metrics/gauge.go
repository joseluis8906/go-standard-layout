package metrics

import "context"

type (
	Gauge struct {
		metric Metric
	}
)

func NewGauge(opts ...option) *Gauge {
	c := &Gauge{}
	for _, opt := range opts {
		opt(&c.metric)
	}

	return c
}

func (g Gauge) Set(ctx context.Context, val float64) error {
	g.metric.mType = GaugeType
	g.metric.value = val

	return collector.Collect(ctx, g.metric)
}

func (g Gauge) WithTag(key, val string) Gauge {
	_, ok := g.metric.tags[key]
	if ok {
		g.metric.tags[key] = val
	}

	return g
}
