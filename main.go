package main

import (
	"net/http"
	"services/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/health", handlers.Health)


	log.Fatal(http.ListenAndServe(":8080", router))
}
