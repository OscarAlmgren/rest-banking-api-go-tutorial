package service

import "github.com/oscaralmgren/rest-banking-api-go-tutorial/domain"

// service - primary port
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomerById(string) (*domain.Customer, error) // use *pointer to be able to return "nil customer"
	DeleteCustomerById(string) (int64, error)
	Create(domain.Customer) (*domain.Customer, error)
}

// DefaultCustomerService service connects the domain as a struct property
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

func (s DefaultCustomerService) DeleteCustomerById(id string) (int64, error) {
	i, err := s.repo.DeleteCustomerById(id)
	if err != nil {
		return -1, err
	}
	return i, nil
}

func (s DefaultCustomerService) Create(domain.Customer) (*domain.Customer, error) {
	var c *domain.Customer
	// TODO implementera den här anropas på rätt sätt

	return c, nil
}

func NewCustomerService(repositry domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repositry}
}
