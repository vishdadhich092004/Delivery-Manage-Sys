package response

type ApiResponseType struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ApiResponse(status int, message string, data interface{}) ApiResponseType {
	return ApiResponseType{
		Success: status < 400,
		Message: message,
		Data:    data,
	}
}
