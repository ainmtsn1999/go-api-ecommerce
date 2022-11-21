package models

import (
	"github.com/ainmtsn1999/go-api-ecommerce/db"
	"gorm.io/gorm"
)

type Merchant struct {
	UserBaseModel
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Street      string `json:"street"`
	CityId      string `json:"city_id"`
	ProvinceId  string `json:"province_id"`
	PictUrl     string `json:"pict_url"`
}

type MerchantTag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Tag struct {
	Id            int `json:"id"`
	MerchantId    int `json:"merchant_id"`
	MerchantTagId int `json:"merchant_tag_id"`
}

type MerchantRequest struct {
	Name        string `json:"name" validate:"required,min=5"`
	PhoneNumber string `json:"phone_number" validate:"required,min=5"`
	Street      string `json:"street" validate:"required,min=5"`
	CityId      string `json:"city_id" validate:"required,min=1"`
	ProvinceId  string `json:"province_id" validate:"required,min=1"`
	PictUrl     string `json:"pict_url" validate:"required,min=1"`
}

func (m *MerchantRequest) ParseToModel() *Merchant {
	return &Merchant{
		Name:        m.Name,
		PhoneNumber: m.PhoneNumber,
		Street:      m.Street,
		CityId:      m.CityId,
		ProvinceId:  m.ProvinceId,
		PictUrl:     m.PictUrl,
	}
}

func CreateMerchant(merchant *Merchant) error {
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(merchant).Error; err != nil {
			return err
		}

		return nil
	})
	return err
}

func UpdateMerchant(merchant *Merchant, authId int) error {
	return db.DB.Model(merchant).Where("auth_id = ?", authId).Updates(merchant).Error
}

func GetMerchantById(id int) (*Merchant, error) {
	var merchant Merchant
	err := db.DB.First(&merchant, "auth_id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &merchant, nil
}

func GetAllMerchant(limit int, page int) (*[]Merchant, error) {
	var merchants []Merchant

	offset := (page - 1) * limit
	err := db.DB.Limit(limit).Offset(offset).Find(&merchants).Error
	if err != nil {
		return nil, err
	}

	return &merchants, nil
}
