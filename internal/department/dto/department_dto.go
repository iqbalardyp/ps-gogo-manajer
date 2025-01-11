package dto

type Department struct {
	Name string
	UserId string
}

type CreateDepartmentPayload struct{
	Name string `json:"name" validate:"required,min=4,max=33"`
	DepartmentId string `json:"departmentId"`
}

