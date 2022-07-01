package log

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func NewLogrus(opts ...OptionFunc) *logrus.Entry {
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
	logger.SetReportCaller(conf.caller)
	logger.SetFormatter(func() logrus.Formatter {
		if conf.formatter == "json" {
			return &logrus.JSONFormatter{}
		}
		return &logrus.TextFormatter{}
	}())

	return logger.WithField("env", env).WithField("app", serviceName)
}
