package config

import (
	"ps-gogo-manajer/db"
	employeeHandler "ps-gogo-manajer/internal/employee/handler"
	employeeRepository "ps-gogo-manajer/internal/employee/repository"
	employeeUsecase "ps-gogo-manajer/internal/employee/usecase"
	auth "ps-gogo-manajer/internal/middleware"
	"ps-gogo-manajer/internal/routes"
	userHandler "ps-gogo-manajer/internal/user/handler"
	userRepository "ps-gogo-manajer/internal/user/repository"
	userUsecase "ps-gogo-manajer/internal/user/usecase"
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

	userRepo := userRepository.NewUserRepository(config.DB.Pool)
	userUseCase := userUsecase.NewUserUseCase(*userRepo)
	userHandler := userHandler.NewUserHandler(*userUseCase, config.Validator)

	// * Middleware
	config.App.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Timeout",
		Timeout:      30 * time.Second,
	}))
	authMiddleware := auth.Auth()

	routes := routes.RouteConfig{
		App:             config.App,
		EmployeeHandler: employeeHandler,
		UserHandler:     userHandler,
		AuthMiddleware:  authMiddleware,
	}

	routes.SetupRoutes()
}
