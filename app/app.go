package app

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/oscaralmgren/rest-banking-api-go-tutorial/domain"
	"github.com/oscaralmgren/rest-banking-api-go-tutorial/service"
	"github.com/rs/zerolog/log"
)

func Start() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", rootGetHandler)

	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	// customers endpoint
	e.GET("/customers", ch.getAllCustomers)
	e.POST("/customers", createCustomer)
	e.GET("/customers/:id", ch.getCustomerById)

	if err := e.Start(":3000"); err != http.ErrServerClosed {
		log.Fatal().Str("Error", err.Error()).Msg("Server error on startup")
	}
}
