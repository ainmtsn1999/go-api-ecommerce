package routers

import (
	"net/http"

	"github.com/ainmtsn1999/go-api-ecommerce/controllers"
	"github.com/ainmtsn1999/go-api-ecommerce/middlewares"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	auth := e.Group("/auth")
	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)

	user := e.Group("/users", middlewares.IsLoggedIn)
	user.GET("", controllers.GetAllUser, middlewares.IsAdmin)
	user.POST("/profile", controllers.CreateUser, middlewares.IsUser)
	user.PUT("/profile", controllers.UpdateUser, middlewares.IsUser)
	user.GET("/profile", controllers.GetUser, middlewares.IsUser)

	merchant := e.Group("/merchants", middlewares.IsLoggedIn)
	merchant.GET("", controllers.GetAllMerchant, middlewares.IsAdmin)
	merchant.POST("/profile", controllers.CreateMerchant, middlewares.IsMerchant)
	merchant.PUT("/profile", controllers.UpdateMerchant, middlewares.IsMerchant)
	merchant.GET("/profile", controllers.GetMerchant, middlewares.IsMerchant)

	product := e.Group("/products")
	product.GET("", controllers.GetAllProduct)
	product.GET("/merchant/:id", controllers.GetAllMerchantProduct)
	product.GET("/id/:id", controllers.GetProduct)
	product.POST("", controllers.CreateProduct, middlewares.IsLoggedIn, middlewares.IsMerchant)
	product.PUT("/id/:id", controllers.UpdateProduct, middlewares.IsLoggedIn, middlewares.IsMerchant)
	product.DELETE("/id/:id", controllers.DeleteProduct, middlewares.IsLoggedIn, middlewares.IsMerchant)

	order := e.Group("/orders", middlewares.IsLoggedIn)
	order.POST("/inquire", controllers.InquireOrder, middlewares.IsUser)
	order.POST("/confirm", controllers.ConfirmOrder, middlewares.IsUser)
	order.PUT("/id/:id/status", controllers.UpdateStatOrder, middlewares.IsMerchant)

	return e
}
