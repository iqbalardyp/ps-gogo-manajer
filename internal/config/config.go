package config

import (
	"net/http"
	"ps-gogo-manajer/db"
	employeeHandler "ps-gogo-manajer/internal/employee/handler"
	employeeRepository "ps-gogo-manajer/internal/employee/repository"
	employeeUsecase "ps-gogo-manajer/internal/employee/usecase"

	departmentHandler "ps-gogo-manajer/internal/department/handler"
	departmentRepository "ps-gogo-manajer/internal/department/repository"
	departmentUsecase "ps-gogo-manajer/internal/department/usecase"
	"ps-gogo-manajer/internal/routes"
	"ps-gogo-manajer/pkg/response"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type BootstrapConfig struct {
	App       *echo.Echo
	DB        *db.Postgres
	Log       *logrus.Logger
	Validator *validator.Validate
}

func Bootstrap(config *BootstrapConfig) {
	employeeRepo := employeeRepository.NewEmployeeRepository(config.DB.Pool)
	employeeUseCase := employeeUsecase.NewEmployeeUsecase(*employeeRepo)
	employeeHandler := employeeHandler.NewEmployeeHandler(*employeeUseCase, config.Validator)

	//department variable
	departmentRepo := departmentRepository.NewDepartmentRepository(config.DB.Pool)
	departmentUsecase := departmentUsecase.NewDepartmentUsecases(*departmentRepo)
	departmentHandler := departmentHandler.NewDepartmentHandler(*departmentUsecase,config.Validator)

	routes := routes.RouteConfig{
		App:             config.App,
		EmployeeHandler: employeeHandler,
		DepartmentHandler : departmentHandler,
	}

	// * Middleware
	config.App.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Timeout",
		Timeout:      30 * time.Second,
	}))

	// Health check
	config.App.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, response.BaseResponse{
			Status:  "Ok",
			Message: "",
		})
	})

	routes.SetupRoutes()
}
