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

type StdoutLogger struct {
	logger *log.Logger // embedded

	_Debugf func(format string, v ...interface{})
	_Infof  func(format string, v ...interface{})
	_Warnf  func(format string, v ...interface{})
	_Errorf func(format string, v ...interface{})
	_Fatalf func(format string, v ...interface{})
}

func NewStderrLogger(prefix string, flag int) *StdoutLogger {
	return &StdoutLogger{logger: log.New(os.Stderr, prefix, flag)}
}

func NewStdoutLogger(prefix string, flag int) *StdoutLogger {
	return &StdoutLogger{logger: log.New(os.Stdout, prefix, flag)}
}

func DefaultStderrLogger(loggingLevel uint8) *StdoutLogger {
	logger := NewStderrLogger("", log.LstdFlags)

	logger.SetLoggingLevel(loggingLevel)

	return logger
}

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

func (sl *StdoutLogger) Debugf(format string, v ...interface{}) {
	sl._Debugf(format, v...)
}

func (sl *StdoutLogger) Infof(format string, v ...interface{}) {
	sl._Infof(format, v...)
}

func (sl *StdoutLogger) Warnf(format string, v ...interface{}) {
	sl._Warnf(format, v...)
}

func (sl *StdoutLogger) Errorf(format string, v ...interface{}) {
	sl._Errorf(format, v...)
}

func (sl *StdoutLogger) Fatalf(format string, v ...interface{}) {
	sl._Fatalf(format, v...)
	os.Exit(1)
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
