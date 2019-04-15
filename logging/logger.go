package logging

import (
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func init() {
	// Log as JSON instead of the default ASCII formatter.
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Only log the warning severity or above.
	// Default log level
	logger.SetLevel(logrus.InfoLevel)
}

func WithField(key string, value interface{}) *logrus.Entry {
	return logger.WithField(key, value)
}

func Info(msg string) {
	logger.Info(msg)
}

func Infof(msg string) {
	logger.Infof(msg)
}

func Debug(msg string) {
	logger.Debug(msg)
}

func Error(trace string, err error) {
	logger.WithFields(logrus.Fields{
		"line": trace,
	}).Error(err)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Println(args ...interface{}) {
	logger.Println(args...)
}
