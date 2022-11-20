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

func CreateMerchant(ctx echo.Context) error {
	merchant := ctx.Get("user").(*jwt.Token)
	claims := merchant.Claims.(jwt.MapClaims)
	authId := claims["auth_id"].(float64)

	var req models.MerchantRequest
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

	resp := views.CreateMerchant(&req, int(authId))
	return WriteJsonResponse(ctx, resp)

}

func UpdateMerchant(ctx echo.Context) error {
	merchant := ctx.Get("user").(*jwt.Token)
	claims := merchant.Claims.(jwt.MapClaims)
	authId := claims["auth_id"].(float64)

	var req models.MerchantRequest
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

	resp := views.UpdateMerchant(&req, int(authId))
	return WriteJsonResponse(ctx, resp)

}

func GetMerchant(ctx echo.Context) error {
	merchant := ctx.Get("user").(*jwt.Token)
	claims := merchant.Claims.(jwt.MapClaims)
	authId := claims["auth_id"].(float64)
	email := claims["email"].(string)

	resp := views.GetMerchantDetail(int(authId), email)
	return WriteJsonResponse(ctx, resp)
}

func GetAllMerchant(ctx echo.Context) error {
	limitStr := ctx.QueryParam("limit")
	pageStr := ctx.QueryParam("page")

	var limit int
	var page int

	if limitStr == "" {
		limit = 25
	} else {
		limit, _ = strconv.Atoi(limitStr)
	}

	if pageStr == "" {
		page = 1
	} else {
		page, _ = strconv.Atoi(pageStr)
	}

	resp := views.GetAllMerchant(limit, page)
	return WriteJsonResponse(ctx, resp)
}
