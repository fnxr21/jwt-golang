package main

import (
	"fmt"
	jwtToken "jwt/jwt"
	"jwt/middleware"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", testcreate)
	e.GET("/login", middleware.Auth(login))
	e.Logger.Fatal(e.Start(":1323"))
}

type claims struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func login(c echo.Context) error {

	userClaims := c.Get("userLogin").(jwt.MapClaims)

	// Type assertion to ensure expected claim type
	name, ok := userClaims["name"].(string)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid claim type for 'name'")
	}

	fmt.Println("User Name:", name)

	return c.JSON(http.StatusOK, "token")
}
func testcreate(c echo.Context) error {

	// generate token
	claims := jwt.MapClaims{}
	claims["id"] = 1
	claims["name"] = "fandi"
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() //2 hours expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)

		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	return c.JSON(http.StatusOK, token)
}
