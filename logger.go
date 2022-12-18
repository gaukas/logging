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
	// Debugf logs a debug message which contains debugging details
	// that are meaningful to developers for debugging purposes ONLY.
	// It SHOULD NOT contain sensitive information and SHOULD be trivial
	// in a production environment.
	Debugf(format string, args ...interface{})

	// Infof logs an info message which contains information that
	// may be useful to users for detailed behavior logging.
	// It MUST not contain any sensitive information and SHOULD
	// indicate normal/expeted behaviors of the program.
	Infof(format string, args ...interface{})

	// Warnf logs a warning message which SHOULD be paid attention to.
	// It MUST not contain any sensitive information and MAY indicate abnormal,
	// unusual, or unexpected behaviors of the program.
	Warnf(format string, args ...interface{})

	// Errorf logs an error message which MUST be paid attention to.
	// It MUST not contain any sensitive information and MUST indicate
	// abnormal, unusual, or unexpected behaviors of the program which
	// does not cause the program to exit but may cause the program to
	// behave incorrectly under certain circumstances.
	Errorf(format string, args ...interface{})

	// Fatalf logs a fatal message which MUST be paid attention to.
	// It MUST not contain any sensitive information and MUST indicate
	// abnormal, unusual, or unexpected behaviors of the program which
	// prevents the program from continuing to run and causes the
	// program to crash or exit.
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

// NewMultiLogger creates a MultiLogger from given loggers.
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
	os.Exit(1) // skipcq: RVV-A0003
}

// DeferredFatalf implements CompatibleLogger interface.
func (ml *MultiLogger) DeferredFatalf(format string, args ...interface{}) {
	for _, l := range ml.loggers {
		l.DeferredFatalf(format, args...)
	}
}
