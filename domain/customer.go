package domain

// business object
type Customer struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateofBirth string
	Status      string
}

// repository
type CustomerRepository interface {
	FindAll() ([]Customer, error)
	FindCustomerById(id string) (Customer, error)
	// UpdateCustomer(id int) (Customer, error)
	// DeleteCustomer(id int) (Customer, error)
}
