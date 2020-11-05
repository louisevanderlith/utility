package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/op"
	"github.com/louisevanderlith/husk/records"
)

type UtilityContext interface {
	GetService(key hsk.Key) (Service, error)
	FindServices(page, size int) (records.Page, error)
	CreateService(obj Service) (hsk.Key, error)
	UpdateService(key hsk.Key, obj Service) error
}

func CreateContext() UtilityContext {
	ctx = context{Services: husk.NewTable(Service{})}

	return ctx
}

func Shutdown() {
	ctx.Services.Save()
}

func Context() UtilityContext {
	return ctx
}

type context struct {
	Services husk.Table
}

var ctx context

func (c context) GetService(key hsk.Key) (Service, error) {
	rec, err := c.Services.FindByKey(key)

	if err != nil {
		return Service{}, err
	}

	return rec.GetValue().(Service), nil
}

func (c context) FindServices(page, size int) (records.Page, error) {
	return c.Services.Find(page, size, op.Everything())
}

func (c context) CreateService(obj Service) (hsk.Key, error) {
	return c.Services.Create(obj)
}

func (c context) UpdateService(key hsk.Key, obj Service) error {
	return c.Services.Update(key, obj)
}
