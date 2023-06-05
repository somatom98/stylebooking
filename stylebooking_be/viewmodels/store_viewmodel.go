package viewmodels

import (
	"time"

	"github.com/somatom98/stylebooking/stylebooking_be/models"
)

type Store struct {
	ID          string                               `json:"id,omitempty"`
	Name        string                               `json:"name,omitempty"`
	Description string                               `json:"description,omitempty"`
	Location    models.StoreLocation                 `json:"location,omitempty"`
	Hours       map[time.Weekday][]models.StoreHours `json:"hours,omitempty"`
	Services    []Service                            `json:"services,omitempty"`
}

func (s *Store) ToModel() models.Store {
	return models.Store{
		Name:        s.Name,
		Description: s.Description,
		Location:    s.Location,
		Hours:       s.Hours,
	}
}

func (s *Store) FromModel(store models.Store) {
	var services []Service
	for _, service := range store.Services {
		var s Service
		s.FromModel(service)
		services = append(services, s)
	}

	s.ID = store.ID.Hex()
	s.Name = store.Name
	s.Description = store.Description
	s.Location = store.Location
	s.Hours = store.Hours
	s.Services = services
}
