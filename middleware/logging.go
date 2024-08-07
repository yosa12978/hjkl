package middleware

import (
	"net/http"
	"time"

	"github.com/yosa12978/hjkl/logging"
)

func Latency(logger logging.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			latency := time.Since(start).Microseconds()
			logger.Info(
				"incoming request",
				"method", r.Method,
				"latency_micro", latency,
				"endpoint", r.URL.Path,
				// Solution for status code logging is to use negroni library ResponseWriter (https://github.com/urfave/negroni)
			)
		})
	}
}
