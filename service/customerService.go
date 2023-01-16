package service

import "github.com/oscaralmgren/rest-banking-api-go-tutorial/domain"

// service - primary port
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomerById(id string) (*domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

// implement Service interface
func (s DefaultCustomerService) GetCustomerById(id string) (*domain.Customer, error) {
	c, err := s.repo.FindCustomerById(id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func NewCustomerService(repositry domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repositry}
}
