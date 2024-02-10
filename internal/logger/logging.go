package logger

import (
	"time"

	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})

	return log
}

func logField(handler string, problem string) logrus.Fields {
	return logrus.Fields{
		"where":   handler,
		"problem": problem,
	}
}

func logError(handler string, problem string, err error) {
	logrus.WithFields(logField(handler, problem)).Error(err)
}
