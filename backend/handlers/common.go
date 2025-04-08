package handlers

import (
	"io"
	"log"
	"myservices/config"
)

var Log = log.New(io.Discard, "[handler]: ", config.LogFlags)
