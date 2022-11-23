package validators

import (
	"errors"

	"github.com/ainmtsn1999/go-api-ecommerce/enums"
	"github.com/go-playground/validator/v10"
)

// func validator
func Validate(u interface{}) error {
	validate := validator.New()
	validate.RegisterValidation("isValidRole", isValidRole)
	validate.RegisterValidation("isValidOrderStatus", isValidOrderStatus)
	validate.RegisterValidation("isValidItemStatus", isValidItemStatus)
	validate.RegisterValidation("isValidActivateAddress", isValidActivateAddress)
	err := validate.Struct(u)

	if err == nil {
		return nil
	}
	myErr := err.(validator.ValidationErrors)
	errString := ""
	for _, e := range myErr {
		errString += e.Field() + " isn't " + e.Tag()
	}
	return errors.New(errString)
}

func isValidRole(fl validator.FieldLevel) bool {
	v, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	elems := []string{enums.Admin, enums.User, enums.Merchant}

	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func isValidOrderStatus(fl validator.FieldLevel) bool {
	v, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	elems := []string{"ON_PROCESS", "CONFIRMED"}

	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func isValidItemStatus(fl validator.FieldLevel) bool {
	v, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	elems := []string{"WAITING", "PICKUP", "ON_THE_WAY", "ARRIVED"}

	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func isValidActivateAddress(fl validator.FieldLevel) bool {
	v, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	elems := "y"

	return (v == elems)
}
