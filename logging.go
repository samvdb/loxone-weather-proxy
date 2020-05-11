package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"runtime/debug"
	"time"
)

// responseWriter is a minimal wrapper for http.ResponseWriter that allows the
// written HTTP status code to be captured for logging.
type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true

	return
}

func LoggingMiddlewar(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.WithFields(log.Fields{
					"trace": debug.Stack(),
					"err":   err,
				})
			}
		}()

		start := time.Now()
		wrapped := wrapResponseWriter(w)
		next.ServeHTTP(wrapped, r)
		log.WithFields(log.Fields{
			"status":   wrapped.status,
			"method":   r.Method,
			"path":     r.URL.EscapedPath(),
			"asl":      r.URL.Query().Get("asl"),
			"coord":    r.URL.Query().Get("coord"),
			"format":   r.URL.Query().Get("format"),
			"duration": time.Since(start),
		}).Info("received request")
	}

	return http.HandlerFunc(fn)
}
