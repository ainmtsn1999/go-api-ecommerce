package views

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func SuccessResponse(msg string, payload interface{}, statusCode int) *Response {
	return &Response{
		Status:  statusCode,
		Message: msg,
		Payload: payload,
	}
}

func ErrorResponse(msg string, err string, statusCode int) *Response {
	return &Response{
		Status:  statusCode,
		Message: msg,
		Error:   err,
	}
}
