package routes

import (
	// "ps-gogo-manajer/internal/employee/handler"

	"github.com/labstack/echo/v4"
)

type RouteConfig struct {
	App             *echo.Echo
	// EmployeeHandler *handler.EmployeeHandler
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
}

func (r *RouteConfig) setupEmployeeRoute(api *echo.Group) {
	employee := api.Group("/employee")
	employee.GET("", r.EmployeeHandler.GetListEmployee)
	employee.POST("", r.EmployeeHandler.CreateEmployee)
	employee.PUT("/:identityNumber", r.EmployeeHandler.UpdateEmployee)
	employee.DELETE("/:identityNumber", r.EmployeeHandler.DeleteEmployee)
}
