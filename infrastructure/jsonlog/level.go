package jsonlog

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Level logrus.Level

const (
	ErrorLevel = Level(logrus.ErrorLevel)
	InfoLevel  = Level(logrus.InfoLevel)
	DebugLevel = Level(logrus.DebugLevel)
)

func ParseLevel(level string) (Level, error) {
	result, err := logrus.ParseLevel(level)
	switch result {
	case logrus.FatalLevel, logrus.PanicLevel, logrus.WarnLevel, logrus.TraceLevel:
		return 0, errors.Errorf("unsupported logging level: `%s`", level)
	default:
		return Level(result), err
	}
}
