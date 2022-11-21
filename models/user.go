package models

import (
	"github.com/ainmtsn1999/go-api-ecommerce/db"
	"gorm.io/gorm"
)

type User struct {
	UserBaseModel
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phone_number"`
	PictUrl     string `json:"pict_url"`
}

type UserRequest struct {
	Name        string `json:"name" validate:"required,min=5"`
	Gender      string `json:"gender" validate:"required,min=1"`
	PhoneNumber string `json:"phone_number" validate:"required,min=5"`
	PictUrl     string `json:"pict_url" validate:"required,min=1"`
}

func (u *UserRequest) ParseToModel() *User {
	return &User{
		Name:        u.Name,
		Gender:      u.Gender,
		PhoneNumber: u.PhoneNumber,
		PictUrl:     u.PictUrl,
	}
}

func CreateUser(user *User) error {
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}

func UpdateUser(user *User, authId int) error {
	return db.DB.Model(user).Where("auth_id = ?", authId).Updates(user).Error
}

func GetUserById(id int) (*User, error) {
	var user User
	err := db.DB.First(&user, "auth_id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetAllUser(limit int, page int) (*[]User, error) {
	var users []User

	offset := (page - 1) * limit
	err := db.DB.Limit(limit).Offset(offset).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return &users, nil
}
