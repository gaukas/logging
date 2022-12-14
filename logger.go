package logging

import "os"

const (
	LOG_DEBUG uint8 = iota
	LOG_INFO
	LOG_WARN
	LOG_ERROR
	LOG_FATAL
	LOG_NOLOG
)

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

// CompatibleLogger is used for internal logger chaining/combining.
type CompatibleLogger interface {
	Logger
	_Fatalf(format string, args ...interface{}) // should not call os.Exit but do the rest of the work
}

// MultiLogger combines multiple loggers into one.
// Implements CompatibleLogger interface -- could be further chained.
type MultiLogger struct {
	loggers []CompatibleLogger
}

func NewMultiLogger(loggers ...CompatibleLogger) *MultiLogger {
	return &MultiLogger{loggers}
}

func (ml *MultiLogger) Debugf(format string, args ...interface{}) {
	for _, l := range ml.loggers {
		l.Debugf(format, args...)
	}
}

func (ml *MultiLogger) Infof(format string, args ...interface{}) {
	for _, l := range ml.loggers {
		l.Infof(format, args...)
	}
}

func (ml *MultiLogger) Warnf(format string, args ...interface{}) {
	for _, l := range ml.loggers {
		l.Warnf(format, args...)
	}
}

func (ml *MultiLogger) Errorf(format string, args ...interface{}) {
	for _, l := range ml.loggers {
		l.Errorf(format, args...)
	}
}

func (ml *MultiLogger) Fatalf(format string, args ...interface{}) {
	ml._Fatalf(format, args...)
	os.Exit(1)
}

func (ml *MultiLogger) _Fatalf(format string, args ...interface{}) {
	for _, l := range ml.loggers {
		l._Fatalf(format, args...)
	}
}
