package middlewares

import (
	"log"
	"net/http"
	"time"
)

type customWriter struct {
	http.ResponseWriter
	statusCode int
}

func (cw *customWriter) WriteHeader(statusCode int) {
	cw.ResponseWriter.WriteHeader(statusCode)
	cw.statusCode = statusCode
}

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		cw := &customWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(cw, r)

		duration := time.Since(start)

		statusCode := cw.statusCode

		log.Printf("%s %s %s %d %s %s", r.Method, r.URL.Path, r.Proto, statusCode, http.StatusText(statusCode), duration)
	})
}
