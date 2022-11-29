package models

import (
	"time"

	"github.com/ainmtsn1999/go-api-ecommerce/db"
)

type Review struct {
	Id        int       `json:"id"`
	OrderId   int       `json:"order_id"`
	ProductId int       `json:"product_id"`
	Rating    int       `json:"rating"`
	Notes     string    `json:"notes"`
	ImgUrl    string    `json:"img_url"`
	CreatedAt time.Time `json:"created_at"`
}

type ReviewRequest struct {
	OrderId   int    `json:"order_id"`
	ProductId int    `json:"product_id"`
	Rating    int    `json:"rating"`
	Notes     string `json:"notes"`
	ImgUrl    string `json:"img_url"`
}

func (r *ReviewRequest) ParseToModel() *Review {
	return &Review{
		OrderId:   r.OrderId,
		ProductId: r.ProductId,
		Rating:    r.Rating,
		Notes:     r.Notes,
		ImgUrl:    r.ImgUrl,
	}
}

func CreateReview(review *Review) error {
	return db.DB.Create(review).Error
}

func GetAllReview(limit int, page int) (*[]Review, error) {
	var reviews []Review

	offset := (page - 1) * limit
	err := db.DB.Limit(limit).Offset(offset).Find(&reviews).Error
	if err != nil {
		return nil, err
	}

	return &reviews, nil
}

func GetAllProductReview(id int) (*[]Review, error) {
	var reviews []Review

	err := db.DB.Where("product_id = ?", id).Find(&reviews).Error
	if err != nil {
		return nil, err
	}

	return &reviews, nil
}

func GetAllOrderReview(id int) (*[]Review, error) {
	var reviews []Review

	err := db.DB.Where("order_id = ?", id).Find(&reviews).Error
	if err != nil {
		return nil, err
	}

	return &reviews, nil
}

func GetReviewById(reviewId int) (*Review, error) {
	var review Review

	err := db.DB.Where("id = ?", reviewId).First(&review).Error
	if err != nil {
		return nil, err
	}

	return &review, nil
}
