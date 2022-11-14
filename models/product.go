package models

type Product struct {
	BaseModel
	MerchantId string `json:"merchant_id"`
	Category   string `json:"category"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	Weight     int    `json:"weight"`
	ImgUrl     string `json:"img_url"`
}
