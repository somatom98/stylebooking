package services

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	sb "github.com/somatom98/stylebooking/stylebooking_be"
	"github.com/somatom98/stylebooking/stylebooking_be/config"
	"github.com/somatom98/stylebooking/stylebooking_be/models"
	vm "github.com/somatom98/stylebooking/stylebooking_be/viewmodels"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationService struct {
	authenticationRepository sb.AuthenticationRepository
	customerRepository       sb.CustomerRepository
	conf                     config.JwtConfiguration
}

func NewAuthenticationService(authenticationRepository sb.AuthenticationRepository, customerRepository sb.CustomerRepository, conf config.JwtConfiguration) *AuthenticationService {
	return &AuthenticationService{
		authenticationRepository: authenticationRepository,
		customerRepository:       customerRepository,
		conf:                     conf,
	}
}

func (s *AuthenticationService) Authenticate(ctx context.Context, customerId string, password string) (vm.Token, error) {
	tokens := vm.Token{}
	authentication, err := s.authenticationRepository.GetByCustomerId(ctx, customerId)
	if err != nil {
		return tokens, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(authentication.Password), []byte(password))
	if err != nil {
		return tokens, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customerId": customerId,
		"iss":        s.conf.Issuer,
		"aud":        s.conf.Audience,
		"exp":        s.conf.AccessTokenDuration,
	})
	tokens.Authentication, err = token.SignedString([]byte(s.conf.Secret))
	if err != nil {
		return tokens, err
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customerId": customerId,
		"iss":        s.conf.Issuer,
		"aud":        s.conf.Audience,
		"exp":        s.conf.RefreshTokenDuration,
	})
	tokens.Refresh, err = token.SignedString([]byte(s.conf.Secret))
	if err != nil {
		return tokens, err
	}
	return tokens, nil
}

func (s *AuthenticationService) Refresh(ctx context.Context, customerId string, refreshToken string) (vm.Token, error) {
	tokens := vm.Token{}
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.conf.Secret), nil
	})
	if err != nil {
		return tokens, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return tokens, sb.ErrInvalidToken{}
	}
	if claims["customerId"] != customerId {
		return tokens, sb.ErrInvalidToken{}
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customerId": customerId,
		"iss":        s.conf.Issuer,
		"aud":        s.conf.Audience,
		"exp":        s.conf.AccessTokenDuration,
	})
	tokens.Authentication, err = token.SignedString([]byte(s.conf.Secret))
	if err != nil {
		return tokens, err
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customerId": customerId,
		"iss":        s.conf.Issuer,
		"aud":        s.conf.Audience,
		"exp":        s.conf.RefreshTokenDuration,
	})
	tokens.Refresh, err = token.SignedString([]byte(s.conf.Secret))
	if err != nil {
		return tokens, err
	}
	return tokens, nil
}

func (s *AuthenticationService) CreatePassword(ctx context.Context, customerId string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	authentication := models.Authentication{
		CustomerId: customerId,
		Password:   string(hashedPassword),
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
	}
	err = s.authenticationRepository.Create(ctx, authentication)
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthenticationService) UpdatePassword(ctx context.Context, customerId string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	authentication := models.Authentication{
		CustomerId: customerId,
		Password:   string(hashedPassword),
		UpdatedAt:  time.Now().UTC(),
	}
	err = s.authenticationRepository.Update(ctx, customerId, authentication)
	if err != nil {
		return err
	}
	return nil
}
