package logger

import (
	"log"

	"github.com/w1png/urlshortener/utils"
)

var LoggerInstance Logger

type Logger interface {
  Log(s string, v ...interface{})
  Info(s string, v ...interface{})
  Error(s string, v ...interface{})
  Fatal(s string, v ...interface{})
}

func InitLogger() error {
  switch utils.ConfigInstance.LoggerType {
  case "console":
    LoggerInstance = NewConsoleLogger()
  default:
    log.Println("No logger type specified, using no logger")
    LoggerInstance = NewNoLogger()
  }

  return nil
}

