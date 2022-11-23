package views

import (
	"net/http"
	"time"

	"github.com/ainmtsn1999/go-api-ecommerce/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type AddressDetail struct {
	Id         int         `json:"id"`
	User       interface{} `json:"user"`
	Address    interface{} `json:"address"`
	AddressTag string      `json:"address_tag"`
	Activate   string      `json:"activate"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	DeletedAt  time.Time   `json:"deleted_at"`
}

func GetAllUserAddresses(id int, limit int, page int) *Response {
	addresses, err := models.GetAllUserAddresses(id, limit, page)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrorResponse("GET_ALL_USER_ADDRESSES_FAILED", "NOT_FOUND", http.StatusNotFound)
		}
		return ErrorResponse("GET_ALL_USER_ADDRESSES_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	pagination := models.Pagination{
		Limit: limit,
		Page:  page,
		Total: len(*addresses),
	}

	resp, err := GetAllAddressesResponse(addresses)
	if err != nil {
		return ErrorResponse("GET_ALL_USER_ADDRESSES_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponseWithQuery("GET_ALL_USER_ADDRESSES_SUCCESS", resp, pagination, http.StatusOK)

}

func GetAllAddressesResponse(addresses *[]models.Address) (*[]AddressDetail, error) {
	var allAddresses []AddressDetail

	for _, address := range *addresses {
		resp, _ := AddressDetailResponse(&address)
		allAddresses = append(allAddresses, *resp)
	}

	return &allAddresses, nil
}

func GetAddressDetail(id int) *Response {
	address, err := models.GetAddressById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrorResponse("GET_ADDRESS_DETAIL_FAILED", "NOT_FOUND", http.StatusNotFound)
		}
		return ErrorResponse("GET_ADDRESS_DETAIL_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	resp, err := AddressDetailResponse(address)
	if err != nil {
		return ErrorResponse("GET_ADDRESS_DETAIL_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponse("GET_ADDRESS_DETAIL_SUCCESS", resp, http.StatusOK)
}

func AddressDetailResponse(address *models.Address) (*AddressDetail, error) {
	user, err := models.GetUserById(address.UserId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return &AddressDetail{
		Id:         address.Id,
		User:       echo.Map{"id": user.AuthId, "name": user.Name, "phone_number": user.PhoneNumber},
		Address:    echo.Map{"street": address.Street, "city_id": address.CityId, "province_id": address.ProvinceId},
		AddressTag: address.AddressTag,
		Activate:   address.Activate,
		CreatedAt:  address.CreatedAt,
		UpdatedAt:  address.UpdatedAt,
		DeletedAt:  address.DeletedAt,
	}, nil
}

func CreateAddress(req *models.AddressRequest, userId int) *Response {
	address := req.ParseToModel()

	address.UserId = userId
	address.CreatedAt = time.Now()
	address.UpdatedAt = time.Now()

	_, err := models.GetAddressByUserId(userId)
	if err != nil {
		address.Activate = "y"
	} else {
		address.Activate = "n"
	}

	err = models.CreateAddress(address)
	if err != nil {
		return ErrorResponse("CREATE_USER_ADDRESS_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponse("CREATE_USER_ADDRESS_SUCCESS", nil, http.StatusCreated)
}

func UpdateAddress(req *models.AddressRequest, addressId int) *Response {
	address := req.ParseToModel()

	address.UpdatedAt = time.Now()

	err := models.UpdateAddress(address, addressId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ErrorResponse("UPDATE_USER_ADDRESS_FAILED", "NOT_FOUND", http.StatusNotFound)
		}
		return ErrorResponse("UPDATE_USER_ADDRESS_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponse("UPDATE_USER_ADDRESS_SUCCESS", nil, http.StatusCreated)
}

func DeleteAddress(addressId int) *Response {

	address, err := models.GetAddressById(addressId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrorResponse("DELETE_ADDRESS_FAILED", "NOT_FOUND", http.StatusNotFound)
		}
		return ErrorResponse("DELETE_ADDRESS_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	if address.Activate == "y" {
		return ErrorResponse("DELETE_ADDRESS_FAILED", "ACTIVE_ADDRESS_CANT_BE_DELETED", http.StatusBadRequest)
	}

	err = models.DeleteAddress(addressId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ErrorResponse("DELETE_USER_ADDRESS_FAILED", "NOT_FOUND", http.StatusNotModified)
		}
		return ErrorResponse("DELETE_USER_ADDRESS_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponse("DELETE_USER_ADDRESS_SUCCESS", nil, http.StatusOK)
}

func UpdateActivateAddress(req *models.UpdateActivateAddressRequest, addressId int) *Response {
	status := req.ParseToModel()

	address, err := models.GetAddressById(addressId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrorResponse("GET_ADDRESS_DETAIL_FAILED", "NOT_FOUND", http.StatusNotFound)
		}
		return ErrorResponse("GET_ADDRESS_DETAIL_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	activate, err := models.GetActivateUserAddressByUserId(address.UserId)
	if err == nil {
		activate.Activate = "n"
		err = models.UpdateAddress(activate, activate.Id)
		if err != nil {
			return ErrorResponse("UPDATE_ACTIVATE_ADDRESS_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		}
	}

	err = models.UpdateAddress(status, addressId)
	if err != nil {
		return ErrorResponse("UPDATE_ACTIVATE_ADDRESS_FAILED", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return SuccessResponse("UPDATE_ACTIVATE_ADDRESS_SUCCESS", nil, http.StatusAccepted)
}
