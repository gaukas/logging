package logging

import (
	"log"
	"os"
)

type FileLogger struct {
	logger *log.Logger // embedded

	_Debugf func(format string, v ...interface{})
	_Infof  func(format string, v ...interface{})
	_Warnf  func(format string, v ...interface{})
	_Errorf func(format string, v ...interface{})
	_Fatalf func(format string, v ...interface{})
}

func NewFileLogger(filename string, prefix string, flag int) *FileLogger {
	logFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil
	}
	return &FileLogger{logger: log.New(logFile, prefix, flag)}
}

func DefaultFileLogger(filename string, loggingLevel uint8) *FileLogger {
	logger := NewFileLogger(filename, "", log.LstdFlags)

	logger.SetLoggingLevel(loggingLevel)

	return logger
}

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

func (fl *FileLogger) Debugf(format string, v ...interface{}) {
	fl._Debugf(format, v...)
}

func (fl *FileLogger) Infof(format string, v ...interface{}) {
	fl._Infof(format, v...)
}

func (fl *FileLogger) Warnf(format string, v ...interface{}) {
	fl._Warnf(format, v...)
}

func (fl *FileLogger) Errorf(format string, v ...interface{}) {
	fl._Errorf(format, v...)
}

func (fl *FileLogger) Fatalf(format string, v ...interface{}) {
	fl._Fatalf(format, v...)
	os.Exit(1)
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
	fl.logger.Fatalf("[FATAL] "+format, v...)
}
