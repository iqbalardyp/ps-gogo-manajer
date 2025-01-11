package dto

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

type GetEmployeeParams struct {
	Limit          int
	Offset         int
	Gender         string
	IdentityNumber string `query:"identityNumber" validate:"omitempty"`
	Name           string `query:"name" validate:"omitempty"`
	DepartmentId   int
}

type CreateEmployeePayload struct {
	IdentityNumber   string `json:"identityNumber" validate:"required,min=5,max=33"`
	Name             string `json:"name" validate:"required,min=4,max=33"`
	Gender           Gender `json:"gender" validate:"required,oneof=male female"`
	DepartmentId     string `json:"departmentId" validate:"required,number"`
	EmployeeImageUri string `json:"employeeImageUri" validate:"omitempty,required,uri"`
}

type PatchEmployeePayload struct {
	IdentityNumber   string `json:"identityNumber" validate:"omitempty,min=5,max=33"`
	Name             string `json:"name" validate:"omitempty,min=4,max=33"`
	Gender           Gender `json:"gender" validate:"omitempty,oneof=male female"`
	DepartmentId     string `json:"departmentId" validate:"omitempty,number"`
	EmployeeImageUri string `json:"employeeImageUri" validate:"omitempty,uri"`
}

type UpdateDeletePathParam struct {
	IdentityNumber string `param:"identityNumber" validate:"required"`
}
