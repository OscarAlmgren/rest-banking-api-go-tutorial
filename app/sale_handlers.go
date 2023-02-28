package app

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v5"
	domain "github.com/oscaralmgren/rest-banking-api-go-tutorial/domain/sale"
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

// e.GET("/sales", getAllSales)
func (sh *SaleHandler) getAllSales(c echo.Context) error {
	sales, err := sh.service.FindAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, sales)
}

// e.POST("/sales", sh.create)
func (sh *SaleHandler) create(c echo.Context) error {
	sale := domain.Sale{
		Item:     "abc",
		Price:    10,
		Quantity: 3,
		Date:     time.Now(),
	}
	inserted, err := sh.service.Create(sale)
	if err != nil {
		return err
	}
	return c.String(http.StatusCreated, "Created sale id:"+inserted)
}

func (sh *SaleHandler) delete(c echo.Context) error {
	id := c.PathParam("id")
	deleted, err := sh.service.Delete(id)
	if err != nil {
		return err
	}
	return c.String(http.StatusAccepted, "Items deleted:"+strconv.Itoa(int(deleted)))

}
