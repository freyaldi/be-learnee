package util

import "net/http"

type response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func SuccessResponse(message string, statusCode int, data interface{}) response {
	return response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}

func ErrorResponse(message string, statusCode int) response {
	return response{
		StatusCode: statusCode,
		Message:    message,
		Data:       nil,
	}
}

func UnauthorizedError() response {
	return response{
		StatusCode: http.StatusUnauthorized,
		Message:    "Unauthorized",
		Data:       nil,
	}
}
