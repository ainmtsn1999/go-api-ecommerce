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

func GetAllUserAddresses(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authId := claims["auth_id"].(float64)

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

	resp := views.GetAllUserAddresses(int(authId), limit, page)
	return WriteJsonResponse(ctx, resp)
}

func GetAddress(ctx echo.Context) error {
	paramId := ctx.Param("id")

	addressId, err := strconv.Atoi(paramId)
	if err != nil {
		resp := views.ErrorResponse("INVALID_REQUEST", "BAD_REQUEST", http.StatusBadRequest)
		return WriteJsonResponse(ctx, resp)
	}

	resp := views.GetAddressDetail(addressId)
	return WriteJsonResponse(ctx, resp)
}

func CreateAddress(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authId := claims["auth_id"].(float64)

	var req models.AddressRequest
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

	resp := views.CreateAddress(&req, int(authId))
	return WriteJsonResponse(ctx, resp)

}

func UpdateAddress(ctx echo.Context) error {
	paramId := ctx.Param("id")

	addressId, err := strconv.Atoi(paramId)
	if err != nil {
		resp := views.ErrorResponse("INVALID_REQUEST", "BAD_REQUEST", http.StatusBadRequest)
		return WriteJsonResponse(ctx, resp)
	}

	var req models.AddressRequest
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

	resp := views.UpdateAddress(&req, addressId)
	return WriteJsonResponse(ctx, resp)

}

func DeleteAddress(ctx echo.Context) error {
	paramId := ctx.Param("id")

	addressId, err := strconv.Atoi(paramId)
	if err != nil {
		resp := views.ErrorResponse("INVALID_REQUEST", "BAD_REQUEST", http.StatusBadRequest)
		return WriteJsonResponse(ctx, resp)
	}

	resp := views.DeleteAddress(addressId)
	return WriteJsonResponse(ctx, resp)

}

func UpdateActivateAddress(ctx echo.Context) error {
	//get orderId
	paramId := ctx.Param("id")

	addressId, err := strconv.Atoi(paramId)
	if err != nil {
		resp := views.ErrorResponse("INVALID_REQUEST", "BAD_REQUEST", http.StatusBadRequest)
		return WriteJsonResponse(ctx, resp)
	}

	//getting and binding data update
	var req models.UpdateActivateAddressRequest
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

	resp := views.UpdateActivateAddress(&req, addressId)

	return WriteJsonResponse(ctx, resp)
}
