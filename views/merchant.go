package views

import (
	"net/http"
	"time"

	"github.com/ainmtsn1999/go-api-ecommerce/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type MerchantDetail struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	PhoneNumber string      `json:"phone_number"`
	PictUrl     string      `json:"pict_url"`
	Address     interface{} `json:"address"`
	Auth        interface{} `json:"auth"`
}

func GetAllMerchant(limit int, page int) *Response {
	merchants, err := models.GetAllMerchant(limit, page)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrorResponse("GET_ALL_MERCHANTS_FAILED", "NOT_FOUND", http.StatusNotFound)
		}
		return ErrorResponse("GET_ALL_MERCHANTS_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}
	pagination := models.Pagination{
		Limit: limit,
		Page:  page,
		Total: len(*merchants),
	}

	resp, err := GetAllMerchantResponse(merchants)
	if err != nil {
		return ErrorResponse("GET_ALL_MERCHANTS_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponseWithQuery("GET_ALL_MERCHANTS_SUCCESS", resp, pagination, http.StatusOK)

}

func GetAllMerchantResponse(merchants *[]models.Merchant) (*[]MerchantDetail, error) {
	var allMerchants []MerchantDetail

	for _, merchant := range *merchants {
		auth, err := models.FindAccById(merchant.AuthId)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, err
			}
			return nil, err
		}

		resp, _ := MerchantDetailResponse(&merchant, auth.Email)
		allMerchants = append(allMerchants, *resp)
	}

	return &allMerchants, nil
}

func GetMerchantDetail(id int, email string) *Response {
	merchant, err := models.GetMerchantDetail(id)
	if err == gorm.ErrRecordNotFound {
		return ErrorResponse("GET_MERCHANT_PROFILE_FAILED", "NOT_FOUND", http.StatusNotFound)
	}

	resp, err := MerchantDetailResponse(merchant, email)
	if err != nil {
		return ErrorResponse("GET_MERCHANT_PROFILE_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponse("GET_MERCHANT_PROFILE_SUCCESS", resp, http.StatusOK)
}

func MerchantDetailResponse(merchant *models.Merchant, email string) (*MerchantDetail, error) {

	return &MerchantDetail{
		Id:          merchant.AuthId,
		Name:        merchant.Name,
		PhoneNumber: merchant.PhoneNumber,
		PictUrl:     merchant.PictUrl,
		Address:     echo.Map{"street": merchant.Street, "city_id": merchant.CityId, "province_id": merchant.ProvinceId},
		Auth:        echo.Map{"email": email},
	}, nil
}

func CreateMerchant(req *models.MerchantRequest, authId int) *Response {
	merchant := req.ParseToModel()

	merchant.AuthId = authId
	merchant.CreatedAt = time.Now()
	merchant.UpdatedAt = time.Now()

	err := models.CreateMerchant(merchant)
	if err != nil {
		return ErrorResponse("CREATE_MERCHANT_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponse("CREATE_MERCHANT_SUCCESS", nil, http.StatusCreated)
}

func UpdateMerchant(req *models.MerchantRequest, authId int) *Response {
	merchant := req.ParseToModel()

	merchant.UpdatedAt = time.Now()

	err := models.UpdateMerchant(merchant, authId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ErrorResponse("UPDATE_MERCHANT_FAILED", "NOT_FOUND", http.StatusNotFound)
		}
		return ErrorResponse("UPDATE_MERCHANT_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponse("UPDATE_MERCHANT_SUCCESS", nil, http.StatusCreated)
}
