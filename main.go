package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/login", login)
	e.Logger.Fatal(e.Start(":1323"))
}

func login(c echo.Context) error {

	return c.JSON(http.StatusOK, "")
}
