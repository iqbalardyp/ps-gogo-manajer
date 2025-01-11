package usecase

import (
	"context"
	"ps-gogo-manajer/internal/employee/dto"
	"ps-gogo-manajer/internal/employee/repository"
	customErrors "ps-gogo-manajer/pkg/custom-errors"

	"github.com/pkg/errors"
)

type EmployeeUsecase struct {
	employeeRepo repository.EmployeeRepository
}

func NewEmployeeUsecase(employeeRepo repository.EmployeeRepository) *EmployeeUsecase {
	return &EmployeeUsecase{
		employeeRepo: employeeRepo,
	}
}

func (u *EmployeeUsecase) CreateEmployee(ctx context.Context, userID int, payload *dto.CreateEmployeePayload) (*dto.Employee, error) {
	isIdentityNumberExists, err := u.employeeRepo.CheckIfEmployeeExists(ctx, userID, payload.IdentityNumber)
	if err != nil {
		return nil, err
	}

	if isIdentityNumberExists {
		return nil, errors.Wrap(customErrors.ErrConflict, "identity number already exists")
	}

	return u.employeeRepo.CreateEmployee(ctx, userID, payload)
}

func (u *EmployeeUsecase) GetListEmployee(ctx context.Context, userID int, payload *dto.GetEmployeeParams) (*[]dto.Employee, error) {
	return u.employeeRepo.GetListEmployee(ctx, userID, payload)
}

func (u *EmployeeUsecase) UpdateEmployee(ctx context.Context, userID int, identityNumber string, payload *dto.PatchEmployeePayload) (*dto.Employee, error) {
	// Validate if employee exists
	isEmployeeExists, err := u.employeeRepo.CheckIfEmployeeExists(ctx, userID, identityNumber)
	if err != nil {
		return nil, err
	}

	if !isEmployeeExists {
		return nil, errors.Wrap(customErrors.ErrNotFound, "employee not found")
	}

	// * Validate if payload's identityNumber exists
	isIdentityNumberExists, err := u.employeeRepo.CheckIfEmployeeExists(ctx, userID, payload.IdentityNumber)
	if err != nil {
		return nil, err
	}

	if isIdentityNumberExists {
		return nil, errors.Wrap(customErrors.ErrConflict, "identity number already exists")
	}

	return u.employeeRepo.UpdateEmployee(ctx, userID, identityNumber, payload)
}

func (u *EmployeeUsecase) DeleteEmployee(ctx context.Context, userID int, identityNumber string) error {
	// * Validate if employee exists
	isEmployeeExists, err := u.employeeRepo.CheckIfEmployeeExists(ctx, userID, identityNumber)
	if err != nil {
		return err
	}

	if !isEmployeeExists {
		return errors.Wrap(customErrors.ErrNotFound, "employee not found")
	}

	return u.employeeRepo.DeleteEmployee(ctx, userID, identityNumber)
}
