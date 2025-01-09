package response

type HealthCheck struct {
	AppName string `json:"app_name"`
	Status  string `json:"status"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
