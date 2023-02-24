package domain

import "time"

type Sale struct {
	Item     string    `json:"item"`
	Price    int       `json:"price"`
	Quantity int       `json:"quantity"`
	Date     time.Time `json:"date"`
}

type SaleRepository interface {
	FindOne() (*Sale, error)
}
