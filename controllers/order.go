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

func UpdateStatOrderItem(ctx echo.Context) error {
	//get orderId,productId
	paramId1 := ctx.Param("orderId")
	paramId2 := ctx.Param("productId")

	orderId, err := strconv.Atoi(paramId1)
	if err != nil {
		resp := views.ErrorResponse("INVALID_REQUEST", "BAD_REQUEST", http.StatusBadRequest)
		return WriteJsonResponse(ctx, resp)
	}

	productId, err := strconv.Atoi(paramId2)
	if err != nil {
		resp := views.ErrorResponse("INVALID_REQUEST", "BAD_REQUEST", http.StatusBadRequest)
		return WriteJsonResponse(ctx, resp)
	}

	//getting and binding data update
	var req models.UpdateStatOrderItemRequest
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

	resp := views.UpdateStatOrderItem(&req, orderId, productId)

	return WriteJsonResponse(ctx, resp)
}

func GetOrder(ctx echo.Context) error {
	paramId := ctx.Param("id")

	orderId, err := strconv.Atoi(paramId)
	if err != nil {
		resp := views.ErrorResponse("INVALID_REQUEST", "BAD_REQUEST", http.StatusBadRequest)
		return WriteJsonResponse(ctx, resp)
	}

	resp := views.GetOrderDetail(orderId)
	return WriteJsonResponse(ctx, resp)
}

func GetAllUserOrder(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["auth_id"].(float64)

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

	resp := views.GetAllUserOrder(int(userId), limit, page)
	return WriteJsonResponse(ctx, resp)
}

func GetAllOrder(ctx echo.Context) error {
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

	resp := views.GetAllOrder(limit, page)
	return WriteJsonResponse(ctx, resp)
}
