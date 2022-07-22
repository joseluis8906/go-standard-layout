package log

import (
	"fmt"
	"runtime"
	"time"
)

type (
	// Logger ...
	Logger interface {
		Panic(...interface{})
		Panicf(string, ...interface{})
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
		WithField(string, interface{}) Logger
		Level() Level
		Message() string
		Data() map[string]interface{}
		Time() time.Time
		AddHook(Hook)
	}

	// Hook ...
	Hook interface {
		Levels() []Level
		Fire(Logger) error
	}

	// Level ...
	Level int
)

// String ...
func (level Level) String() string {
	if b, err := level.MarshalText(); err == nil {
		return string(b)
	}

	return "unknown"
}

// MarshalText ...
func (level Level) MarshalText() ([]byte, error) {
	switch level {
	case TraceLevel:
		return []byte("trace"), nil
	case DebugLevel:
		return []byte("debug"), nil
	case InfoLevel:
		return []byte("info"), nil
	case WarnLevel:
		return []byte("warning"), nil
	case ErrorLevel:
		return []byte("error"), nil
	case FatalLevel:
		return []byte("fatal"), nil
	case PanicLevel:
		return []byte("panic"), nil
	}

	return nil, fmt.Errorf("not a valid log level %d", level)
}

const (
	// PanicLevel ...
	PanicLevel = Level(0)
	// FatalLevel ...
	FatalLevel = Level(1)
	// ErrorLevel ...
	ErrorLevel = Level(2)
	// WarnLevel ...
	WarnLevel = Level(3)
	// InfoLevel ...
	InfoLevel = Level(4)
	// DebugLevel ...
	DebugLevel = Level(5)
	// TraceLevel ...
	TraceLevel = Level(6)
)

var (
	logger Logger
)

// SetLogger ...
func SetLogger(aLogger Logger) {
	logger = aLogger
}

func init() {
	SetLogger(NoopLogger())
}

// Panic ...
func Panic(args ...interface{}) {
	caller().Panic(args...)
}

// Panicf ...
func Panicf(format string, args ...interface{}) {
	caller().Panicf(format, args...)
}

// Fatal ...
func Fatal(args ...interface{}) {
	caller().Fatal(args...)
}

// Fatalf ...
func Fatalf(format string, args ...interface{}) {
	caller().Fatalf(format, args...)
}

// Error ...
func Error(args ...interface{}) {
	caller().Error(args...)
}

// Errorf ...
func Errorf(format string, args ...interface{}) {
	caller().Errorf(format, args...)
}

// Warn ...
func Warn(args ...interface{}) {
	caller().Warn(args)
}

// Warnf ...
func Warnf(format string, args ...interface{}) {
	caller().Warnf(format, args...)
}

// Info ...
func Info(args ...interface{}) {
	caller().Info(args...)
}

// Infof ...
func Infof(format string, args ...interface{}) {
	caller().Infof(format, args...)
}

// Debug ...
func Debug(args ...interface{}) {
	caller().Debug(args...)
}

// Debugf ...
func Debugf(format string, args ...interface{}) {
	caller().Debugf(format, args...)
}

// Trace ...
func Trace(args ...interface{}) {
	caller().Trace(args...)
}

// Tracef ...
func Tracef(format string, args ...interface{}) {
	caller().Tracef(format, args...)
}

func caller() Logger {
	pc, file, line, ok := runtime.Caller(2)
	if ok {
		funcName := runtime.FuncForPC(pc).Name()
		return logger.WithField("file", fmt.Sprintf("%s:%d", file, line)).WithField("func", funcName)
	}

	return logger
}
