package handler

import (
	"ps-gogo-manajer/internal/department/dto"
	"ps-gogo-manajer/internal/department/usecase"
	customErrors "ps-gogo-manajer/pkg/custom-errors"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)


type DepartmentHandler struct {
	departmentUsecase usecase.DepartmentUsecase
	validator *validator.Validate
}

func newDepartmentHandler(department usecase.DepartmentUsecase, validator *validator.Validate) *DepartmentHandler{
	return &DepartmentHandler{
		departmentUsecase: department,
		validator: validator,
	}
}

func(h DepartmentHandler) CreateDepartment(ctx echo.Context) error{
	var payload dto.CreateDepartmentPayload

	if err := ctx.Bind(&payload); err != nil {
		err = errors.Wrap(customErrors.ErrBadRequest, err.Error())
	}
}

