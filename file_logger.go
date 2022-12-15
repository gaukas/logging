package logging

import (
	"log"
	"os"
)

// FileLogger defines a logger that writes to a file. It implements the Logger interface.
type FileLogger struct {
	logger *log.Logger // embedded

	_Debugf func(format string, v ...interface{})
	_Infof  func(format string, v ...interface{})
	_Warnf  func(format string, v ...interface{})
	_Errorf func(format string, v ...interface{})
	_Fatalf func(format string, v ...interface{})
}

// NewFileLogger creates a new FileLogger with the given filename, prefix and flag.
func NewFileLogger(filename string, prefix string, flag int) *FileLogger {
	logFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil
	}
	return &FileLogger{logger: log.New(logFile, prefix, flag)}
}

// DefaultFileLogger creates a new FileLogger with the given filename at the given logging level.
// It sets the prefix to "" and the flag to log.LstdFlags.
func DefaultFileLogger(filename string, loggingLevel uint8) *FileLogger {
	logger := NewFileLogger(filename, "", log.LstdFlags)
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
func (fl *FileLogger) SetLoggingLevel(level uint8) {
	if level <= LOG_DEBUG {
		fl._Debugf = fl._debugf
	} else {
		fl._Debugf = func(format string, v ...interface{}) {}
	}

	if level <= LOG_INFO {
		fl._Infof = fl._infof
	} else {
		fl._Infof = func(format string, v ...interface{}) {}
	}

	if level <= LOG_WARN {
		fl._Warnf = fl._warnf
	} else {
		fl._Warnf = func(format string, v ...interface{}) {}
	}

	if level <= LOG_ERROR {
		fl._Errorf = fl._errorf
	} else {
		fl._Errorf = func(format string, v ...interface{}) {}
	}

	if level <= LOG_FATAL {
		fl._Fatalf = fl._fatalf
	} else {
		fl._Fatalf = func(format string, v ...interface{}) {}
	}
}

// Debugf implements the Logger interface.
func (fl *FileLogger) Debugf(format string, v ...interface{}) {
	fl._Debugf(format, v...)
}

// Infof implements the Logger interface.
func (fl *FileLogger) Infof(format string, v ...interface{}) {
	fl._Infof(format, v...)
}

// Warnf implements the Logger interface.
func (fl *FileLogger) Warnf(format string, v ...interface{}) {
	fl._Warnf(format, v...)
}

// Errorf implements the Logger interface.
func (fl *FileLogger) Errorf(format string, v ...interface{}) {
	fl._Errorf(format, v...)
}

// Fatalf implements the Logger interface. It calls os.Exit(1) after logging
// regardless of the logging level.
func (fl *FileLogger) Fatalf(format string, v ...interface{}) {
	fl._Fatalf(format, v...)
	os.Exit(1)
}

// DeferredFatalf implements the CompatibleLogger interface. It does everything
// that Fatalf does, except it does not call os.Exit(1).
func (fl *FileLogger) DeferredFatalf(format string, v ...interface{}) {
	fl._Fatalf(format, v...)
}

func (fl *FileLogger) _debugf(format string, v ...interface{}) {
	fl.logger.Printf("[DEBUG] "+format, v...)
}

func (fl *FileLogger) _infof(format string, v ...interface{}) {
	fl.logger.Printf("[INFO] "+format, v...)
}

func (fl *FileLogger) _warnf(format string, v ...interface{}) {
	fl.logger.Printf("[WARN] "+format, v...)
}

func (fl *FileLogger) _errorf(format string, v ...interface{}) {
	fl.logger.Printf("[ERROR] "+format, v...)
}

func (fl *FileLogger) _fatalf(format string, v ...interface{}) {
	fl.logger.Printf("[FATAL] "+format, v...)
}
