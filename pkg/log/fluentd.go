package log

import (
	"fmt"
	"time"

	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/sirupsen/logrus"
)

type (
	FluentHook struct {
		client *fluent.Fluent
		levels []logrus.Level
	}
)

func NewFluentHook(opts ...OptionFunc) *FluentHook {
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
		Fatalf("error trying to connect to fluentd: %v", err)
	}

	return &FluentHook{
		client: client,
		levels: []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
			logrus.WarnLevel,
			logrus.InfoLevel,
			logrus.DebugLevel,
			logrus.TraceLevel,
		},
	}
}

func (h *FluentHook) Close() error {
	return h.client.Close()
}

func (h *FluentHook) Fire(e *logrus.Entry) error {
	data := map[string]string{
		"level":   e.Level.String(),
		"message": e.Message,
		"time":    e.Time.Format(time.RFC3339Nano),
	}

	for k, v := range e.Data {
		data[k] = fmt.Sprintf("%s", v)
	}

	if e.HasCaller() {
		data["file"] = fmt.Sprintf("%s:%d", e.Caller.File, e.Caller.Line)
		data["func"] = e.Caller.Function
	}

	tag, ok := data["app"]
	if !ok {
		tag = "fluentd"
	}

	return h.client.PostWithTime(tag, e.Time, data)
}

func (h *FluentHook) Levels() []logrus.Level {
	return h.levels
}
