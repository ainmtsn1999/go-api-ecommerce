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
	user.POST("/profile/address", controllers.CreateAddress, middlewares.IsUser)
	user.PUT("/address/id/:id", controllers.UpdateAddress, middlewares.IsUser)
	user.DELETE("/address/id/:id", controllers.DeleteAddress, middlewares.IsUser)
	user.GET("/address/id/:id", controllers.GetAddress, middlewares.IsUser)
	user.GET("/profile/address", controllers.GetAllUserAddresses, middlewares.IsUser)
	user.PUT("/address/id/:id/activate", controllers.UpdateActivateAddress, middlewares.IsUser)

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
	order.PUT("/id/:id/status", controllers.UpdateStatOrder, middlewares.IsUser)
	order.GET("/id/:id", controllers.GetOrder, middlewares.IsUser)
	order.GET("/histories/me", controllers.GetAllUserOrder, middlewares.IsUser)
	order.GET("/histories/list", controllers.GetAllOrder, middlewares.IsAdmin)
	order.PUT("/id/:orderId/product/:productId/status", controllers.UpdateStatOrderItem, middlewares.IsMerchant)

	review := e.Group("/reviews")
	review.GET("", controllers.GetAllReview, middlewares.IsLoggedIn, middlewares.IsAdmin)
	review.GET("/id/:id", controllers.GetReview, middlewares.IsLoggedIn, middlewares.IsUser)
	review.GET("/product/:id", controllers.GetAllProductReview)
	review.GET("/order/:id", controllers.GetAllOrderReview, middlewares.IsLoggedIn, middlewares.IsUser)
	review.GET("/user/:id", controllers.GetAllOrderReview, middlewares.IsLoggedIn, middlewares.IsUser)
	review.POST("", controllers.CreateReview, middlewares.IsLoggedIn, middlewares.IsUser)

	return e
}
