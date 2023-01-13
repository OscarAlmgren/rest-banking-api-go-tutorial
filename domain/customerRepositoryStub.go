package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{1, "Oscar", "Nacka", "13131", "1982-01-08", "active"},
		{2, "Ida", "Kvarnholmen", "13131", "1991-03-13", "active"},
		{3, "Hugo", "Lilla Kvarnholmen", "12345", "2020-01-18", "inactive"},
		{3, "Henry", "Kvarnholmensparken", "67890", "2022-07-28", "inactive"},
	}

	return CustomerRepositoryStub{customers: customers}
}
