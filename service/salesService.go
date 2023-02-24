package service

import "github.com/oscaralmgren/rest-banking-api-go-tutorial/domain"

type SalesService interface {
	FindOne() (*domain.Sale, error)
}

type DefaultSaleService struct {
	repo domain.SaleRepository
}

func (s DefaultSaleService) FindOne() (*domain.Sale, error) {
	sale, err := s.repo.FindOne()
	if err != nil {
		return nil, err
	}
	return sale, nil
}

func NewSaleService(repositry domain.SaleRepository) DefaultSaleService {
	return DefaultSaleService{repositry}
}
