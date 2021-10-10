package utility

type CustomResponse struct {
	IsSuccess bool        `json:"is_success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

func NewSuccessResponse(message string, data interface{}) CustomResponse {
	return CustomResponse{
		IsSuccess: true,
		Message:   message,
		Data:      data,
	}
}

func NewFailResponse(message string) CustomResponse {
	return CustomResponse{
		IsSuccess: false,
		Message:   message,
		Data:      nil,
	}
}