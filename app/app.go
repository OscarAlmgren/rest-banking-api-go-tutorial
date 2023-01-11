package app

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func Start() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", rootGetHandler)

	// customers endpoint
	e.GET("/customers", getAllCustomers)
	e.POST("/customers", createCustomer)
	e.GET("/customers/:id", getCustomer)

	if err := e.Start(":3000"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
