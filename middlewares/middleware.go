package middlewares

import (
	"github.com/ainmtsn1999/go-api-ecommerce/config"
	"github.com/ainmtsn1999/go-api-ecommerce/enums"
	"github.com/ainmtsn1999/go-api-ecommerce/models"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var secret = config.GetEnvVariable("JWT_SECRET")

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(secret),
})

func IsUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		authId := claims["auth_id"].(float64)

		auth, err := models.FindAccById(int(authId))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return echo.ErrNotFound
			}
			return echo.ErrInternalServerError
		}

		if auth.Role != enums.User {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}

func IsMerchant(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		authId := claims["auth_id"].(float64)

		auth, err := models.FindAccById(int(authId))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return echo.ErrNotFound
			}
			return echo.ErrInternalServerError
		}

		if auth.Role != enums.Merchant {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		authId := claims["auth_id"].(float64)

		auth, err := models.FindAccById(int(authId))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return echo.ErrNotFound
			}
			return echo.ErrInternalServerError
		}

		if auth.Role != enums.Admin {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}

// func IsExpired(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		user := c.Request().Header.Get("Authorization")
// 		fmt.Print(user)
// 		return next(c)
// 	}
// }
