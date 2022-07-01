package metrics

import (
	"context"
)

const (
	CounterType   = Type(1)
	GaugeType     = Type(2)
	HistogramType = Type(3)
	SummaryType   = Type(4)
)

type (
	Type int8

	Metric struct {
		mType     Type
		name      string
		namespace string
		subsystem string
		help      string
		tags      map[string]string
		value     float64
	}

	option func(*Metric)
)

func (m Metric) Type() Type {
	return m.mType
}

func (m Metric) Name() string {
	return m.name
}

func (m Metric) Namespace() string {
	return m.namespace
}

func (m Metric) Subsystem() string {
	return m.subsystem
}

func (m Metric) Help() string {
	return m.help
}

func (m Metric) Tags() map[string]string {
	return m.tags
}

func (m Metric) Value() float64 {
	return m.value
}

func Name(name string) option {
	return func(m *Metric) {
		m.name = name
	}
}

func Namespace(namespace string) option {
	return func(m *Metric) {
		m.namespace = namespace
	}
}

func Subsystem(subsystem string) option {
	return func(m *Metric) {
		m.subsystem = subsystem
	}
}

func Help(help string) option {
	return func(m *Metric) {
		m.help = help
	}
}

func Tags(keys ...string) option {
	return func(m *Metric) {
		if m.tags == nil {
			m.tags = map[string]string{}
		}

		for _, key := range keys {
			m.tags[key] = ""
		}
	}
}

type (
	Collector interface {
		Collect(context.Context, Metric) error
	}
)

var (
	collector Collector
)

func SetCollector(aCollector Collector) {
	collector = aCollector
}
