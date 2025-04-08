package services

import (
	"io"
	"log"
	"myservices/config"
)

var Log = log.New(io.Discard, "[service]: ", config.LogFlags)
