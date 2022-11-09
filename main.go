package main

import (
	"fmt"
	"net/http"

	"github.com/ainmtsn1999/go-api-ecommerce/config"
	"github.com/labstack/echo"
)

func main() {

	// db, err := db.ConnectDB()
	// if err != nil {
	// 	panic(err)
	// }

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	port := fmt.Sprintf(":%s", config.GetEnvVariable("APP_PORT"))
	e.Logger.Fatal(e.Start(port))
}
