package middleware

import (
	"io"
	"log"
	"myservices/common"
	"net/http"
)

var RouterLog = log.New(io.Discard, "[router]: ", common.LogFlags)

func LogRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r * http.Request)  {
		RouterLog.Printf("%s -> %s", r.RemoteAddr, r.RequestURI)

		next.ServeHTTP(w, r)
	})
}
