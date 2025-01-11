package config

import (
	"net/http"
	// "os"
	"ps-gogo-manajer/db"
	// employeeHandler "ps-gogo-manajer/internal/employee/handler"
	// employeeRepository "ps-gogo-manajer/internal/employee/repository"
	// employeeUsecase "ps-gogo-manajer/internal/employee/usecase"
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
	// employeeRepo := employeeRepository.NewEmployeeRepository(config.DB.Pool)
	// employeeUseCase := employeeUsecase.NewEmployeeUsecase(*employeeRepo)
	// employeeHandler := employeeHandler.NewEmployeeHandler(*employeeUseCase)
	routes := routes.RouteConfig{
		App:             config.App,
		// EmployeeHandler: employeeHandler,
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
