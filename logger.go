package logging

import "os"

const (
	LOG_DEBUG uint8 = iota // lowest logging level, all debug details will be logged
	LOG_INFO
	LOG_WARN
	LOG_ERROR
	LOG_FATAL
	LOG_NOLOG // highest logging level, no log will be written, but Fatalf will still exit the program
)

// Logger defines an interface for logging output.
// Debugf, Infof, Warnf, Errorf, and Fatalf are the logging functions.
// The Fatalf function will call os.Exit(1) after logging, even when logging
// level is higher than LOG_FATAL.
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
	DeferredFatalf(format string, args ...interface{}) // should not call os.Exit but do the rest of the work.
}

// MultiLogger combines multiple loggers into one.
// Implements CompatibleLogger interface -- could be further chained.
type MultiLogger struct {
	loggers []CompatibleLogger
}

func NewMultiLogger(loggers ...CompatibleLogger) *MultiLogger {
	return &MultiLogger{loggers}
}

// Debugf implements Logger interface.
func (ml *MultiLogger) Debugf(format string, args ...interface{}) {
	for _, l := range ml.loggers {
		l.Debugf(format, args...)
	}
}

// Infof implements Logger interface.
func (ml *MultiLogger) Infof(format string, args ...interface{}) {
	for _, l := range ml.loggers {
		l.Infof(format, args...)
	}
}

// Warnf implements Logger interface.
func (ml *MultiLogger) Warnf(format string, args ...interface{}) {
	for _, l := range ml.loggers {
		l.Warnf(format, args...)
	}
}

// Errorf implements Logger interface.
func (ml *MultiLogger) Errorf(format string, args ...interface{}) {
	for _, l := range ml.loggers {
		l.Errorf(format, args...)
	}
}

// Fatalf implements Logger interface.
func (ml *MultiLogger) Fatalf(format string, args ...interface{}) {
	ml.DeferredFatalf(format, args...)
	os.Exit(1)
}

func (ml *MultiLogger) DeferredFatalf(format string, args ...interface{}) {
	for _, l := range ml.loggers {
		l.DeferredFatalf(format, args...)
	}
}
