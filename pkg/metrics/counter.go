package metrics

import (
	"context"
)

type (
	Counter struct {
		metric Metric
	}
)

func NewCounter(opts ...option) *Counter {
	c := &Counter{}
	for _, opt := range opts {
		opt(&c.metric)
	}

	return c
}

func (c Counter) Inc(ctx context.Context) error {
	c.metric.mType = CounterType

	return collector.Collect(ctx, c.metric)
}

func (c Counter) WithTag(key, val string) Counter {
	_, ok := c.metric.tags[key]
	if ok {
		c.metric.tags[key] = val
	}

	return c
}
