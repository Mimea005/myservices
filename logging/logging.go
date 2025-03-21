package logging

import (
	"io"
	"myservices/common"
	"myservices/handlers"
	"myservices/middleware"
	"myservices/services"
	"os"
)

func ConfigureLogging() {

	for _, filter := range common.Config.LogFilters {
		switch filter {
		case "handler":
			handlers.Log.SetOutput(os.Stderr)
			fallthrough
		case "service":
			services.Log.SetOutput(os.Stderr)
		case "router":
			middleware.RouterLog.SetOutput(os.Stderr)
		case "-info":
			common.Log.SetOutput(io.Discard)
		default:
		}
	}

}
