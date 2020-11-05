package core

import "github.com/louisevanderlith/husk/validation"

type Service struct {
	Url string `hsk:"size(256)"`
}

func (o Service) Valid() error {
	return validation.Struct(o)
}
