package log

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Entry

// SetLogger ...
func SetLogger(aLogger *logrus.Entry) {
	logger = aLogger
}

// Logger ...
func Logger() *logrus.Entry {
	return logger
}

// Noop ...
func Noop() *logrus.Entry {
	logger := logrus.New()
	logger.SetOutput(ioutil.Discard)

	return logger.WithField("type", "Not Operation Logger")
}
