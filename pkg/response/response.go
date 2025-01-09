package response

import "net/http"

type HealthCheck struct {
	AppName string `json:"app_name"`
	Status  string `json:"status"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}



type CustomError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

// implement error interface
func (c *CustomError) Error() string {
	return c.Message
}

func Conflict(msg string) *CustomError {
	return &CustomError{
		Message:    msg,
		StatusCode: http.StatusConflict,
	}
}

func NotFound(msg string) *CustomError {
	return &CustomError{
		Message:    msg,
		StatusCode: http.StatusNotFound,
	}
}

func BadRequest(msg string) *CustomError {
	return &CustomError{
		Message:    msg,
		StatusCode: http.StatusBadRequest,
	}
}

func Unauthorized(msg string) *CustomError {
	return &CustomError{
		Message:    msg,
		StatusCode: http.StatusUnauthorized,
	}
}

func ServerError(msg string) *CustomError {
	return &CustomError{
		Message:    msg,
		StatusCode: http.StatusInternalServerError,
	}
}