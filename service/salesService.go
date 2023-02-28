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
	Create(domain.Sale) (string, error)
	Delete(id string) (int64, error)
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

func (s DefaultSaleService) Create(sale domain.Sale) (string, error) {
	inserted, err := s.repo.Create(sale)
	if err != nil {
		return err.Error(), err
	}
	return inserted, nil
}

func (s DefaultSaleService) Delete(id string) (int64, error) {
	deleted, err := s.repo.Delete(id)
	if err != nil {
		return -1, err
	}
	return deleted, nil
}
