package repository

import (
	"context"
	"fmt"
	"ps-gogo-manajer/internal/employee/dto"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeRepository struct {
	pool *pgxpool.Pool
}

func NewEmployeeRepository(pool *pgxpool.Pool) *EmployeeRepository {
	return &EmployeeRepository{pool: pool}
}

const (
	queryGetEmployeeByIdentityNumber = `
	SELECT 
		name,
		identity_number,
		gender,
		department_id,
		employee_image_uri
	FROM employees
	WHERE
		user_id = $1
		AND identity_number = $2;`
)

// Temporary
func (er *EmployeeRepository) GetEmployeeByIdentityNumber(ctx context.Context, identityNumber string) (*dto.Employee, error) {
	var employee dto.Employee
	var employeeImageUri pgtype.Text

	// Placeholder
	userID := "1"

	err := er.pool.QueryRow(ctx, queryGetEmployeeByIdentityNumber, userID, identityNumber).Scan(
		&employee.Name,
		&employee.IdentityNumber,
		&employee.Gender,
		&employee.DepartmentId,
		&employeeImageUri,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, err
		}
		fmt.Println("REPO:", err.Error())
		return nil, err
	}

	employee.EmployeeImageUri = employeeImageUri.String
	return &employee, nil
}
