package app

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	customer "github.com/oscaralmgren/rest-banking-api-go-tutorial/domain/customer"
	sale "github.com/oscaralmgren/rest-banking-api-go-tutorial/domain/sale"
	"github.com/oscaralmgren/rest-banking-api-go-tutorial/service"
	"github.com/rs/zerolog/log"
)

func Start() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", rootGetHandler)

	ch := CustomerHandler{service.NewCustomerService(customer.NewCustomerRepositoryDb())}
	// customers endpoint
	e.GET("/customers", ch.getAllCustomers)
	e.POST("/customers", ch.create)
	e.GET("/customers/:id", ch.getCustomerById)
	e.DELETE("customers/:id", ch.deleteCustomerById)

	sh := SaleHandler{service.NewSaleService(sale.NewSaleRepositoryDb())}

	e.GET("/sales/getOneSale", sh.getOneSale)
	e.GET("/sales", sh.getAllSales)

	if err := e.Start(":3000"); err != http.ErrServerClosed {
		log.Fatal().Str("Error", err.Error()).Msg("Server error on startup")
	}
}
