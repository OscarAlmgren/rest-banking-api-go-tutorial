package app

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
	"github.com/oscaralmgren/rest-banking-api-go-tutorial/service"
)

// DTO (data transformation object)

type CustomerHandler struct {
	service service.CustomerService
}

func rootGetHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}

func (ch *CustomerHandler) getAllCustomers(c echo.Context) error {

	customers, err := ch.service.GetAllCustomers()
	if err != nil {
		return err
	}
	// customers := []Customer{
	// 	{"Oscar", "Nacka", "13131"},
	// 	{"Ida", "Kvarnholmen", "13131"},
	// 	{"Hugo", "Lilla Kvarnholmen", "12345"},
	// }

	return c.JSON(http.StatusOK, customers)
}

// e.GET("/customers/:id", getCustomer)
func (ch *CustomerHandler) getCustomerById(c echo.Context) error {
	id := c.PathParam("id")
	customer, err := ch.service.GetCustomerById(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, customer)
}

// e.DELETE("/customers/:id", deleteCustomerById)
func (ch *CustomerHandler) deleteCustomerById(c echo.Context) error {
	id := c.PathParam("id")
	rowsAffected, err := ch.service.DeleteCustomerById(id)
	if err != nil {
		return err
	}
	return c.String(http.StatusAccepted, "Rows affected:"+strconv.Itoa(int(rowsAffected)))

}

// e.POST("/customers", createCustomer)
func (ch *CustomerHandler) create(c echo.Context) error {
	return c.String(http.StatusCreated, "Post create customer")
}
