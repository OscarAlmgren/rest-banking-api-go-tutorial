package app

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/oscaralmgren/rest-banking-api-go-tutorial/service"
)

type SaleHandler struct {
	service service.SaleService
}

// e.GET("/sales/getOne", getOneSale)
func (sh *SaleHandler) getOneSale(c echo.Context) error {
	sale, err := sh.service.FindOne()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, sale)
}
