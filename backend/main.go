package main

import (
	"flag"
	"log"
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
	logging.Configure(Config)
	log.Println("Start")

	router := router.NewRouter()
	router.UseMiddleware(middleware.LogRequests)

	// Register routes in api package
	router.OnSubPath("/x/", api.ApiRouter)

	router.HandleFunc("GET /health", handlers.Health)

	// Default handler to return not found
	router.HandleFunc("/", http.NotFound)



	if l, err := net.Listen("tcp", Config.BindAddress); err != nil {
		log.Fatal(err.Error())
	} else {
		log.Println("Serving on: " + Config.BindAddress)
		log.Fatal(http.Serve(l, router))
	}
}
