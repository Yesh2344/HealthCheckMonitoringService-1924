package logger

import (
	"github.com/sirupsen/logrus"
)

// Logger represents the application logger
type Logger interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

// NewLogger returns a new application logger
func NewLogger() Logger {
	return &logger{logger: logrus.New()}
}

type logger struct {
	logger *logrus.Logger
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}