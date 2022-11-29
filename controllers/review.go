package controllers

import (
	"net/http"
	"strconv"

	"github.com/ainmtsn1999/go-api-ecommerce/models"
	"github.com/ainmtsn1999/go-api-ecommerce/validators"
	"github.com/ainmtsn1999/go-api-ecommerce/views"
	"github.com/labstack/echo"
)

func CreateReview(ctx echo.Context) error {
	var req models.ReviewRequest
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

	resp := views.CreateReview(&req)
	return WriteJsonResponse(ctx, resp)

}

func GetReview(ctx echo.Context) error {
	paramId := ctx.Param("id")

	reviewId, err := strconv.Atoi(paramId)
	if err != nil {
		resp := views.ErrorResponse("INVALID_REQUEST", "BAD_REQUEST", http.StatusBadRequest)
		return WriteJsonResponse(ctx, resp)
	}

	resp := views.GetReviewDetail(reviewId)
	return WriteJsonResponse(ctx, resp)
}

func GetAllOrderReview(ctx echo.Context) error {
	paramId := ctx.Param("id")

	orderId, err := strconv.Atoi(paramId)
	if err != nil {
		resp := views.ErrorResponse("INVALID_REQUEST", "BAD_REQUEST", http.StatusBadRequest)
		return WriteJsonResponse(ctx, resp)
	}

	resp := views.GetAllOrderReview(orderId)
	return WriteJsonResponse(ctx, resp)
}

func GetAllProductReview(ctx echo.Context) error {
	paramId := ctx.Param("id")

	productId, err := strconv.Atoi(paramId)
	if err != nil {
		resp := views.ErrorResponse("INVALID_REQUEST", "BAD_REQUEST", http.StatusBadRequest)
		return WriteJsonResponse(ctx, resp)
	}

	resp := views.GetAllProductReview(productId)
	return WriteJsonResponse(ctx, resp)
}

func GetAllUserReview(ctx echo.Context) error {
	paramId := ctx.Param("id")

	userId, err := strconv.Atoi(paramId)
	if err != nil {
		resp := views.ErrorResponse("INVALID_REQUEST", "BAD_REQUEST", http.StatusBadRequest)
		return WriteJsonResponse(ctx, resp)
	}

	resp := views.GetAllProductReview(userId)
	return WriteJsonResponse(ctx, resp)
}

func GetAllReview(ctx echo.Context) error {
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

	resp := views.GetAllReview(limit, page)
	return WriteJsonResponse(ctx, resp)
}
