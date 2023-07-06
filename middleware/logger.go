package middleware

import (
	"net/http"
	"time"

	"github.com/w1png/urlshortener/logger"
)

func LoggingMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    startTime := time.Now()
    next.ServeHTTP(w, r)

    logger.LoggerInstance.Info(" [%s] %s Took: %s", r.Method, r.RequestURI, time.Since(startTime))
  })
}

