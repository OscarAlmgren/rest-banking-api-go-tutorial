package domain

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

// FindCustomerById implements CustomerRepository
func (d CustomerRepositoryDb) FindCustomerById(id string) (*Customer, error) {
	var c Customer
	customerFindSql := "SELECT customer_id as id, name, date_of_birth as dateofbirth, city, zipcode, status FROM customers WHERE customer_id = ?"

	err := d.client.Get(&c, customerFindSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			return nil, err
		}
	}
	return &c, nil
}

// FindAll returns all customers from db
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

// DeleteCustomersById implements delete
func (d CustomerRepositoryDb) DeleteCustomerById(id string) (int64, error) {
	deleteCustomerSql := "DELETE FROM customers WHERE customer_id = ?"

	sqlResult, err := d.client.Exec(deleteCustomerSql, id)
	if err != nil {
		return -1, err
	}
	rowsAffected, err := sqlResult.RowsAffected()
	if err != nil {
		return -1, err
	}
	return rowsAffected, nil
}

func (d CustomerRepositoryDb) Create(c Customer) (*Customer, error) {
	createCustomerSql := "INSERT INTO customers (customer_id, name, date_of_birth, city, zipcode, status) VALUES (?,?,?,?,?,?)"

	sqlInsert, err := d.client.Exec(createCustomerSql, c.Id, c.Name, c.DateofBirth, c.City, c.Zipcode, c.Status)
	if err != nil {
		return nil, err // sql insert error. Handle this better!
	}
	newId, err := sqlInsert.LastInsertId()
	if err != nil {
		return nil, err // some error, what can be caught here?
	}
	c.Id = int(newId)
	return &c, nil
}

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
