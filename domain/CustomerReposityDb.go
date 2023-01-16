package domain

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

// FindCustomerById implements CustomerRepository
func (CustomerRepositoryDb) FindCustomerById(id string) (Customer, error) {
	panic("unimplemented")
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {

	findAllSql := "SELECT `customer_id`, `name`, `date_of_birth`, `city`, `zipcode`, `status` FROM customers"

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		return nil, err
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.DateofBirth, &c.City, &c.Zipcode, &c.Status)
		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}

	return customers, nil
}

// func (d CustomerRepositoryDb) FindCustomerById(id string) (Customer, error) {

// }
func NewCustomerRepositoryDb() CustomerRepositoryDb {
	// return CustomerRepositoryDb{client: dbClient}  dbClient *sql.DB
	client, err := sqlx.Open("mysql", "root:oscar-camp-tutorial@tcp(192.168.205.5:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxIdleTime(time.Minute * 2)
	client.SetMaxOpenConns(5)
	client.SetMaxIdleConns(5)
	return CustomerRepositoryDb{client: client}
}
