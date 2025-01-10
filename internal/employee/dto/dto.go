package dto

import "github.com/jackc/pgx/v5/pgtype"

type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
)

type Employee struct {
	Name             string `json:"name"`
	IdentityNumber   string `json:"identityNumber"`
	Gender           Gender `json:"gender"`
	DepartmentId     string `json:"departmentId"`
	EmployeeImageUri string `json:"employeeImageUri"`
}

// TODO:
type GetEmployeeParams struct {
	Limit          int
	Offset         int
	Gender         pgtype.Text
	IdentityNumber pgtype.Text `query:"identityNumber" validate:"omitempty"`
	Name           pgtype.Text `query:"name" validate:"omitempty"`
	DepartmentId   *int
}

type CreateEmployeePayload struct {
	IdentityNumber   string `json:"identityNumber" validate:"required,min=5,max=33"`
	Name             string `json:"name" validate:"required,min=4,max=33"`
	Gender           Gender `json:"gender" validate:"required,oneof=male female"`
	DepartmentId     string `json:"departmentId" validate:"required,number"` // Bisa lgsg int/bigint ga ya?
	EmployeeImageUri string `json:"employeeImageUri" validate:"omitempty,required,uri"`
}

type PatchEmployeePayload struct {
	IdentityNumber   string `json:"identityNumber" validate:"min=5,max=33"`
	Name             string `json:"name" validate:"min=4,max=33"`
	Gender           Gender `json:"gender" validate:"oneof=male female"`
	DepartmentId     string `json:"departmentId" validate:"number"` // Bisa lgsg int/bigint ga ya?
	EmployeeImageUri string `json:"employeeImageUri" validate:"uri"`
}

type UpdateDeletePathParam struct {
	IdentityNumber string `param:"identityNumber" validate:"required"`
}
