package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	echo := echo.New()
	echo.GET("/", rootGetHandler)
	echo.Logger.Fatal(echo.Start(":3000"))
}

func rootGetHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}
