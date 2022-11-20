package views

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload,omitempty"`
	Query   interface{} `json:"query,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func SuccessResponse(msg string, payload interface{}, statusCode int) *Response {
	return &Response{
		Status:  statusCode,
		Message: msg,
		Payload: payload,
	}
}

func ErrorResponse(msg string, err interface{}, statusCode int) *Response {
	return &Response{
		Status:  statusCode,
		Message: msg,
		Error:   err,
	}
}

func SuccessResponseWithQuery(msg string, payload interface{}, query interface{}, statusCode int) *Response {
	return &Response{
		Status:  statusCode,
		Message: msg,
		Payload: payload,
		Query:   query,
	}
}
