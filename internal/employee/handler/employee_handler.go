package handler

import (
	"net/http"
	"ps-gogo-manajer/internal/employee/dto"
	"ps-gogo-manajer/internal/employee/usecase"
	customErrors "ps-gogo-manajer/pkg/custom-errors"
	customValidators "ps-gogo-manajer/pkg/custom-validators"
	"ps-gogo-manajer/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type EmployeeHandler struct {
	employeeUsecase usecase.EmployeeUsecase
	validator       *validator.Validate
}

const (
	DEFAULT_LIMIT  = 5
	DEFAULT_OFFSET = 0
)

func NewEmployeeHandler(employeeUsecase usecase.EmployeeUsecase, validator *validator.Validate) *EmployeeHandler {
	return &EmployeeHandler{
		employeeUsecase: employeeUsecase,
		validator:       validator,
	}
}

func (h EmployeeHandler) CreateEmployee(ctx echo.Context) error {
	var payload dto.CreateEmployeePayload
	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(response.WriteErrorResponse(err))
	}

	// TODO: Add payload validator
	if err := h.validator.Struct(payload); err != nil {
		err = errors.Wrap(customErrors.ErrBadRequest, err.Error())
		return ctx.JSON(response.WriteErrorResponse(err))
	}

	// TODO: Get userID from auth token
	userID := 1

	employee, err := h.employeeUsecase.CreateEmployee(ctx.Request().Context(), userID, &payload)
	if err != nil {
		return ctx.JSON(response.WriteErrorResponse(err))
	}

	return ctx.JSON(http.StatusCreated, employee)
}

func (h EmployeeHandler) GetListEmployee(ctx echo.Context) error {
	genderStr := ctx.QueryParam("gender")
	gender, isValid := customValidators.ParseGender(genderStr)
	if !isValid {
		return ctx.JSON(http.StatusOK, make([]string, 0))
	}

	departmentIDStr := ctx.QueryParam("departmentID")
	departmentID, isValid := customValidators.ParseDepartmentID(departmentIDStr)
	if !isValid {
		return ctx.JSON(http.StatusOK, make([]string, 0))
	}

	limitStr := ctx.QueryParam("limit")
	offsetStr := ctx.QueryParam("offset")

	limit := customValidators.ParseLimitOffset(limitStr, DEFAULT_LIMIT)
	offset := customValidators.ParseLimitOffset(offsetStr, DEFAULT_OFFSET)

	payload := dto.GetEmployeeParams{
		Limit:        limit,
		Offset:       offset,
		Gender:       gender,
		DepartmentId: departmentID,
	}

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(response.WriteErrorResponse(err))
	}

	// TODO: get userID from auth token
	userID := 1

	employees, err := h.employeeUsecase.GetListEmployee(ctx.Request().Context(), userID, &payload)
	if err != nil {
		return ctx.JSON(response.WriteErrorResponse(err))
	}

	if len(*employees) == 0 {
		return ctx.JSON(http.StatusOK, make([]string, 0))
	}

	return ctx.JSON(http.StatusOK, &employees)
}

func (h EmployeeHandler) UpdateEmployee(ctx echo.Context) error {
	identityNumber := ctx.Param("identityNumber")
	if identityNumber == "" {
		err := errors.Wrap(customErrors.ErrBadRequest, "identity number required")
		return ctx.JSON(response.WriteErrorResponse(err))
	}

	var payload dto.PatchEmployeePayload
	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(response.WriteErrorResponse(err))
	}
	// TODO: add payload validator

	// TODO: get userID from auth token
	userID := 1

	employee, err := h.employeeUsecase.UpdateEmployee(ctx.Request().Context(), userID, identityNumber, &payload)
	if err != nil {
		return ctx.JSON(response.WriteErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, employee)
}

func (h EmployeeHandler) DeleteEmployee(ctx echo.Context) error {
	var payload dto.UpdateDeletePathParam
	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(response.WriteErrorResponse(err))
	}

	// TODO: get userID from auth token
	userID := 1

	err := h.employeeUsecase.DeleteEmployee(ctx.Request().Context(), userID, payload.IdentityNumber)
	if err != nil {
		return ctx.JSON(response.WriteErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, response.BaseResponse{
		Status:  http.StatusText(http.StatusOK),
		Message: "deleted",
	})
}
