package domain

// business object
type Customer struct {
	Id          int    `json:"customer_id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateofBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}

// repository secondary port interface
type CustomerRepository interface {
	FindAll() ([]Customer, error)
	FindCustomerById(string) (*Customer, error) // use *pointer to be able to return "nil customer"
	// UpdateCustomer(id int) (Customer, error)
	DeleteCustomerById(id string) (int64, error) // delete customer, no customer returned, only iff error
	Create(Customer) (*Customer, error)
}
