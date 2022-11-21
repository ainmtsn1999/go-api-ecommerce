package models

import (
	"errors"

	"github.com/ainmtsn1999/go-api-ecommerce/db"
)

type Product struct {
	BaseModel
	MerchantId int    `json:"merchant_id"`
	Category   string `json:"category"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	Weight     int    `json:"weight"`
	ImgUrl     string `json:"img_url"`
}

type ProductRequest struct {
	Name     string `json:"name" validate:"required,min=5"`
	Category string `json:"category" validate:"required,min=5"`
	Desc     string `json:"desc" validate:"required,min=5"`
	Weight   int    `json:"weight" validate:"required,numeric"`
	Price    int    `json:"price" validate:"required,numeric"`
	Stock    int    `json:"stock" validate:"required,numeric"`
	ImgUrl   string `json:"img_url" validate:"required,min=5"`
}

func (p *ProductRequest) ParseToModel() *Product {
	return &Product{
		Name:     p.Name,
		Category: p.Category,
		Desc:     p.Desc,
		Weight:   p.Weight,
		Price:    p.Price,
		Stock:    p.Stock,
		ImgUrl:   p.ImgUrl,
	}
}

func GetAllProduct(limit int, page int) (*[]Product, error) {
	var products []Product

	offset := (page - 1) * limit
	err := db.DB.Limit(limit).Offset(offset).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return &products, nil
}

func GetAllMerchantProduct(id int, limit int, page int) (*[]Product, error) {
	var products []Product

	offset := (page - 1) * limit
	err := db.DB.Where("merchant_id = ?", id).Limit(limit).Offset(offset).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return &products, nil
}

func GetProductById(productId int) (*Product, error) {
	var product Product

	err := db.DB.Where("id = ?", productId).First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func CreateProduct(product *Product) error {
	return db.DB.Create(product).Error
}

func UpdateProduct(productId int, product *Product) error {
	return db.DB.Model(product).Where("id = ?", productId).Updates(product).Error
}

func DeleteProduct(productId int) error {

	if db.DB.Where("id = ?", productId).Delete(Product{}).RowsAffected == 0 {
		return errors.New("NOT_AFFECTED")
	}
	return nil
}
