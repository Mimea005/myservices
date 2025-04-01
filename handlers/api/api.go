package api

import (
	"bytes"
	"myservices/handlers"
	"myservices/middleware"
	"myservices/router"
	"myservices/services"
	"net/http"
)

// Subrouter for all api endpoints (except for /health (at least for now))
var ApiRouter = router.NewRouter()

func init() {
	ApiRouter.UseMiddleware(middleware.ContentType("text/html"))

	ApiRouter.HandleFunc("GET /ping", ping)
}


func ping(w http.ResponseWriter, r *http.Request) {
	handlers.Log.Println("Pinged")

	buff := bytes.Buffer{}

	err := services.Templates.ExecuteTemplate(&buff, "ping.html", r.RemoteAddr)
	if err != nil {
		handlers.Log.Println(err.Error())
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	} else {
		w.Write(buff.Bytes())
	}
}
