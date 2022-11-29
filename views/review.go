package views

import (
	"net/http"
	"time"

	"github.com/ainmtsn1999/go-api-ecommerce/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type ReviewDetail struct {
	Id      int         `json:"id"`
	OrderId int         `json:"order_id"`
	User    interface{} `json:"user"`
	Product interface{} `json:"product"`
	Rating  int         `json:"rating"`
	Notes   string      `json:"notes"`
	ImgUrl  string      `json:"img_url"`
}

func CreateReview(req *models.ReviewRequest) *Response {
	review := req.ParseToModel()

	review.CreatedAt = time.Now()

	err := models.CreateReview(review)
	if err != nil {
		return ErrorResponse("CREATE_REVIEW_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponse("CREATE_REVIEW_SUCCESS", nil, http.StatusCreated)
}

func GetAllReview(limit int, page int) *Response {
	reviews, err := models.GetAllReview(limit, page)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrorResponse("GET_ALL_REVIEWS_FAILED", "NOT_FOUND", http.StatusNotFound)
		}
		return ErrorResponse("GET_ALL_REVIEWS_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	pagination := models.Pagination{
		Limit: limit,
		Page:  page,
		Total: len(*reviews),
	}

	resp, err := GetAllReviewResponse(reviews)
	if err != nil {
		return ErrorResponse("GET_ALL_REVIEWS_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponseWithQuery("GET_ALL_REVIEWS_SUCCESS", resp, pagination, http.StatusOK)

}

func GetAllUserReview(id int) *Response {
	orders, err := models.GetAllOrderByUserId(id, 0, 0)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrorResponse("GET_ALL_USER_REVIEWS_DETAIL_FAILED", "NOT_FOUND", http.StatusNotFound)
		}
		return ErrorResponse("GET_ALL_USER_REVIEWS_DETAIL_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	resp, err := GetAllUserReviewResponse(orders)
	if err != nil {
		return ErrorResponse("GET_ALL_USER_REVIEWS_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponse("GET_ALL_USER_REVIEWS_SUCCESS", resp, http.StatusOK)
}

func GetAllUserReviewResponse(orders *[]models.Order) (*[]ReviewDetail, error) {

	var allReviews []ReviewDetail

	for _, order := range *orders {
		reviews, _ := models.GetAllOrderReview(order.Id)
		for _, review := range *reviews {
			resp, _ := ReviewDetailResponse(&review)
			allReviews = append(allReviews, *resp)
		}
	}

	return &allReviews, nil
}

func GetAllOrderReview(id int) *Response {
	reviews, err := models.GetAllOrderReview(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrorResponse("GET_ALL_ORDER_REVIEWS_FAILED", "NOT_FOUND", http.StatusNotFound)
		}
		return ErrorResponse("GET_ALL_ORDER_REVIEWS_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	resp, err := GetAllReviewResponse(reviews)
	if err != nil {
		return ErrorResponse("GET_ALL_ORDER_REVIEWS_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponse("GET_ALL_ORDER_REVIEWS_SUCCESS", resp, http.StatusOK)

}

func GetAllProductReview(id int) *Response {
	reviews, err := models.GetAllProductReview(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrorResponse("GET_ALL_PRODUCT_REVIEWS_FAILED", "NOT_FOUND", http.StatusNotFound)
		}
		return ErrorResponse("GET_ALL_PRODUCT_REVIEWS_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	resp, err := GetAllReviewResponse(reviews)
	if err != nil {
		return ErrorResponse("GET_ALL_PRODUCT_REVIEWS_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponse("GET_ALL_PRODUCT_REVIEWS_SUCCESS", resp, http.StatusOK)

}

func GetAllReviewResponse(reviews *[]models.Review) (*[]ReviewDetail, error) {
	var allReviews []ReviewDetail

	for _, review := range *reviews {
		resp, _ := ReviewDetailResponse(&review)
		allReviews = append(allReviews, *resp)
	}

	return &allReviews, nil
}

func GetReviewDetail(id int) *Response {
	review, err := models.GetReviewById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrorResponse("GET_REVIEW_FAILED", "NOT_FOUND", http.StatusNotFound)
		}
		return ErrorResponse("GET_REVIEW_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	resp, err := ReviewDetailResponse(review)
	if err != nil {
		return ErrorResponse("GET_REVIEW_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponse("GET_REVIEW_SUCCESS", resp, http.StatusOK)
}

func ReviewDetailResponse(review *models.Review) (*ReviewDetail, error) {

	order, err := models.GetOrderById(review.OrderId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	user, err := models.GetUserById(order.UserId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	product, err := models.GetProductById(review.ProductId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	return &ReviewDetail{
		Id:      review.Id,
		OrderId: order.Id,
		User:    echo.Map{"id": user.AuthId, "name": user.Name},
		Product: echo.Map{"id": product.Id, "name": product.Name},
		Rating:  review.Rating,
		Notes:   review.Notes,
		ImgUrl:  review.ImgUrl}, nil
}
