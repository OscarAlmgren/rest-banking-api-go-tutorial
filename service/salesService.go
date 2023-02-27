package service

import domain "github.com/oscaralmgren/rest-banking-api-go-tutorial/domain/sale"

// Interface definition
type SaleService interface {
	FindOne() (*domain.Sale, error)
}

type DefaultSaleService struct {
	repo domain.SaleRepository
}

func NewSaleService(repositry domain.SaleRepository) DefaultSaleService {
	return DefaultSaleService{repositry}
}

// Interface implementations from here.
func (s DefaultSaleService) FindOne() (*domain.Sale, error) {
	sale, err := s.repo.FindOne()
	if err != nil {
		return nil, err
	}
	return sale, nil
}
