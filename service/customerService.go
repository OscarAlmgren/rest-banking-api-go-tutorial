package service

import "github.com/oscaralmgren/rest-banking-api-go-tutorial/domain"

// service - primary port
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomerById(id string) (domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

// implement Service interface
func (s DefaultCustomerService) GetCustomerById(id string) (domain.Customer, error) {
	return s.repo.FindCustomerById(id)
}

func NewCustomerService(repositry domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repositry}
}
