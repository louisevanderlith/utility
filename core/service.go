package core

import (
	"github.com/louisevanderlith/husk/validation"
	"time"
)

type Service struct {
	Duration    time.Duration
	Location    string
	Description string `hsk:"size(256)"`
}

func (o Service) Valid() error {
	return validation.Struct(o)
}
