package usecase

import (
	"context"
	"ps-gogo-manajer/internal/employee/dto"
	"ps-gogo-manajer/internal/employee/repository"
)

type EmployeeUsecase struct {
	employeeRepo repository.EmployeeRepository
}

func NewEmployeeUsecase(employeeRepo repository.EmployeeRepository) *EmployeeUsecase {
	return &EmployeeUsecase{
		employeeRepo: employeeRepo,
	}
}

func (u *EmployeeUsecase) GetEmployeeByIdentityNumber(ctx context.Context, identityNumber string) (*dto.Employee, error) {
	return u.employeeRepo.GetEmployeeByIdentityNumber(ctx, identityNumber)
}
