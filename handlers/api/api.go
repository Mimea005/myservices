package api

import (
	"myservices/handlers"
	"myservices/middleware"
	"myservices/services"
	"net/http"

	"github.com/gorilla/mux"
)

// Subrouter for all api endpoints (except for /health (at least for now))
func ConfigureRoutes(r *mux.Router) {

	// Make sure all repsonses under this router are answered as html content
	r.Use(middleware.ContentHtml)

	r.HandleFunc("/ping", ping)
	r.HandleFunc("/services", listServices)
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

