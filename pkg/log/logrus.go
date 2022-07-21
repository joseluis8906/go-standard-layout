package log

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func NewLogrus(opts ...OptionFunc) Logger {
	conf := &config{}
	for _, opt := range opts {
		opt(conf)
	}

	serviceName := filepath.Base(os.Args[0])
	env := conf.env
	if env == "" {
		env = "DEV"
	}

	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.Level(conf.level))
	logger.SetFormatter(func() logrus.Formatter {
		if conf.formatter == "json" {
			return &logrus.JSONFormatter{}
		}
		return &logrus.TextFormatter{}
	}())

	entry := logger.WithField("env", env).WithField("app", serviceName)

	return &Entry{entry}
}

// Entry ,,,
type Entry struct {
	entry *logrus.Entry
}

// WithField ...
func (e *Entry) WithField(key string, value interface{}) Logger {
	e.entry = e.entry.WithField(key, value)
	return e
}

// Fatal ...
func (e *Entry) Fatal(args ...interface{}) {
	e.entry.Fatal(args...)
}

// Fatalf ...
func (e *Entry) Fatalf(format string, args ...interface{}) {
	e.entry.Fatalf(format, args...)
}

// Error ..
func (e *Entry) Error(args ...interface{}) {
	e.entry.Error(args...)
}

// Errorf ...
func (e *Entry) Errorf(format string, args ...interface{}) {
	e.entry.Errorf(format, args...)
}

// Warn ...
func (e *Entry) Warn(args ...interface{}) {
	e.entry.Warn(args...)
}

// Warnf ...
func (e *Entry) Warnf(format string, args ...interface{}) {
	e.entry.Warnf(format, args...)
}

// Info ...
func (e *Entry) Info(args ...interface{}) {
	e.entry.Info(args...)
}

// Infof ...
func (e *Entry) Infof(format string, args ...interface{}) {
	e.entry.Infof(format, args...)
}

// Debug ...
func (e *Entry) Debug(args ...interface{}) {
	e.entry.Debug(args...)
}

// Debugf ...
func (e *Entry) Debugf(format string, args ...interface{}) {
	e.entry.Debugf(format, args...)
}

// Trace ...
func (e *Entry) Trace(args ...interface{}) {
	e.entry.Trace(args...)
}

// Tracef ...
func (e *Entry) Tracef(format string, args ...interface{}) {
	e.entry.Tracef(format, args...)
}

// AddHook ...
func (e *Entry) AddHook(hook logrus.Hook) {
	e.entry.Logger.AddHook(hook)
}
