package log

type (
	Logger interface {
		Fatal(...interface{})
		Fatalf(string, ...interface{})
		Error(...interface{})
		Errorf(string, ...interface{})
		Warn(...interface{})
		Warnf(string, ...interface{})
		Info(...interface{})
		Infof(string, ...interface{})
		Debug(...interface{})
		Debugf(string, ...interface{})
		Trace(...interface{})
		Tracef(string, ...interface{})
	}
)

var (
	logger Logger
)

func SetLogger(aLogger Logger) {
	logger = aLogger
}

func init() {
	SetLogger(NoopLogger())
}

func Fatal(a ...interface{}) {
	logger.Fatal(a...)
}

func Fatalf(format string, a ...interface{}) {
	logger.Fatalf(format, a...)
}

func Error(a ...interface{}) {
	logger.Error(a...)
}

func Errorf(format string, a ...interface{}) {
	logger.Errorf(format, a...)
}

func Warn(a ...interface{}) {
	logger.Warn(a)
}

func Warnf(format string, a ...interface{}) {
	logger.Warnf(format, a...)
}

func Info(a ...interface{}) {
	logger.Info(a...)
}

func Infof(format string, a ...interface{}) {
	logger.Infof(format, a...)
}

func Debug(a ...interface{}) {
	logger.Debug(a...)
}

func Debugf(format string, a ...interface{}) {
	logger.Debugf(format, a...)
}

func Trace(a ...interface{}) {
	logger.Trace(a...)
}

func Tracef(format string, a ...interface{}) {
	logger.Tracef(format, a...)
}
