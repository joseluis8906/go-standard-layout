package log

import "github.com/sirupsen/logrus"

type (
	noopLogger struct{}
)

func NoopLogger() Logger {
	return &noopLogger{}
}

func (n *noopLogger) Fatal(_ ...interface{}) {}

func (n *noopLogger) Fatalf(_ string, _ ...interface{}) {}

func (n *noopLogger) Error(_ ...interface{}) {}

func (n *noopLogger) Errorf(_ string, _ ...interface{}) {}

func (n *noopLogger) Warn(_ ...interface{}) {}

func (n *noopLogger) Warnf(_ string, _ ...interface{}) {}

func (n *noopLogger) Info(_ ...interface{}) {}

func (n *noopLogger) Infof(_ string, _ ...interface{}) {}

func (n *noopLogger) Debug(_ ...interface{}) {}

func (n *noopLogger) Debugf(_ string, _ ...interface{}) {}

func (n *noopLogger) Trace(_ ...interface{}) {}

func (n *noopLogger) Tracef(_ string, _ ...interface{}) {}

func (n *noopLogger) WithField(_ string, _ interface{}) Logger { return n }

func (n *noopLogger) AddHook(_ logrus.Hook) {}
