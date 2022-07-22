package fluentd

import (
	"fmt"
	"time"

	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/joseluis8906/go-standard-layout/pkg/log"
)

type (
	// Hook ...
	Hook struct {
		client *fluent.Fluent
		levels []log.Level
	}
)

// NewHook ...
func NewHook(opts ...OptionFunc) *Hook {
	conf := &config{}
	for _, opt := range opts {
		opt(conf)
	}

	client, err := fluent.New(fluent.Config{
		FluentHost: conf.fluentdHost,
		FluentPort: conf.fluentdPort,
		Async:      conf.fluentdAsync,
	})

	if err != nil {
		log.Fatalf("error trying to connect to fluentd: %v", err)
	}

	return &Hook{
		client: client,
		levels: []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
			log.WarnLevel,
			log.InfoLevel,
			log.DebugLevel,
			log.TraceLevel,
		},
	}
}

// Close ...
func (h *Hook) Close() error {
	return h.client.Close()
}

// Fire ...
func (h *Hook) Fire(e log.Logger) error {
	data := map[string]string{
		"level":   e.Level().String(),
		"message": e.Message(),
		"time":    e.Time().Format(time.RFC3339Nano),
	}

	for k, v := range e.Data() {
		data[k] = fmt.Sprintf("%s", v)
	}

	tag, ok := data["app"]
	if !ok {
		tag = "fluentd"
	}

	return h.client.PostWithTime(tag, e.Time(), data)
}

// Levels ...
func (h *Hook) Levels() []log.Level {
	return h.levels
}
