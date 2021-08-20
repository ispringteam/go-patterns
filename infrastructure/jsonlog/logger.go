package jsonlog

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/ispringtech/go-patterns/infrastructure/jsonlog/hooks"
	"github.com/ispringtech/go-patterns/infrastructure/log"
)

const appNameKey = "app_name"

type logger struct {
	logrus.FieldLogger
}

var fieldMap = logrus.FieldMap{
	logrus.FieldKeyTime: "@timestamp",
	logrus.FieldKeyMsg:  "message",
}

//	NewLogger creates a new logger which does the following:
//	- Formats logs in JSON
//	- Uses logging level from config
//	- Adds fields:
//		- "app_name" - application name from given Config
//		- "@timestamp" - RFC3339Nano timestamp
//		- "level" - logging level
//		- "stack" - stack trace for errors
//		- "message" - log message
func NewLogger(config *Config) log.MainLogger {
	impl := logrus.New()
	impl.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		FieldMap:        fieldMap,
		PrettyPrint:     config.PrettyPrint,
	})
	impl.SetLevel(logrus.Level(config.Level))
	impl.AddHook(new(hooks.StackTraceHook))

	return &logger{
		FieldLogger: impl.WithField(appNameKey, config.AppName),
	}
}

func (l *logger) WithField(key string, value interface{}) log.Logger {
	return &logger{l.FieldLogger.WithField(key, value)}
}

func (l *logger) WithFields(fields log.Fields) log.Logger {
	return &logger{l.FieldLogger.WithFields(logrus.Fields(fields))}
}

func (l *logger) Error(err error, args ...interface{}) {
	l.FieldLogger.WithError(err).Error(args...)
}

func (l *logger) FatalError(err error, args ...interface{}) {
	l.FieldLogger.WithError(err).Fatal(args...)
}
