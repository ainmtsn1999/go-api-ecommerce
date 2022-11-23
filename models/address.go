package models

import (
	"errors"

	"github.com/ainmtsn1999/go-api-ecommerce/db"
	"gorm.io/gorm"
)

type Address struct {
	BaseModel
	UserId     int    `json:"user_id"`
	Street     string `json:"street"`
	CityId     string `json:"city_id"`
	ProvinceId string `json:"province_id"`
	AddressTag string `json:"address_tag"`
	Activate   string `json:"activate"`
}

type AddressRequest struct {
	Street     string `json:"street" validate:"required,min=5"`
	CityId     string `json:"city_id" validate:"required,min=1"`
	ProvinceId string `json:"province_id" validate:"required,min=1"`
	AddressTag string `json:"address_tag" validate:"required,min=3"`
}

type UpdateActivateAddressRequest struct {
	Activate string `json:"activate" validate:"isValidActivateAddress"`
}

func (m *AddressRequest) ParseToModel() *Address {
	return &Address{
		Street:     m.Street,
		CityId:     m.CityId,
		ProvinceId: m.ProvinceId,
		AddressTag: m.AddressTag,
	}
}

func (t *UpdateActivateAddressRequest) ParseToModel() *Address {
	return &Address{
		Activate: t.Activate,
	}
}

func CreateAddress(address *Address) error {
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(address).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func UpdateAddress(address *Address, addressId int) error {
	return db.DB.Model(address).Where("id = ?", addressId).Updates(address).Error
}

func DeleteAddress(addressId int) error {

	if db.DB.Where("id = ?", addressId).Delete(Address{}).RowsAffected == 0 {
		return errors.New("NOT_AFFECTED")
	}
	return nil
}

func GetAddressById(id int) (*Address, error) {
	var address Address
	err := db.DB.First(&address, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &address, nil
}

func GetAddressByUserId(userId int) (*Address, error) {
	var address Address
	err := db.DB.First(&address, "user_id = ?", userId).Error
	if err != nil {
		return nil, err
	}

	return &address, nil
}

func GetActivateUserAddressByUserId(userId int) (*Address, error) {
	var address Address
	err := db.DB.Where("user_id = ? AND activate = ?", userId, "y").First(&address).Error
	if err != nil {
		return nil, err
	}

	return &address, nil
}

func GetAllUserAddresses(id int, limit int, page int) (*[]Address, error) {
	var addresses []Address

	offset := (page - 1) * limit
	err := db.DB.Where("user_id = ?", id).Limit(limit).Offset(offset).Find(&addresses).Error
	if err != nil {
		return nil, err
	}

	return &addresses, nil
}
