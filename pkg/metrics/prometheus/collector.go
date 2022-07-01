package prometheus

import (
	"context"
	"fmt"
	"sync"

	"github.com/joseluis8906/go-standard-layout/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type (
	Collector struct {
		store struct {
			keys       map[string]struct{}
			counters   map[string]*prometheus.CounterVec
			gauges     map[string]*prometheus.GaugeVec
			histograms map[string]*prometheus.HistogramVec
			summaries  map[string]*prometheus.SummaryVec
		}
		mux sync.Mutex
	}
)

func NewCollector() *Collector {
	c := Collector{}
	c.store.keys = map[string]struct{}{}
	c.store.counters = map[string]*prometheus.CounterVec{}
	c.store.gauges = map[string]*prometheus.GaugeVec{}
	c.store.histograms = map[string]*prometheus.HistogramVec{}
	c.store.summaries = map[string]*prometheus.SummaryVec{}

	return &c
}

func (c *Collector) Collect(ctx context.Context, metric metrics.Metric) error {
	key := c.getKey(metric)
	_, ok := c.store.keys[key]
	if !ok {
		c.add(key, metric)
	}

	c.publish(metric)

	return nil
}

func (c *Collector) add(key string, m metrics.Metric) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.store.keys[key] = struct{}{}
	switch m.Type() {
	case metrics.CounterType:
		counter := promauto.NewCounterVec(prometheus.CounterOpts{
			Subsystem: m.Subsystem(),
			Namespace: m.Namespace(),
			Name:      m.Name(),
			Help:      m.Help(),
		}, c.getLabelNames(m))

		c.store.counters[key] = counter

	case metrics.GaugeType:
		gauge := promauto.NewGaugeVec(prometheus.GaugeOpts{
			Subsystem: m.Subsystem(),
			Namespace: m.Namespace(),
			Name:      m.Name(),
			Help:      m.Help(),
		}, c.getLabelNames(m))

		c.store.gauges[key] = gauge

	case metrics.HistogramType:
		histogram := promauto.NewHistogramVec(prometheus.HistogramOpts{
			Subsystem: m.Subsystem(),
			Namespace: m.Namespace(),
			Name:      m.Name(),
			Help:      m.Help(),
		}, c.getLabelNames(m))

		c.store.histograms[key] = histogram

	case metrics.SummaryType:
		summary := promauto.NewSummaryVec(prometheus.SummaryOpts{
			Subsystem: m.Subsystem(),
			Namespace: m.Namespace(),
			Name:      m.Name(),
			Help:      m.Help(),
		}, c.getLabelNames(m))

		c.store.summaries[key] = summary
	}
}

func (c *Collector) publish(m metrics.Metric) {
	switch m.Type() {
	case metrics.CounterType:
		counter := c.store.counters[c.getKey(m)]
		counter.With(m.Tags()).Inc()

	case metrics.GaugeType:
		gauge := c.store.gauges[c.getKey(m)]
		gauge.With(m.Tags()).Set(m.Value())

	case metrics.HistogramType:
		histogram := c.store.histograms[c.getKey(m)]
		histogram.With(m.Tags()).Observe(m.Value())

	case metrics.SummaryType:
		summary := c.store.summaries[c.getKey(m)]
		summary.With(m.Tags()).Observe(m.Value())
	}
}

func (c *Collector) getKey(m metrics.Metric) string {
	return fmt.Sprintf("%s.%s.%s", m.Subsystem(), m.Namespace(), m.Name())
}

func (c *Collector) getLabelNames(m metrics.Metric) []string {
	labels := []string{}
	for k := range m.Tags() {
		labels = append(labels, k)
	}

	return labels
}
