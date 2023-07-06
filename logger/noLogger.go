package logger

import "os"

type NoLogger struct {}

func (l *NoLogger) Log(s string, v ...interface{}) {}
func (l *NoLogger) Info(s string, v ...interface{}) {}
func (l *NoLogger) Error(s string, v ...interface{}) {}
func (l *NoLogger) Fatal(s string, v ...interface{}) {
  os.Exit(1)
}

func NewNoLogger() Logger {
  return &NoLogger{}
}

