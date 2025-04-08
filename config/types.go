package config

import (
	"fmt"
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

type Server struct {
	BindAddress string
	LogFilters Filters
	CoolifyURL string
	CoolifyToken string
}
