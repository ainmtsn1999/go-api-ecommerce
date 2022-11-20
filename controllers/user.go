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

func CreateUser(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authId := claims["auth_id"].(float64)

	var req models.UserRequest
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

	resp := views.CreateUser(&req, int(authId))
	return WriteJsonResponse(ctx, resp)

}

func UpdateUser(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authId := claims["auth_id"].(float64)

	var req models.UserRequest
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

	resp := views.UpdateUser(&req, int(authId))
	return WriteJsonResponse(ctx, resp)

}

func GetUser(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authId := claims["auth_id"].(float64)
	email := claims["email"].(string)

	resp := views.GetUserDetail(int(authId), email)
	return WriteJsonResponse(ctx, resp)
}

func GetAllUser(ctx echo.Context) error {
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

	resp := views.GetAllUser(limit, page)
	return WriteJsonResponse(ctx, resp)
}
