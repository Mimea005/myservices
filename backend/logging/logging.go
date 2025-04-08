package logging

import (
	"io"
	"log"
	"myservices/config"
	"myservices/handlers"
	"myservices/middleware"
	"myservices/services"
	"os"
)


// Configure outputstreams for different log scopes 
func Configure(config config.Server) {
	for _, filter := range config.LogFilters {
		switch filter {
		case "handler":
			handlers.Log.SetOutput(os.Stderr)
			fallthrough
		case "service":
			services.Log.SetOutput(os.Stderr)
		case "router":
			middleware.RouterLog.SetOutput(os.Stderr)
		case "-info":
			log.Default().SetOutput(io.Discard)
		default:
		}
	}
}
