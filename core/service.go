package core

import (
	"github.com/louisevanderlith/husk/validation"
	"time"
)

type Service struct {
	Duration    time.Duration
	StartTime   time.Time
	Location    string
	Description string `hsk:"size(256)"`
}

func (o Service) Valid() error {
	return validation.Struct(o)
}
