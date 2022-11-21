package models

import (
	"github.com/ainmtsn1999/go-api-ecommerce/db"
	"gorm.io/gorm"
)

type Order struct {
	BaseModel
	UserId      int    `json:"user_id"`
	TotalWeight int    `json:"total_weight"`
	TotalPrice  int    `json:"total_price"`
	Status      string `json:"status"`
}

type Order_Item struct {
	OrderId   int `json:"order_id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type OrderRequest struct {
	Items []ItemRequest `json:"items" validate:"required"`
}

type ItemRequest struct {
	ProductId int `json:"product_id" validate:"required,numeric"`
	Quantity  int `json:"quantity" validate:"required,numeric"`
}

type UpdateStatOrderRequest struct {
	Status string `json:"status" validate:"isValidStatus"`
}

func (t *UpdateStatOrderRequest) ParseToModel() *Order {
	return &Order{
		Status: t.Status,
	}
}

func CreateOrder(order *Order, items *[]Order_Item) error {
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}

		for _, item := range *items {
			item.OrderId = order.Id
			if err := tx.Create(item).Error; err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func UpdateStatOrder(order *Order, orderId int) error {

	return db.DB.Model(order).Where("id = ?", orderId).Updates(order).Error

}
