package routes

import (
	// "ps-gogo-manajer/internal/employee/handler"

	employeeHandler "ps-gogo-manajer/internal/employee/handler"
	departmentHandler "ps-gogo-manajer/internal/department/handler"

	"github.com/labstack/echo/v4"
)

type RouteConfig struct {
	App             *echo.Echo
	EmployeeHandler *employeeHandler.EmployeeHandler
	DepartmentHandler *departmentHandler.DepartmentHandler
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
	r.setupDepartmentRoute(v1)
}

func (r *RouteConfig) setupEmployeeRoute(api *echo.Group) {
	employee := api.Group("/employee")
	employee.GET("", r.EmployeeHandler.GetListEmployee)
	employee.POST("", r.EmployeeHandler.CreateEmployee)
	employee.PATCH("/:identityNumber", r.EmployeeHandler.UpdateEmployee)
	employee.DELETE("/:identityNumber", r.EmployeeHandler.DeleteEmployee)
}

func (r *RouteConfig) setupDepartmentRoute(api *echo.Group){
	department := api.Group("/department")

	department.GET("",r.DepartmentHandler.GetListDepartment)
	department.POST("", r.DepartmentHandler.CreateDepartment)
	department.PATCH("/:departmentId", r.DepartmentHandler.UpdateDepartment)
	department.DELETE("/:departmentId",r.DepartmentHandler.DeleteDepartment)
}
