package log

type (
	noopLogger struct{}
)

func NoopLogger() Logger {
	return &noopLogger{}
}

func (n *noopLogger) Fatal(a ...interface{}) {}

func (n *noopLogger) Fatalf(format string, a ...interface{}) {}

func (n *noopLogger) Error(a ...interface{}) {}

func (n *noopLogger) Errorf(format string, a ...interface{}) {}

func (n *noopLogger) Warn(a ...interface{}) {}

func (n *noopLogger) Warnf(format string, a ...interface{}) {}

func (n *noopLogger) Info(a ...interface{}) {}

func (n *noopLogger) Infof(format string, a ...interface{}) {}

func (n *noopLogger) Debug(a ...interface{}) {}

func (n *noopLogger) Debugf(format string, a ...interface{}) {}

func (n *noopLogger) Trace(a ...interface{}) {}

func (n *noopLogger) Tracef(format string, a ...interface{}) {}
