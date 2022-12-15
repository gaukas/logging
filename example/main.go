package main

import "github.com/gaukas/logging"

func main() {
	stdlogger := logging.DefaultStderrLogger(logging.LOG_DEBUG)
	stdlogger.Debugf("This is a debug message")
	stdlogger.Infof("This is an info message")
	stdlogger.Warnf("This is a warning message")
	stdlogger.Errorf("This is an error message")

	filelogger := logging.DefaultFileLogger("test.log", logging.LOG_DEBUG)
	filelogger.Debugf("This is a debug message")
	filelogger.Infof("This is an info message")
	filelogger.Warnf("This is a warning message")
	filelogger.Errorf("This is an error message")

	// combining the two loggers
	multilogger := logging.NewMultiLogger(stdlogger, filelogger)
	multilogger.Debugf("This is a debug message")
	multilogger.Infof("This is an info message")
	multilogger.Warnf("This is a warning message")
	multilogger.Errorf("This is an error message")

	// Fatalf will exit the program
	multilogger.Fatalf("This is a fatal message")
}
