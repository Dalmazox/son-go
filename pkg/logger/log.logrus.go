package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.JSONFormatter{})
}

func Info(msg string, fields logrus.Fields) {
	log.WithFields(fields).Info(msg)
}

func Fatal(msg string, fields logrus.Fields) {
	log.WithFields(fields).Fatal(msg)
}

func Error(msg string, fields logrus.Fields) {
	log.WithFields(fields).Error(msg)
}
