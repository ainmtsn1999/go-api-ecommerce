package views

import (
	"net/http"
	"time"

	"github.com/ainmtsn1999/go-api-ecommerce/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type ProductDetail struct {
	Id       int         `json:"id"`
	Merchant interface{} `json:"merchant"`
	Category string      `json:"category"`
	Name     string      `json:"name"`
	Desc     string      `json:"desc"`
	Price    int         `json:"price"`
	Stock    int         `json:"stock"`
	Weight   int         `json:"weight"`
	ImgUrl   string      `json:"img_url"`
}

func GetAllProduct(limit int, page int) *Response {
	products, err := models.GetAllProduct(limit, page)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrorResponse("GET_ALL_PRODUCTS_FAILED", "NOT_FOUND", http.StatusNotFound)
		}
		return ErrorResponse("GET_ALL_PRODUCTS_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}
	pagination := models.Pagination{
		Limit: limit,
		Page:  page,
		Total: len(*products),
	}

	resp, err := GetAllProductResponse(products)
	if err != nil {
		return ErrorResponse("GET_ALL_PRODUCTS_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponseWithQuery("GET_ALL_PRODUCTS_SUCCESS", resp, pagination, http.StatusOK)

}

func GetAllMerchantProduct(id int, limit int, page int) *Response {
	products, err := models.GetAllMerchantProduct(id, limit, page)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrorResponse("GET_ALL_MERCHANT_PRODUCTS_FAILED", "NOT_FOUND", http.StatusNotFound)
		}
		return ErrorResponse("GET_ALL_MERCHANT_PRODUCTS_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	pagination := models.Pagination{
		Limit: limit,
		Page:  page,
		Total: len(*products),
	}

	resp, err := GetAllProductResponse(products)
	if err != nil {
		return ErrorResponse("GET_ALL_MERCHANT_PRODUCTS_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponseWithQuery("GET_ALL_MERCHANT_PRODUCTS_SUCCESS", resp, pagination, http.StatusOK)

}

func GetAllProductResponse(products *[]models.Product) (*[]ProductDetail, error) {
	var allProducts []ProductDetail

	for _, product := range *products {
		resp, _ := ProductDetailResponse(&product)
		allProducts = append(allProducts, *resp)
	}

	return &allProducts, nil
}

func GetProductDetail(id int) *Response {
	product, err := models.GetProductById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrorResponse("GET_PRODUCT_FAILED", "NOT_FOUND", http.StatusNotFound)
		}
		return ErrorResponse("GET_PRODUCT_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	resp, err := ProductDetailResponse(product)
	if err != nil {
		return ErrorResponse("GET_PRODUCT_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponse("GET_PRODUCT_SUCCESS", resp, http.StatusOK)
}

func ProductDetailResponse(product *models.Product) (*ProductDetail, error) {
	merchant, err := models.GetMerchantById(product.MerchantId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	return &ProductDetail{
		Id:       product.Id,
		Merchant: echo.Map{"id": merchant.AuthId, "name": merchant.Name},
		Category: product.Category,
		Name:     product.Name,
		Desc:     product.Desc,
		Price:    product.Price,
		Stock:    product.Stock,
		Weight:   product.Weight,
		ImgUrl:   product.ImgUrl,
	}, nil
}

func CreateProduct(req *models.ProductRequest, authId int) *Response {
	product := req.ParseToModel()

	product.MerchantId = authId
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	err := models.CreateProduct(product)
	if err != nil {
		return ErrorResponse("CREATE_PRODUCT_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponse("CREATE_PRODUCT_SUCCESS", nil, http.StatusCreated)
}

func UpdateProduct(req *models.ProductRequest, productId int) *Response {
	product := req.ParseToModel()
	product.UpdatedAt = time.Now()

	err := models.UpdateProduct(productId, product)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ErrorResponse("UPDATE_PRODUCT_FAILED", "NOT_FOUND", http.StatusNotModified)
		}
		return ErrorResponse("UPDATE_PRODUCT_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponse("UPDATE_PRODUCT_SUCCESS", nil, http.StatusAccepted)
}

func DeleteProduct(productId int) *Response {

	err := models.DeleteProduct(productId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ErrorResponse("DELETE_PRODUCT_FAILED", "NOT_FOUND", http.StatusNotModified)
		}
		return ErrorResponse("DELETE_PRODUCT_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponse("DELETE_PRODUCT_SUCCESS", nil, http.StatusOK)
}
