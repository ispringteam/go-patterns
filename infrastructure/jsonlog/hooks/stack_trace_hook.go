package hooks

import (
	"github.com/sirupsen/logrus"

	"github.com/ispringteam/go-patterns/infrastructure/errors"
)

type StackTraceHook struct{}

const keyStack = "stack"

func (h *StackTraceHook) Fire(entry *logrus.Entry) error {
	val, ok := entry.Data[logrus.ErrorKey]
	if !ok {
		return nil
	}

	err, ok := val.(error)
	if !ok {
		return nil
	}

	if err == nil {
		delete(entry.Data, logrus.ErrorKey)
		return nil
	}

	trace := errors.GetStackTrace(err)
	if trace != nil {
		entry.Data[keyStack] = trace
	}

	entry.Data[logrus.ErrorKey] = err.Error()

	return nil
}

func (h *StackTraceHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
