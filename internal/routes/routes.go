package routes

import (
	"net/http"
	employeeHandler "ps-gogo-manajer/internal/employee/handler"
	userHandler "ps-gogo-manajer/internal/user/handler"
	"ps-gogo-manajer/pkg/response"

	"github.com/labstack/echo/v4"
)

type RouteConfig struct {
	App             *echo.Echo
	EmployeeHandler *employeeHandler.EmployeeHandler
	UserHandler     *userHandler.UserHandler
	AuthMiddleware  echo.MiddlewareFunc
}

func (r *RouteConfig) SetupRoutes() {
	r.setupPublicRoutes()
	r.setupAuthRoutes()
}

func (r *RouteConfig) setupPublicRoutes() {
	r.App.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, response.BaseResponse{
			Status:  "Ok",
			Message: "",
		})
	})
	auth := r.App.Group("/auth")
	auth.POST("/", r.UserHandler.AuthenticateUser)
}
func (r *RouteConfig) setupAuthRoutes() {
	v1 := r.App.Group("/v1", r.AuthMiddleware)

	r.setupEmployeeRoute(v1)
	r.setupUserRoute(v1)
}

func (r *RouteConfig) setupEmployeeRoute(api *echo.Group) {
	employee := api.Group("/employee")
	employee.GET("", r.EmployeeHandler.GetListEmployee)
	employee.POST("", r.EmployeeHandler.CreateEmployee)
	employee.PUT("/:identityNumber", r.EmployeeHandler.UpdateEmployee)
	employee.DELETE("/:identityNumber", r.EmployeeHandler.DeleteEmployee)
}

func (r *RouteConfig) setupUserRoute(api *echo.Group) {
	user := api.Group("/user")
	user.GET("/", r.UserHandler.GetUser)
	user.PATCH("/", r.UserHandler.UpdateUser)
}
