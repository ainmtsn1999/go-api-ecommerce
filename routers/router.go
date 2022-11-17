package routers

import (
	"net/http"

	"github.com/ainmtsn1999/go-api-ecommerce/controllers"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/auth/register", controllers.Register)
	e.POST("/auth/login", controllers.Login)

	return e
}
