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

type CustomerRepository interface {
	GetAll(context.Context) ([]models.Customer, error)
	GetById(context.Context, string) (models.Customer, error)
	GetByEmail(context.Context, string) (models.Customer, error)
	Create(context.Context, models.Customer) (string, error)
	Update(context.Context, string, models.Customer) error
	Delete(context.Context, string) error
}

type AuthenticationRepository interface {
	GetByCustomerId(context.Context, string) (models.Authentication, error)
	Create(context.Context, models.Authentication) error
	Delete(context.Context, string) error
	Update(context.Context, string, models.Authentication) error
}

type StoreService interface {
	GetAll(context.Context) ([]vm.Store, error)
	GetById(context.Context, string) (vm.Store, error)
	Create(context.Context, vm.Store) error
	AddService(context.Context, string, vm.Service) error
	UpdateService(context.Context, string, string, vm.Service) error
	DeleteService(context.Context, string, string) error
}

type CustomerService interface {
	GetById(context.Context, string) (models.Customer, error)
	SignUp(context.Context, vm.SignUpRequest) (vm.SignUpResponse, error)
	LogIn(context.Context, vm.SignInRequest) (vm.Token, error)
}

type AuthenticationService interface {
	Authenticate(context.Context, string, string) (vm.Token, error)
	Refresh(context.Context, string, string) (vm.Token, error)
	CreatePassword(context.Context, string, string) error
	UpdatePassword(context.Context, string, string) error
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

type ErrWrongPassword struct{}

func (e ErrWrongPassword) Error() string {
	return "Wrong password"
}

type ErrCustomerNotFound struct {
	Id string
}

func (e ErrCustomerNotFound) Error() string {
	return "Customer with id " + e.Id + " not found"
}

type ErrInvalidToken struct{}

func (e ErrInvalidToken) Error() string {
	return "Invalid token"
}
