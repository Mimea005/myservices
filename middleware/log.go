package middleware

import (
	"io"
	"log"
	"myservices/common"
	"net/http"
)

type logResponseWriter struct {
	http.ResponseWriter
	status int
}

func newLogResponseWriter(w http.ResponseWriter) *logResponseWriter {
	return &logResponseWriter{w, http.StatusOK}
}

func (w *logResponseWriter) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

var RouterLog = log.New(io.Discard, "[router]: ", common.LogFlags)

func LogRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r * http.Request)  {
		writer := newLogResponseWriter(w)
		RouterLog.Printf("%s -> %s", r.RemoteAddr, r.RequestURI)
		next.ServeHTTP(writer, r)
		RouterLog.Printf("%s <- %d", r.RemoteAddr, writer.status)
	})
}
