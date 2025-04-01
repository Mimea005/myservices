package main

import (
	"flag"
	"myservices/common"
	"myservices/handlers"
	"myservices/handlers/api"
	"myservices/logging"
	"myservices/middleware"
	"myservices/router"
	"net"
	"net/http"
)

func main() {
	flag.Parse()
	logging.ConfigureLogging()
	common.Log.Println("Start")

	router := router.NewRouter()
	router.UseMiddleware(middleware.LogRequests)

	// Register routes in api package
	router.OnSubPath("/x/", api.ApiRouter)

	router.HandleFunc("GET /health", handlers.Health)

	// Default handler to return not found
	router.HandleFunc("/", http.NotFound)



	if l, err := net.Listen("tcp", common.Config.BindAddress); err != nil {
		common.Log.Fatal(err.Error())
	} else {
		common.Log.Println("Serving on: " + common.Config.BindAddress)
		common.Log.Fatal(http.Serve(l, router))
	}
}
