package views

import (
	"net/http"
	"time"

	"github.com/ainmtsn1999/go-api-ecommerce/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type InquireOrderDetail struct {
	User      UserOrder     `json:"user"`
	Order     []OrderDetail `json:"order"`
	TotWeight int           `json:"total_weight"`
	TotPrice  int           `json:"total_price"`
}

type UserOrder struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type OrderDetail struct {
	Item     ItemDetail `json:"item"`
	Quantity int        `json:"quantity"`
}

type ItemDetail struct {
	Id       int         `json:"id"`
	Merchant interface{} `json:"merchant"`
	Name     string      `json:"name"`
	Weight   int         `json:"weight"`
	Price    int         `json:"price"`
	ImgUrl   string      `json:"img_url"`
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

	var allOrders []OrderDetail
	var totWeight = 0
	var totPrice = 0

	user, err := models.GetUserById(userId)
	if err == gorm.ErrRecordNotFound {
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

		allOrders = append(allOrders, OrderDetail{
			Item:     *dtl,
			Quantity: item.Quantity,
		})
	}

	return &InquireOrderDetail{
		User:      *resp,
		Order:     allOrders,
		TotWeight: totWeight,
		TotPrice:  totPrice,
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

		totWeight += item.Quantity * product.Weight
		totPrice += item.Quantity * product.Price

		allOrderItems = append(allOrderItems, models.Order_Item{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		})
	}

	var order = models.Order{
		UserId:      userId,
		TotalWeight: totWeight,
		TotalPrice:  totPrice,
		Status:      "WAITING",
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

	err := models.UpdateStatOrder(status, orderId)
	if err != nil {
		return ErrorResponse("UPDATE_STATUS_TRANSACTION_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponse("UPDATE_STATUS_TRANSACTION_SUCCESS", nil, http.StatusAccepted)
}
