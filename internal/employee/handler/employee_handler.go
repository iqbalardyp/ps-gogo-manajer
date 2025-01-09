package handler

import (
	"net/http"
	"ps-gogo-manajer/internal/employee/usecase"
	"ps-gogo-manajer/pkg/response"

	"github.com/labstack/echo/v4"
)

type EmployeeHandler struct {
	employeeUsecase usecase.EmployeeUsecase
}

func NewEmployeeHandler(employeeUsecase usecase.EmployeeUsecase) *EmployeeHandler {
	return &EmployeeHandler{
		employeeUsecase: employeeUsecase,
	}
}

func (h EmployeeHandler) GetEmployeeByIdentityNumber(ctx echo.Context) error {
	identityNumber := ctx.Param("identityNumber")
	if identityNumber == "" {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Identity Number Required",
		})
	}
	employee, err := h.employeeUsecase.GetEmployeeByIdentityNumber(ctx.Request().Context(), identityNumber)
	if employee == nil {
		return ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Message: "Not found",
		})
	}
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: "Internal Server Error",
		})
	}
	return ctx.JSON(http.StatusOK, employee)
}
