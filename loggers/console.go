package loggers

import (
	"log"
	"os"
)

type ConsoleLogger struct {
	debug *log.Logger
	warn  *log.Logger
	info  *log.Logger
	error *log.Logger
}

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{
		debug: log.New(os.Stdout, "Debug: ", log.LstdFlags),
		warn:  log.New(os.Stdout, "Warn: ", log.LstdFlags),
		info:  log.New(os.Stdout, "Info: ", log.LstdFlags),
		error: log.New(os.Stderr, "Error: ", log.LstdFlags),
	}
}

func (l *ConsoleLogger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *ConsoleLogger) Warn(v ...interface{}) {
	l.warn.Println(v...)
}

func (l *ConsoleLogger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *ConsoleLogger) Error(v ...interface{}) {
	l.error.Println(v...)
}
