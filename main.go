package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

func main() {
	e := echo.New()
	e.GET("/", rootGetHandler)

	e.GET("/customers", getAllCustomers)

	e.Logger.Fatal(e.Start(":3000"))
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
