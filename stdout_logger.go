package logging

import (
	"log"
	"os"
)

var (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
)

// StdoutLogger defines a logger that writes to stdout or stderr. It
// implements the Logger interface.
type StdoutLogger struct {
	logger *log.Logger // embedded

	_Debugf func(format string, v ...interface{})
	_Infof  func(format string, v ...interface{})
	_Warnf  func(format string, v ...interface{})
	_Errorf func(format string, v ...interface{})
	_Fatalf func(format string, v ...interface{})
}

// NewStderrLogger returns a new StdoutLogger that writes to stderr with
// specified prefix and flag. See log package from standard library for
// information about prefix and flag.
func NewStderrLogger(prefix string, flag int) *StdoutLogger {
	return &StdoutLogger{logger: log.New(os.Stderr, prefix, flag)}
}

// NewStdoutLogger returns a new StdoutLogger that writes to stdout with
// specified prefix and flag. See log package from standard library for
// information about prefix and flag.
func NewStdoutLogger(prefix string, flag int) *StdoutLogger {
	return &StdoutLogger{logger: log.New(os.Stdout, prefix, flag)}
}

// DefaultStderrLogger returns a new StdoutLogger that writes to stderr with
// no prefix and log.LstdFlags flag. It sets the logging level as specified.
func DefaultStderrLogger(loggingLevel uint8) *StdoutLogger {
	logger := NewStderrLogger("", log.LstdFlags)

	logger.SetLoggingLevel(loggingLevel)

	return logger
}

// SetLoggingLevel sets the logging level for the logger. The logging level
// is expected to be one of the following:
//
//	LOG_DEBUG
//	LOG_INFO
//	LOG_WARN
//	LOG_ERROR
//	LOG_FATAL
//	LOG_NOLOG
func (sl *StdoutLogger) SetLoggingLevel(level uint8) {
	if level <= LOG_DEBUG {
		sl._Debugf = sl._debugf
	} else {
		sl._Debugf = func(format string, v ...interface{}) {}
	}

	if level <= LOG_INFO {
		sl._Infof = sl._infof
	} else {
		sl._Infof = func(format string, v ...interface{}) {}
	}

	if level <= LOG_WARN {
		sl._Warnf = sl._warnf
	} else {
		sl._Warnf = func(format string, v ...interface{}) {}
	}

	if level <= LOG_ERROR {
		sl._Errorf = sl._errorf
	} else {
		sl._Errorf = func(format string, v ...interface{}) {}
	}

	if level <= LOG_FATAL {
		sl._Fatalf = sl._fatalf
	} else {
		sl._Fatalf = func(format string, v ...interface{}) {}
	}
}

// Debugf implements the Logger interface. It prints the message type in blue
// when logging level is set to LOG_DEBUG.
func (sl *StdoutLogger) Debugf(format string, v ...interface{}) {
	sl._Debugf(format, v...)
}

// Infof implements the Logger interface. It prints the message type in green
// when logging level is set to LOG_INFO or lower.
func (sl *StdoutLogger) Infof(format string, v ...interface{}) {
	sl._Infof(format, v...)
}

// Warnf implements the Logger interface. It prints the message type in yellow
// when logging level is set to LOG_WARN or lower.
func (sl *StdoutLogger) Warnf(format string, v ...interface{}) {
	sl._Warnf(format, v...)
}

// Errorf implements the Logger interface. It prints the message type in red
// when logging level is set to LOG_ERROR or lower.
func (sl *StdoutLogger) Errorf(format string, v ...interface{}) {
	sl._Errorf(format, v...)
}

// Fatalf implements the Logger interface. It prints the full message in red
// when logging level is set to LOG_FATAL or lower. It calls os.Exit(1) regardless
// of the logging level.
func (sl *StdoutLogger) Fatalf(format string, v ...interface{}) {
	sl._Fatalf(format, v...)
	os.Exit(1) // skipcq: RVV-A0003
}

// DeferredFatalf implements the CompatibleLogger interface. It does everything
// that Fatalf does, except it does not call os.Exit(1).
func (sl *StdoutLogger) DeferredFatalf(format string, v ...interface{}) {
	sl._Fatalf(format, v...)
}

// _debugf prints the message type in blue.
func (sl *StdoutLogger) _debugf(format string, v ...interface{}) {
	sl.logger.Printf("["+blue+"DEBUG"+reset+"] "+format, v...)
}

// _infof prints the message type in green.
func (sl *StdoutLogger) _infof(format string, v ...interface{}) {
	sl.logger.Printf("["+green+"INFO"+reset+"] "+format, v...)
}

// _warnf prints the message type in yellow.
func (sl *StdoutLogger) _warnf(format string, v ...interface{}) {
	sl.logger.Printf("["+yellow+"WARN"+reset+"] "+format, v...)
}

// _errorf prints the message type in red.
func (sl *StdoutLogger) _errorf(format string, v ...interface{}) {
	sl.logger.Printf("["+red+"ERROR"+reset+"] "+format, v...)
}

// _fatalf prints the full message in red.
func (sl *StdoutLogger) _fatalf(format string, v ...interface{}) {
	sl.logger.Printf(red+"[FATAL] "+format+reset, v...)
}
