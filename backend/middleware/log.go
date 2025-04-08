package middleware

import (
	"io"
	"log"
	"myservices/config"
	"net/http"
)

// middleware uses its own log scope
var RouterLog = log.New(io.Discard, "[router]: ", config.LogFlags)

// ResponseWriter Wrapper to keep information of the result for logging purposes
type logResponseWriter struct {
	http.ResponseWriter
	status int
}

func wrapResponseWriter(w http.ResponseWriter) *logResponseWriter {
	return &logResponseWriter{w, http.StatusOK}
}

func (w *logResponseWriter) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// Middleware to log request origin and response statuscode
func LogRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r * http.Request)  {
		writer := wrapResponseWriter(w)

		RouterLog.Printf("%s -> %s", r.RemoteAddr, r.RequestURI)

		next.ServeHTTP(writer, r)

		RouterLog.Printf("%s <- %d", r.RemoteAddr, writer.status)
	})
}
