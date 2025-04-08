package config

import (
	"log"
	"os"
)

// Common logging format for server
const LogFlags = log.Lmsgprefix | log.LstdFlags

// For logging non debug info
var Log = log.New(os.Stdout, "[ ]: ", LogFlags)
