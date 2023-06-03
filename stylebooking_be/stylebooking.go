package stylebooking_be

import (
	"context"

	"github.com/somatom98/stylebooking/stylebooking_be/models"
)

type ServiceRepository interface {
	GetAll(context.Context) ([]models.Service, error)
	GetById(context.Context, string) (models.Service, error)
	Create(context.Context, models.Service) error
}
