package main

import (
	"flag"
	"myservices/common"
	"myservices/handlers"
	"myservices/handlers/api"
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
	router.Use(middleware.LogRequests)

	// Register routes in api package
	api.ConfigureRoutes(router.PathPrefix("/x").Subrouter())

	router.HandleFunc("/health", handlers.Health)

	// Add a handler to log requests not matched
	router.NotFoundHandler = router.NewRoute().HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.Log.Println("Not found")
		http.NotFound(w, r)
	}).GetHandler()

	if l, err := net.Listen("tcp", common.Config.BindAddress); err != nil {
		common.Log.Fatal(err.Error())
	} else {
		common.Log.Println("Serving on: " + common.Config.BindAddress)
		common.Log.Fatal(http.Serve(l, router))
	}
}
