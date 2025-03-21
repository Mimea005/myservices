package services

import (
	"io"
	"log"
	"myservices/common"
)

var Log = log.New(io.Discard, "[service]: ", common.LogFlags)
