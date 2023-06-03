package viewmodels

import (
	"time"

	"github.com/somatom98/stylebooking/stylebooking_be/models"
)

type Store struct {
	Name        string                               `json:"name,omitempty"`
	Description string                               `json:"description,omitempty"`
	Location    models.StoreLocation                 `json:"location,omitempty"`
	Hours       map[time.Weekday][]models.StoreHours `json:"hours,omitempty"`
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
	s.Name = store.Name
	s.Description = store.Description
	s.Location = store.Location
	s.Hours = store.Hours
}
