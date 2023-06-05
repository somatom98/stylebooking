package viewmodels

import (
	"github.com/somatom98/stylebooking/stylebooking_be/models"
)

type Service struct {
	ID          string  `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
}

func (s *Service) ToModel() models.Service {
	return models.Service{
		Name:        s.Name,
		Description: s.Description,
		Price:       s.Price,
	}
}

func (s *Service) FromModel(service models.Service) {
	s.ID = service.ID.Hex()
	s.Name = service.Name
	s.Description = service.Description
	s.Price = service.Price
}
