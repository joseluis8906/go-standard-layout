package log

import (
	"fmt"
	"runtime"

	"github.com/sirupsen/logrus"
)

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
		WithField(string, interface{}) Logger
		AddHook(logrus.Hook)
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
	caller().Fatal(a...)
}

func Fatalf(format string, a ...interface{}) {
	caller().Fatalf(format, a...)
}

func Error(a ...interface{}) {
	caller().Error(a...)
}

func Errorf(format string, a ...interface{}) {
	caller().Errorf(format, a...)
}

func Warn(a ...interface{}) {
	caller().Warn(a)
}

func Warnf(format string, a ...interface{}) {
	caller().Warnf(format, a...)
}

func Info(a ...interface{}) {
	caller().Info(a...)
}

func Infof(format string, a ...interface{}) {
	caller().Infof(format, a...)
}

func Debug(a ...interface{}) {
	caller().Debug(a...)
}

func Debugf(format string, a ...interface{}) {
	caller().Debugf(format, a...)
}

func Trace(a ...interface{}) {
	caller().Trace(a...)
}

func Tracef(format string, a ...interface{}) {
	caller().Tracef(format, a...)
}

func caller() Logger {
	pc, file, line, ok := runtime.Caller(2)
	if ok {
		funcName := runtime.FuncForPC(pc).Name()
		return logger.WithField("file", fmt.Sprintf("%s:%d", file, line)).WithField("func", funcName)
	}

	return logger
}
