package controllers

import (
	"net/http"

	"github.com/ainmtsn1999/go-api-ecommerce/models"
	"github.com/ainmtsn1999/go-api-ecommerce/validators"
	"github.com/ainmtsn1999/go-api-ecommerce/views"
	"github.com/labstack/echo"
)

func Register(ctx echo.Context) error {
	var req models.Auth
	err := ctx.Bind(&req)
	if err != nil {
		resp := views.ErrorResponse("INVALID_REQUEST", "BAD_REQUEST", http.StatusBadRequest)
		return WriteJsonResponse(ctx, resp)

	}

	errString := validators.Validate(req)
	if errString != nil {
		resp := views.ErrorResponse("INPUT_NOT_VALID", errString.Error(), http.StatusBadRequest)
		return WriteJsonResponse(ctx, resp)

	}

	resp := views.Register(&req)
	return WriteJsonResponse(ctx, resp)
}

func Login(ctx echo.Context) error {
	var req models.AuthLogin
	err := ctx.Bind(&req)
	if err != nil {
		resp := views.ErrorResponse("INVALID_REQUEST", "BAD_REQUEST", http.StatusBadRequest)
		return WriteJsonResponse(ctx, resp)
	}

	errString := validators.Validate(req)
	if errString != nil {
		resp := views.ErrorResponse("INPUT_NOT_VALID", errString.Error(), http.StatusBadRequest)
		return WriteJsonResponse(ctx, resp)

	}

	resp := views.Login(&req)
	return WriteJsonResponse(ctx, resp)
}
