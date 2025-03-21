package handlers

import (
	"io"
	"log"
	"myservices/common"
)

var Log = log.New(io.Discard, "[handler]: ", common.LogFlags)
