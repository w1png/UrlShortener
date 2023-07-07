package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
  httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
    Name: "urlshortener_http_duration_seconds",
    Help: "Duration of HTTP requests.",
  }, []string{"path"})
  httpRequests = promauto.NewCounterVec(prometheus.CounterOpts{
    Name: "urlshortener_http_requests_total",
    Help: "Total number of HTTP requests.",
  }, []string{"path", "method"})
)

func PrometheusDurationMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    route := mux.CurrentRoute(r)
    path, _ := route.GetPathTemplate()
    timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
    next.ServeHTTP(w, r)
    timer.ObserveDuration()
  })
}

func PrometheusCounterMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    route := mux.CurrentRoute(r)
    path, _ := route.GetPathTemplate()
    httpRequests.WithLabelValues(path, r.Method).Inc()
    next.ServeHTTP(w, r)
  })
}

