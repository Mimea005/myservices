package common

import (
	"flag"
)

type serverConfig struct {
	BindAddress string
	LogFilters Filters
}

// Runtime configuration variables. Currently set as arguments
var Config = serverConfig{
	BindAddress: "",
	LogFilters: []string{"handler", "service", "router"},
}

func init() {
	flag.Var(&Config.LogFilters, "filter", "filter the log by category")
	flag.StringVar(&Config.BindAddress, "bind", ":8080", "address to bind to")
}

