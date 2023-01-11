package app

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

// DTO (data transformation object)
type Customer struct {
	Name    string `json:"name" xml:"name" form:"name" query:"name"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

func rootGetHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}

func getAllCustomers(c echo.Context) error {
	customers := []Customer{
		{"Oscar", "Nacka", "13131"},
		{"Ida", "Kvarnholmen", "13131"},
		{"Hugo", "Lilla Kvarnholmen", "12345"},
	}

	return c.JSON(http.StatusOK, customers)
}

// e.GET("/customers/:id", getCustomer)
func getCustomer(c echo.Context) error {
	id := c.PathParam("id")
	customer := Customer{id, "Stockholm", "12345"}
	return c.JSON(http.StatusOK, customer)
}

// e.POST("/customers", createCustomer)
func createCustomer(c echo.Context) error {
	return c.String(http.StatusCreated, "Post create customer")
}
