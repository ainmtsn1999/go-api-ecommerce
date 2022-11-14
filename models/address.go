package models

type Address struct {
	BaseModel
	UserId     string `json:"user_id"`
	Street     string `json:"street"`
	CityId     string `json:"city_id"`
	ProvinceId string `json:"province_id"`
	AddressTag string `json:"address_tag"`
}
