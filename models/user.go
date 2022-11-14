package models

type User struct {
	UserBaseModel
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phone_number"`
	PictUrl     string `json:"pict_url"`
}
