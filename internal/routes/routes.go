package routes

import (
	employeeHandler "ps-gogo-manajer/internal/employee/handler"
	fileHandler "ps-gogo-manajer/internal/files/handler"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v4"
)

type RouteConfig struct {
	App             *echo.Echo
	S3Client        *s3.Client
	EmployeeHandler *employeeHandler.EmployeeHandler
	FileHandler     *fileHandler.FileHandler
}

func (r *RouteConfig) SetupRoutes() {
	// r.setupPublicRoutes()
	r.setupAuthRoutes()
}

func (r *RouteConfig) setupPublicRoutes() {
	// TODO: add public routes
	// auth := r.App.Group("/auth")
}
func (r *RouteConfig) setupAuthRoutes() {
	v1 := r.App.Group("/v1")

	// TODO: use echo-jwt
	// v1.Use(echojwt.WithConfig(echojwt.Config{
	// 	SigningKey:             []byte("secret"),
	// }))
	r.setupEmployeeRoute(v1)
	r.setupFileRoutes(v1)
}

func (r *RouteConfig) setupEmployeeRoute(api *echo.Group) {
	employee := api.Group("/employee")
	employee.GET("", r.EmployeeHandler.GetListEmployee)
	employee.POST("", r.EmployeeHandler.CreateEmployee)
	employee.PUT("/:identityNumber", r.EmployeeHandler.UpdateEmployee)
	employee.DELETE("/:identityNumber", r.EmployeeHandler.DeleteEmployee)
}

func (r *RouteConfig) setupFileRoutes(api *echo.Group) {
	api.POST("/file", r.FileHandler.UploadFile)
}
