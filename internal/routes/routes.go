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
	r.SetupEmployeeRoute()
}

func (r *RouteConfig) SetupEmployeeRoute() {
	// employee := r.App.Group("/employee")
	// employee.GET("/:identityNumber", r.EmployeeHandler.GetEmployeeByIdentityNumber)
}
