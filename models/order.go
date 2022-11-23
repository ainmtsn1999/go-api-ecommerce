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
	OrderId   int    `json:"order_id"`
	ProductId int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Notes     string `json:"notes"`
	Status    string `json:"status"`
}

type OrderRequest struct {
	Items []ItemRequest `json:"items" validate:"required"`
}

type ItemRequest struct {
	ProductId int    `json:"product_id" validate:"required,numeric"`
	Quantity  int    `json:"quantity" validate:"required,numeric"`
	Notes     string `json:"notes"`
}

type UpdateStatOrderRequest struct {
	Status string `json:"status" validate:"isValidOrderStatus"`
}

type UpdateStatOrderItemRequest struct {
	Status string `json:"status" validate:"isValidItemStatus"`
}

func (t *UpdateStatOrderRequest) ParseToModel() *Order {
	return &Order{
		Status: t.Status,
	}
}

func (t *UpdateStatOrderItemRequest) ParseToModel() *Order_Item {
	return &Order_Item{
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

func UpdateOrder(order *Order, orderId int) error {

	return db.DB.Model(order).Where("id = ?", orderId).Updates(order).Error

}

func UpdateOrderItem(item *Order_Item, orderId int, productId int) error {

	return db.DB.Model(item).Where("order_id = ? AND product_id = ?", orderId, productId).Updates(item).Error

}

func GetAllItemByOrderId(id int) (*[]Order_Item, error) {
	var items []Order_Item

	err := db.DB.Where("order_id = ?", id).Find(&items).Error
	if err != nil {
		return nil, err
	}

	return &items, nil
}

func GetOrderById(id int) (*Order, error) {
	var order Order

	err := db.DB.Where("id = ?", id).First(&order).Error
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func GetAllOrderByUserId(id int, limit int, page int) (*[]Order, error) {
	var orders []Order

	offset := (page - 1) * limit
	err := db.DB.Where("user_id = ?", id).Limit(limit).Offset(offset).Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return &orders, nil
}
