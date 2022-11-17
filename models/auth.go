package models

import (
	"github.com/ainmtsn1999/go-api-ecommerce/db"
	"gorm.io/gorm"
)

type Auth struct {
	Id       int    `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Role     string `json:"role" validate:"required,isValidRole"`
}

type AuthLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// func model
func Register(acc *Auth) error {
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(acc).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}
func FindAccByEmail(email string) (*Auth, error) {
	var auth Auth
	err := db.DB.Where("email=?", email).First(&auth).Error
	if err != nil {
		return nil, err
	}

	return &auth, nil
}
