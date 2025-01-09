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
	Limit          int    `query:"limit" validate:"number,gte=0"`
	Offset         int    `query:"offset" validate:"number,gte=0"`
	IdentityNumber string `query:"identityNumber"`
	Name           string `query:"name"`
	Gender         Gender `query:"gender" validate:"oneof=male female"`
	DepartmentId   string `query:"departmentId" validate:"number"`
}

type CreateEmployeePayload struct {
	IdentityNumber   string `json:"identityNumber" validate:"required,min=5,max=33"`
	Name             string `json:"name" validate:"required,min=4,max=33"`
	Gender           Gender `json:"gender" validate:"required,oneof=male female"`
	DepartmentId     string `json:"departmentId" validate:"required,number"` // Bisa lgsg int/bigint ga ya?
	EmployeeImageUri string `json:"employeeImageUri" validate:"required,uri"`
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
