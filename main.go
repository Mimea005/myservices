package main

import (
	"flag"
	"myservices/common"
	"myservices/handlers"
	"myservices/logging"
	"myservices/middleware"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	flag.Parse()
	logging.ConfigureLogging()
	common.Log.Println("Start")

	router := mux.NewRouter()

	router.HandleFunc("/health", handlers.Health)

	// Add a handler to log requests not matched
	router.NotFoundHandler = router.NewRoute().HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.Log.Println("Not found")
		http.NotFound(w, r)
	}).GetHandler()

	router.Use(middleware.LogRequests)

	if l, err := net.Listen("tcp", common.Config.BindAddress); err != nil {
		common.Log.Fatal(err.Error())
	} else {
		common.Log.Println("Serving on: " + common.Config.BindAddress)
		common.Log.Fatal(http.Serve(l, router))
	}
}
