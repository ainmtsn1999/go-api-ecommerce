package views

import (
	"net/http"
	"time"

	"github.com/ainmtsn1999/go-api-ecommerce/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type UserDetail struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	Gender      string      `json:"gender"`
	PhoneNumber string      `json:"phone_number"`
	PictUrl     string      `json:"pict_url"`
	Address     interface{} `json:"address"`
	Auth        interface{} `json:"auth"`
}

func GetAllUser(limit int, page int) *Response {
	users, err := models.GetAllUser(limit, page)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrorResponse("GET_ALL_USERS_FAILED", "NOT_FOUND", http.StatusNotFound)
		}
		return ErrorResponse("GET_ALL_USERS_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}
	pagination := models.Pagination{
		Limit: limit,
		Page:  page,
		Total: len(*users),
	}

	resp, err := GetAllUserResponse(users)
	if err != nil {
		return ErrorResponse("GET_ALL_USERS_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponseWithQuery("GET_ALL_USERS_SUCCESS", resp, pagination, http.StatusOK)

}

func GetAllUserResponse(users *[]models.User) (*[]UserDetail, error) {
	var allUsers []UserDetail

	for _, user := range *users {
		auth, err := models.FindAccById(user.AuthId)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, err
			}
			return nil, err
		}

		resp, _ := UserDetailResponse(&user, auth.Email)
		allUsers = append(allUsers, *resp)
	}

	return &allUsers, nil
}

func GetUserDetail(id int, email string) *Response {
	user, err := models.GetUserById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrorResponse("GET_USER_PROFILE_FAILED", "NOT_FOUND", http.StatusNotFound)
		}
		return ErrorResponse("GET_USER_PROFILE_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	resp, err := UserDetailResponse(user, email)
	if err != nil {
		return ErrorResponse("GET_USER_PROFILE_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponse("GET_USER_PROFILE_SUCCESS", resp, http.StatusOK)
}

func UserDetailResponse(user *models.User, email string) (*UserDetail, error) {

	var address interface{}

	activeAddress, err := models.GetActivateUserAddressByUserId(user.AuthId)

	if err == nil {
		address = echo.Map{"street": activeAddress.Street, "city_id": activeAddress.CityId, "province_id": activeAddress.ProvinceId, "address_tag": activeAddress.AddressTag}
	} else {
		address = echo.Map{"message": "ADDRESS_NOT_FOUND"}
	}

	return &UserDetail{
		Id:          user.AuthId,
		Name:        user.Name,
		Gender:      user.Gender,
		PhoneNumber: user.PhoneNumber,
		PictUrl:     user.PictUrl,
		Address:     address,
		Auth:        echo.Map{"email": email},
	}, nil
}

func CreateUser(req *models.UserRequest, authId int) *Response {
	user := req.ParseToModel()

	user.AuthId = authId
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err := models.CreateUser(user)
	if err != nil {
		return ErrorResponse("CREATE_USER_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponse("CREATE_USER_SUCCESS", nil, http.StatusCreated)
}

func UpdateUser(req *models.UserRequest, authId int) *Response {
	user := req.ParseToModel()

	user.UpdatedAt = time.Now()

	err := models.UpdateUser(user, authId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ErrorResponse("UPDATE_USER_FAILED", "NOT_FOUND", http.StatusNotModified)
		}
		return ErrorResponse("UPDATE_USER_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponse("UPDATE_USER_SUCCESS", nil, http.StatusAccepted)
}
