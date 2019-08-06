package log

import (
	"github.com/sirupsen/logrus"
)

// Fields - map with additional log entry fields
type Fields logrus.Fields

// Logger - simplified logger abstraction
type Logger interface {
	WithField(string, interface{}) Logger
	WithFields(Fields) Logger

	// Use this level only for local development needs
	Debug(...interface{})

	// Use this level for information that helps to find the reason of failure,
	//  e.g. to trace REST API calls
	Info(...interface{})

	// Use this level for failures
	Error(error, ...interface{})
}

// MainLogger - Logger which can also report fatal errors
type MainLogger interface {
	Logger
	FatalError(error, ...interface{})
}
