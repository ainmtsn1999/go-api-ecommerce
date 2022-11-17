package views

import (
	"net/http"

	"github.com/ainmtsn1999/go-api-ecommerce/helper"
	"github.com/ainmtsn1999/go-api-ecommerce/models"
	"gorm.io/gorm"
)

func Register(req *models.Auth) *Response {

	isRegistered, _ := models.FindAccByEmail(req.Email)
	if isRegistered != nil {
		return ErrorResponse("EMAIL_ALREADY_REGISTERED", "UNPROCESSABLE_ENTITY", http.StatusUnprocessableEntity)
	}

	hashedPassword, err := helper.GeneratePassword(req.Password)
	if err != nil {
		return ErrorResponse("FAILED_TO_HASH_PASSWORD", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	req.Password = hashedPassword

	err = models.Register(req)
	if err != nil {
		return ErrorResponse("REGISTER_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponse("REGISTER_SUCCESS", nil, http.StatusCreated)
}

func Login(req *models.AuthLogin) *Response {
	auth, err := models.FindAccByEmail(req.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrorResponse("ACCOUNT_NOT_FOUND", "NOT_FOUND", http.StatusNotFound)
		}

		return ErrorResponse("LOGIN_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	err = helper.ValidatePassword(auth.Password, req.Password)
	if err != nil {
		return ErrorResponse("INVALID_CREDENTIAL", "UNAUTHORIZED", http.StatusUnauthorized)
	}
	tokenPayload := helper.Token{
		AuthId: auth.Id,
		Email:  auth.Email,
	}

	tokenString, err := helper.GenerateToken(&tokenPayload)
	if err != nil {
		return ErrorResponse("LOGIN_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	tokenResponse := map[string]string{
		"token": tokenString,
	}

	return SuccessResponse("LOGIN_SUCCESS", tokenResponse, http.StatusOK)
}
