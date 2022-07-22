package logrus

import (
	"os"
	"path/filepath"
	"time"

	"github.com/joseluis8906/go-standard-layout/pkg/log"
	"github.com/sirupsen/logrus"
)

// New ...
func New(opts ...OptionFunc) *Entry {
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
func (e *Entry) WithField(key string, value interface{}) log.Logger {
	e.entry = e.entry.WithField(key, value)
	return e
}

// Panic ...
func (e *Entry) Panic(args ...interface{}) {
	e.entry.Panic(args...)
}

// Panicf ...
func (e *Entry) Panicf(format string, args ...interface{}) {
	e.entry.Panicf(format, args...)
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
func (e *Entry) AddHook(aHook log.Hook) {
	realHook := hook{
		LevelFunc: aHook.Levels,
		FireFunc:  aHook.Fire,
	}

	e.entry.Logger.AddHook(realHook)
}

// Level ...
func (e *Entry) Level() log.Level {
	return log.Level(e.entry.Level)
}

// Time ...
func (e *Entry) Time() time.Time {
	return e.entry.Time
}

// Message ...
func (e *Entry) Message() string {
	return e.entry.Message
}

// Data ...
func (e *Entry) Data() map[string]interface{} {
	return e.entry.Data
}

type hook struct {
	LevelFunc func() []log.Level
	FireFunc  func(log.Logger) error
}

func (h hook) Levels() []logrus.Level {
	levels := h.LevelFunc()
	logrusLevels := make([]logrus.Level, len(levels))

	for i, level := range levels {
		logrusLevels[i] = logrus.Level(level)
	}

	return logrusLevels
}

func (h hook) Fire(e *logrus.Entry) error {
	realEntry := &Entry{entry: e}
	return h.FireFunc(realEntry)
}
