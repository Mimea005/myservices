package main

import (
	"flag"
	"myservices/config"
)

// Runtime configuration variables. Currently set as arguments
var Config = config.Server {
	BindAddress: "",
	LogFilters: []string{"handler", "service", "router"},
	CoolifyURL: "",
	CoolifyToken: "",
}

func init() {
	flag.Var(&Config.LogFilters, "filter", "filter the log by category")
	flag.StringVar(&Config.BindAddress, "bind", ":8080", "address to bind to")
	flag.StringVar(&Config.CoolifyURL, "coolify-url", "", "coolify url")
	flag.StringVar(&Config.CoolifyToken, "coolify-token", "", "Bearer token for coolify")
}
