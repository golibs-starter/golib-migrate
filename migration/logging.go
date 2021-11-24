package migration

import (
	"github.com/golang-migrate/migrate/v4"
	"gitlab.id.vin/vincart/golib/log"
	"strings"
)

type Logging struct {
}

func NewLogging() migrate.Logger {
	return &Logging{}
}

func (l Logging) Printf(format string, v ...interface{}) {
	log.Infof(strings.Trim(format, "\n"), v...)
}

func (l Logging) Verbose() bool {
	return true
}
