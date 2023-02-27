package service

import domain "github.com/oscaralmgren/rest-banking-api-go-tutorial/domain/sale"

type DefaultSaleService struct {
	repo domain.SaleRepository
}

func NewSaleService(repositry domain.SaleRepository) DefaultSaleService {
	return DefaultSaleService{repositry}
}

// Interface definition
type SaleService interface {
	FindOne() (*domain.Sale, error)
	FindAll() ([]domain.Sale, error)
}

// Interface implementations from here.
func (s DefaultSaleService) FindOne() (*domain.Sale, error) {
	sale, err := s.repo.FindOne()
	if err != nil {
		return nil, err
	}
	return sale, nil
}

func (s DefaultSaleService) FindAll() ([]domain.Sale, error) {
	sales, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return sales, nil
}
