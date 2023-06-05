package stylebooking_be

import (
	"context"

	"github.com/somatom98/stylebooking/stylebooking_be/models"
	vm "github.com/somatom98/stylebooking/stylebooking_be/viewmodels"
)

type ServiceRepository interface {
	GetAll(context.Context) ([]models.Service, error)
	GetById(context.Context, string) (models.Service, error)
	Create(context.Context, models.Service) error
}

type StoreRepository interface {
	GetAll(context.Context) ([]models.Store, error)
	GetById(context.Context, string) (models.Store, error)
	Create(context.Context, models.Store) error
	AddService(context.Context, string, models.Service) error
	UpdateService(context.Context, string, string, models.Service) error
	DeleteService(context.Context, string, string) error
}

type StoreService interface {
	GetAll(context.Context) ([]vm.Store, error)
	GetById(context.Context, string) (vm.Store, error)
	Create(context.Context, vm.Store) error
	AddService(context.Context, string, vm.Service) error
	UpdateService(context.Context, string, string, vm.Service) error
	DeleteService(context.Context, string, string) error
}

type ErrStoreNotFound struct {
	Id string
}

func (e ErrStoreNotFound) Error() string {
	return "Store with id " + e.Id + " not found"
}

type ErrServiceNotFound struct {
	Id      string
	StoreId string
}

func (e ErrServiceNotFound) Error() string {
	return "Service with id " + e.Id + " not found in store with id " + e.StoreId
}
