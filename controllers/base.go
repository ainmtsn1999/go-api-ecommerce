package controllers

import (
	"github.com/ainmtsn1999/go-api-ecommerce/views"
	"github.com/labstack/echo"
)

func WriteJsonResponse(ctx echo.Context, payload *views.Response) error {
	return ctx.JSON(payload.Status, payload)
}
