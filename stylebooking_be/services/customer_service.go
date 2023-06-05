package services

import (
	"context"

	sb "github.com/somatom98/stylebooking/stylebooking_be"
	m "github.com/somatom98/stylebooking/stylebooking_be/models"
	vm "github.com/somatom98/stylebooking/stylebooking_be/viewmodels"
)

type CustomerService struct {
	customerRepository    sb.CustomerRepository
	authenticationService sb.AuthenticationService
}

func NewCustomerService(customerRepository sb.CustomerRepository, authenticationService sb.AuthenticationService) *CustomerService {
	return &CustomerService{
		customerRepository:    customerRepository,
		authenticationService: authenticationService,
	}
}

func (s *CustomerService) GetById(ctx context.Context, id string) (m.Customer, error) {
	customer, err := s.customerRepository.GetById(ctx, id)
	if err != nil {
		return m.Customer{}, err
	}

	return customer, nil
}

func (s *CustomerService) SignUp(ctx context.Context, request vm.SignUpRequest) (vm.SignUpResponse, error) {
	customer := m.Customer{
		Email:    request.Email,
		Password: request.Password,
		Name:     request.Name,
		Surname:  request.Surname,
		Phone:    request.Phone,
	}

	id, err := s.customerRepository.Create(ctx, customer)
	if err != nil {
		return vm.SignUpResponse{}, err
	}

	s.authenticationService.CreatePassword(ctx, id, request.Password)

	return vm.SignUpResponse{ID: id}, nil
}

func (s *CustomerService) SignIn(ctx context.Context, request vm.SignInRequest) (vm.Token, error) {
	customer, err := s.customerRepository.GetByEmail(ctx, request.Email)
	if err != nil {
		return vm.Token{}, err
	}

	token, err := s.authenticationService.Authenticate(ctx, customer.ID.Hex(), request.Password)
	if err != nil {
		return vm.Token{}, err
	}

	return token, nil
}
