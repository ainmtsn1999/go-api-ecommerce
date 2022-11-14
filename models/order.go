package models

type Order struct {
	BaseModel
	UserId    int    `json:"user_id"`
	TotPrice  int    `json:"total_price"`
	TotWeight int    `json:"total_weight"`
	Status    string `json:"status"`
}

type OrderItem struct {
	OrderId   int `json:"order_id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
