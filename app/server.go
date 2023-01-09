package app

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
)

func Start() {
	e := echo.New()
	e.GET("/", rootGetHandler)

	e.GET("/customers", getAllCustomers)

	if err := e.Start(":3000"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
