package response

type ErrorApiResponse struct {
	Code    uint   `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func ErrorHttp(code uint, message string) ErrorApiResponse {
	return ErrorApiResponse{
		Code:    code,
		Message: message,
	}
}
