package logger

import "log"

type ConsoleLogger struct {}

func (l *ConsoleLogger) Log(s string, v ...interface{}) {
  log.Printf(s, v...)
}

func (l *ConsoleLogger) Info(s string, v ...interface{}) {
  log.Printf("[INFO]" + s, v...)
}

func (l *ConsoleLogger) Error(s string, v ...interface{}) {
  log.Printf("[ERROR]" + s, v...)
}

func (l *ConsoleLogger) Fatal(s string, v ...interface{}) {
  log.Fatalf("[FATAL]" + s, v...)
}

func NewConsoleLogger() Logger {
  return &ConsoleLogger{}
}


