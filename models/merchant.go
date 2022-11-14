package models

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
