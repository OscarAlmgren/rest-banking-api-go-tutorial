package domain

import "time"

type Sale struct {
	Item     string    `json:"item"`
	Price    float64   `json:"price"`
	Quantity int       `json:"quantity"`
	Date     time.Time `json:"date"`
}

type SaleRepository interface {
	FindOne() (*Sale, error)
	FindAll() ([]Sale, error)
}
