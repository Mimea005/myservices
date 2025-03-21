package common

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Filters []string

func (l *Filters) String() string {
	str := new(string)
	for _, filter := range *l {
		*str = fmt.Sprintf("%s,%s", *str, filter)
	}
	return *str
}

func (l *Filters) Set(value string) error {
	filters := strings.Split(value, ",")
	*l = filters
	return nil
}

// Common logging format for server
const LogFlags = log.Lmsgprefix | log.LstdFlags

// For logging non debug info
var Log = log.New(os.Stdout, "[ ]: ", LogFlags)
