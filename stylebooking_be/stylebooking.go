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
}

type StoreService interface {
	GetAll(context.Context) ([]vm.Store, error)
	GetById(context.Context, string) (vm.Store, error)
	Create(context.Context, vm.Store) error
}
