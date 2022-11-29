package views

import (
	"net/http"
	"time"

	"github.com/ainmtsn1999/go-api-ecommerce/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type InquireOrderDetail struct {
	User        UserOrder         `json:"user"`
	Order       []OrderItemDetail `json:"order"`
	TotalWeight int               `json:"total_weight"`
	TotalPrice  int               `json:"total_price"`
}

type UserOrder struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type OrderItemDetail struct {
	Item     ItemDetail `json:"item"`
	Quantity int        `json:"quantity"`
	Notes    string     `json:"notes"`
	Status   string     `json:"status,omitempty"`
}

type ItemDetail struct {
	Id       int         `json:"id"`
	Merchant interface{} `json:"merchant"`
	Name     string      `json:"name"`
	Weight   int         `json:"weight"`
	Price    int         `json:"price"`
	ImgUrl   string      `json:"img_url"`
}

type OrderDetail struct {
	Id          int               `json:"id"`
	User        UserOrder         `json:"user"`
	Order       []OrderItemDetail `json:"order"`
	TotalWeight int               `json:"total_weight"`
	TotalPrice  int               `json:"total_price"`
	Status      string            `json:"status"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	DeletedAt   time.Time         `json:"deleted_at"`
}

func GetAllOrder(limit int, page int) *Response {
	orders, err := models.GetAllOrder(limit, page)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrorResponse("GET_ALL_ORDERS_DETAIL_FAILED", "NOT_FOUND", http.StatusNotFound)
		}
		return ErrorResponse("GET_ALL_ORDERS_DETAIL_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}
	pagination := models.Pagination{
		Limit: limit,
		Page:  page,
		Total: len(*orders),
	}

	resp, err := GetAllOrderResponse(orders)
	if err != nil {
		return ErrorResponse("GET_ALL_ORDERS_DETAIL_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponseWithQuery("GET_ALL_ORDERS_DETAIL_SUCCESS", resp, pagination, http.StatusOK)

}

func GetAllUserOrder(id int, limit int, page int) *Response {
	orders, err := models.GetAllOrderByUserId(id, limit, page)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrorResponse("GET_ALL_USER_ORDERS_DETAIL_FAILED", "NOT_FOUND", http.StatusNotFound)
		}
		return ErrorResponse("GET_ALL_USER_ORDERS_DETAIL_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}
	pagination := models.Pagination{
		Limit: limit,
		Page:  page,
		Total: len(*orders),
	}

	resp, err := GetAllOrderResponse(orders)
	if err != nil {
		return ErrorResponse("GET_ALL_USER_ORDERS_DETAIL_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponseWithQuery("GET_ALL_USER_ORDERS_DETAIL_SUCCESS", resp, pagination, http.StatusOK)

}

func GetAllOrderResponse(orders *[]models.Order) (*[]OrderDetail, error) {
	var allOrders []OrderDetail

	for _, order := range *orders {
		resp, _ := OrderDetailResponse(&order)
		allOrders = append(allOrders, *resp)
	}

	return &allOrders, nil
}

func GetOrderDetail(id int) *Response {
	order, err := models.GetOrderById(id)
	if err == gorm.ErrRecordNotFound {
		return ErrorResponse("GET_ORDER_DETAIL_FAILED", "NOT_FOUND", http.StatusNotFound)
	}

	resp, err := OrderDetailResponse(order)
	if err != nil {
		return ErrorResponse("GET_ORDER_DETAIL_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponse("GET_ORDER_DETAIL_SUCCESS", resp, http.StatusOK)
}

func OrderDetailResponse(order *models.Order) (*OrderDetail, error) {

	var allOrders []OrderItemDetail

	user, err := models.GetUserById(order.UserId)
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	resp, err := UserOrderResponse(user)
	if err != nil {
		return nil, err
	}

	itemdtl, err := models.GetAllItemByOrderId(order.Id)
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	for _, item := range *itemdtl {
		product, err := models.GetProductById(item.ProductId)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, err
			}
			return nil, err
		}
		dtl, err := ItemDetailResponse(product)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, err
			}
			return nil, err
		}

		allOrders = append(allOrders, OrderItemDetail{
			Item:     *dtl,
			Quantity: item.Quantity,
			Notes:    item.Notes,
			Status:   item.Status,
		})
	}

	return &OrderDetail{
		Id:          order.Id,
		User:        *resp,
		Order:       allOrders,
		TotalWeight: order.TotalWeight,
		TotalPrice:  order.TotalPrice,
		Status:      order.Status,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
		DeletedAt:   order.DeletedAt,
	}, nil
}

func InquireOrder(req *models.OrderRequest, userId int) *Response {

	for _, item := range req.Items {
		product, err := models.GetProductById(item.ProductId)
		if err != nil {
			return ErrorResponse("GET_DETAIL_PRODUCT_FAILED", "NOT_FOUND", http.StatusNotFound)
		}

		if item.Quantity > product.Stock {
			return ErrorResponse("INQUIRY_ORDER_FAILED", "UNPROCESSABLE_ENTITY", http.StatusUnprocessableEntity)
		}
	}

	resp, err := InquireOrderDetailResponse(req, userId)
	if err != nil {
		return ErrorResponse("GET_DETAIL_INQUIRE_ORDER_FAILED", "NOT_FOUND", http.StatusNotFound)
	}

	return SuccessResponse("INQUIRY_ORDER_SUCCESS", resp, http.StatusCreated)
}

func InquireOrderDetailResponse(req *models.OrderRequest, userId int) (*InquireOrderDetail, error) {

	var allOrders []OrderItemDetail
	var totWeight = 0
	var totPrice = 0

	user, err := models.GetUserById(userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	resp, err := UserOrderResponse(user)
	if err != nil {
		return nil, err
	}

	for _, item := range req.Items {
		product, err := models.GetProductById(item.ProductId)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, err
			}
			return nil, err
		}

		totWeight += product.Weight * item.Quantity
		totPrice += product.Price * item.Quantity

		dtl, err := ItemDetailResponse(product)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, err
			}
			return nil, err
		}

		allOrders = append(allOrders, OrderItemDetail{
			Item:     *dtl,
			Quantity: item.Quantity,
			Notes:    item.Notes,
		})
	}

	return &InquireOrderDetail{
		User:        *resp,
		Order:       allOrders,
		TotalWeight: totWeight,
		TotalPrice:  totPrice,
	}, nil
}

func ItemDetailResponse(product *models.Product) (*ItemDetail, error) {
	merchant, err := models.GetMerchantById(product.MerchantId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	return &ItemDetail{
		Id:       product.Id,
		Merchant: echo.Map{"id": merchant.AuthId, "name": merchant.Name},
		Name:     product.Name,
		Weight:   product.Weight,
		Price:    product.Price,
		ImgUrl:   product.ImgUrl,
	}, nil
}

func UserOrderResponse(user *models.User) (*UserOrder, error) {
	auth, err := models.FindAccById(user.AuthId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return &UserOrder{
		Id:          user.AuthId,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Email:       auth.Email,
	}, nil
}

func ConfirmOrder(req *models.OrderRequest, userId int) *Response {

	var allOrderItems []models.Order_Item
	var totWeight = 0
	var totPrice = 0
	for _, item := range req.Items {
		product, err := models.GetProductById(item.ProductId)
		if err != nil {
			return ErrorResponse("GET_DETAIL_PRODUCT_FAILED", "NOT_FOUND", http.StatusNotFound)
		}

		if item.Quantity > product.Stock {
			return ErrorResponse("INQUIRY_ORDER_FAILED", "UNPROCESSABLE_ENTITY", http.StatusUnprocessableEntity)
		}

		newStock := product.Stock - item.Quantity

		updateStock := models.Product{
			Stock: newStock,
		}
		err = models.UpdateProduct(product.Id, &updateStock)
		if err != nil {
			return ErrorResponse("UPDATE_STOCK_PRODUCT_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		}

		totWeight += item.Quantity * product.Weight
		totPrice += item.Quantity * product.Price

		allOrderItems = append(allOrderItems, models.Order_Item{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
			Notes:     item.Notes,
			Status:    "WAITING",
		})
	}

	var order = models.Order{
		UserId:      userId,
		TotalWeight: totWeight,
		TotalPrice:  totPrice,
		Status:      "ON_PROCESS",
		BaseModel: models.BaseModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	err := models.CreateOrder(&order, &allOrderItems)
	if err != nil {
		return ErrorResponse("CONFIRM_ORDER_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponse("CONFIRM_ORDER_SUCCESS", nil, http.StatusCreated)
}

func UpdateStatOrder(req *models.UpdateStatOrderRequest, orderId int) *Response {
	status := req.ParseToModel()

	err := models.UpdateOrder(status, orderId)
	if err != nil {
		return ErrorResponse("UPDATE_STATUS_ORDER_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponse("UPDATE_STATUS_ORDER_SUCCESS", nil, http.StatusAccepted)
}

func UpdateStatOrderItem(req *models.UpdateStatOrderItemRequest, orderId int, productId int) *Response {
	status := req.ParseToModel()

	err := models.UpdateOrderItem(status, orderId, productId)
	if err != nil {
		return ErrorResponse("UPDATE_STATUS_ORDER_ITEM_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponse("UPDATE_STATUS_ORDER_ITEM_SUCCESS", nil, http.StatusAccepted)
}
