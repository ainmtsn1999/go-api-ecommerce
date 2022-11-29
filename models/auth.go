package models

import (
	"time"

	"github.com/ainmtsn1999/go-api-ecommerce/db"
	"gorm.io/gorm"
)

type Auth struct {
	Id       int       `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
	LoginAt  time.Time `json:"login_at"`
}

type AuthRequest struct {
	Id       int    `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Role     string `json:"role" validate:"required,isValidRole"`
}

type AuthLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (r *AuthRequest) ParseToModel() *Auth {
	return &Auth{
		Id:       r.Id,
		Email:    r.Email,
		Password: r.Password,
		Role:     r.Role,
	}
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

func FindAccById(id int) (*Auth, error) {
	var auth Auth
	err := db.DB.Where("id=?", id).First(&auth).Error
	if err != nil {
		return nil, err
	}

	return &auth, nil
}

func UpdateAuth(authId int, auth *Auth) error {
	return db.DB.Model(auth).Where("id = ?", authId).Updates(auth).Error
}
