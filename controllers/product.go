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

func CreateProduct(ctx echo.Context) error {
	merchant := ctx.Get("user").(*jwt.Token)
	claims := merchant.Claims.(jwt.MapClaims)
	authId := claims["auth_id"].(float64)

	var req models.ProductRequest
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

	resp := views.CreateProduct(&req, int(authId))
	return WriteJsonResponse(ctx, resp)

}

func UpdateProduct(ctx echo.Context) error {
	paramId := ctx.Param("id")

	productId, err := strconv.Atoi(paramId)
	if err != nil {
		resp := views.ErrorResponse("INVALID_REQUEST", "BAD_REQUEST", http.StatusBadRequest)
		return WriteJsonResponse(ctx, resp)
	}

	var req models.ProductRequest
	err1 := ctx.Bind(&req)
	if err1 != nil {
		resp := views.ErrorResponse("INVALID_REQUEST", "BAD_REQUEST", http.StatusBadRequest)
		return WriteJsonResponse(ctx, resp)

	}

	errString := validators.Validate(req)
	if errString != nil {
		resp := views.ErrorResponse("INPUT_NOT_VALID", errString.Error(), http.StatusBadRequest)
		return WriteJsonResponse(ctx, resp)

	}

	resp := views.UpdateProduct(&req, productId)
	return WriteJsonResponse(ctx, resp)

}

func DeleteProduct(ctx echo.Context) error {
	paramId := ctx.Param("id")

	productId, err := strconv.Atoi(paramId)
	if err != nil {
		resp := views.ErrorResponse("INVALID_REQUEST", "BAD_REQUEST", http.StatusBadRequest)
		return WriteJsonResponse(ctx, resp)
	}

	resp := views.DeleteProduct(productId)
	return WriteJsonResponse(ctx, resp)

}

func GetProduct(ctx echo.Context) error {
	paramId := ctx.Param("id")

	productId, err := strconv.Atoi(paramId)
	if err != nil {
		resp := views.ErrorResponse("INVALID_REQUEST", "BAD_REQUEST", http.StatusBadRequest)
		return WriteJsonResponse(ctx, resp)
	}

	resp := views.GetProductDetail(productId)
	return WriteJsonResponse(ctx, resp)
}

func GetAllMerchantProduct(ctx echo.Context) error {
	paramId := ctx.Param("id")

	merchantId, err := strconv.Atoi(paramId)
	if err != nil {
		resp := views.ErrorResponse("INVALID_REQUEST", "BAD_REQUEST", http.StatusBadRequest)
		return WriteJsonResponse(ctx, resp)
	}

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

	resp := views.GetAllMerchantProduct(merchantId, limit, page)
	return WriteJsonResponse(ctx, resp)
}

func GetAllProduct(ctx echo.Context) error {
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

	resp := views.GetAllProduct(limit, page)
	return WriteJsonResponse(ctx, resp)
}
