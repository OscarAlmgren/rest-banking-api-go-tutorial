package domain

import "time"

type Sale struct {
	// ID string `json:"id", omitempty`
	Item     string    `json:"item"`
	Price    float64   `json:"price"`
	Quantity int       `json:"quantity"`
	Date     time.Time `json:"date"`
}

type SaleRepository interface {
	FindOne() (*Sale, error)
	FindAll() ([]Sale, error)
	Create(Sale) (string, error)
	Delete(id string) (int64, error)
}
