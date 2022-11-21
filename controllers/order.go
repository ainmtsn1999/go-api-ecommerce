package controllers

import (
	"net/http"
	"strconv"

	"github.com/ainmtsn1999/go-api-ecommerce/models"
	"github.com/ainmtsn1999/go-api-ecommerce/validators"
	"github.com/ainmtsn1999/go-api-ecommerce/views"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func InquireOrder(ctx echo.Context) error {
	merchant := ctx.Get("user").(*jwt.Token)
	claims := merchant.Claims.(jwt.MapClaims)
	authId := claims["auth_id"].(float64)

	var req models.OrderRequest
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

	resp := views.InquireOrder(&req, int(authId))
	return WriteJsonResponse(ctx, resp)

}

func ConfirmOrder(ctx echo.Context) error {
	merchant := ctx.Get("user").(*jwt.Token)
	claims := merchant.Claims.(jwt.MapClaims)
	authId := claims["auth_id"].(float64)

	var req models.OrderRequest
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

	resp := views.ConfirmOrder(&req, int(authId))
	return WriteJsonResponse(ctx, resp)

}

func UpdateStatOrder(ctx echo.Context) error {
	//get orderId
	paramId := ctx.Param("id")

	orderId, err := strconv.Atoi(paramId)
	if err != nil {
		resp := views.ErrorResponse("INVALID_REQUEST", "BAD_REQUEST", http.StatusBadRequest)
		return WriteJsonResponse(ctx, resp)
	}

	//getting and binding data update
	var req models.UpdateStatOrderRequest
	err = ctx.Bind(&req)
	if err != nil {
		resp := views.ErrorResponse("INVALID_REQUEST", "BAD_REQUEST", http.StatusBadRequest)
		return WriteJsonResponse(ctx, resp)
	}

	errString := validators.Validate(req)
	if errString != nil {
		resp := views.ErrorResponse("INPUT_NOT_VALID", errString.Error(), http.StatusBadRequest)
		return WriteJsonResponse(ctx, resp)
	}

	resp := views.UpdateStatOrder(&req, orderId)

	return WriteJsonResponse(ctx, resp)
}
