package usecase

import (
	"context"
	"ps-gogo-manajer/internal/department/dto"
	"ps-gogo-manajer/internal/department/repository"
)

type DepartmentUsecase struct{
	departmentRepo repository.DepartmentRepository
}

func NewDepartmentUsecases(departmentRepo repository.DepartmentRepository)*DepartmentUsecase{
	return &DepartmentUsecase{
		departmentRepo: departmentRepo,
	}
}

func (u *DepartmentUsecase) CreateEmployee(ctx context.Context, userID int, payload *dto.CreateDepartmentPayload) (*dto.Department, error) {
	
	return u.departmentRepo.CreateDepartment(ctx, userID, payload)
}