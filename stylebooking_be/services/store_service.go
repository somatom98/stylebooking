package services

import (
	"context"

	sb "github.com/somatom98/stylebooking/stylebooking_be"
	vm "github.com/somatom98/stylebooking/stylebooking_be/viewmodels"
)

type StoreService struct {
	storeRepository sb.StoreRepository
}

func NewStoreService(storeRepository sb.StoreRepository) *StoreService {
	return &StoreService{
		storeRepository: storeRepository,
	}
}

func (s *StoreService) GetAll(ctx context.Context) ([]vm.Store, error) {
	stores, err := s.storeRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var viewmodels []vm.Store
	for _, store := range stores {
		viewmodels = append(viewmodels, vm.Store{
			Name:        store.Name,
			Description: store.Description,
			Location:    store.Location,
			Hours:       store.Hours,
		})
	}

	return viewmodels, nil
}

func (s *StoreService) GetById(ctx context.Context, id string) (vm.Store, error) {
	store, err := s.storeRepository.GetById(ctx, id)
	if err != nil {
		return vm.Store{}, err
	}

	storeVm := vm.Store{}
	storeVm.FromModel(store)

	return storeVm, nil
}

func (s *StoreService) Create(ctx context.Context, store vm.Store) error {
	return s.storeRepository.Create(ctx, store.ToModel())
}
