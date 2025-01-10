package repository

import (
	"context"
	"ps-gogo-manajer/internal/employee/dto"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type EmployeeRepository struct {
	pool *pgxpool.Pool
}

func NewEmployeeRepository(pool *pgxpool.Pool) *EmployeeRepository {
	return &EmployeeRepository{pool: pool}
}

const (
	queryCheckIfEmployeeExists = `
	SELECT EXISTS (
		SELECT id
		FROM employees
		WHERE
			user_id = @userID
			AND identity_number = @identityNumber
	) is_exists;`
	//  TODO:
	queryGetListEmployee = `
	SELECT
		name,
		identity_number,
		gender,
		department_id,
		employee_image_uri
	FROM employees
	WHERE
		user_id = @userID
		-- AND (@identityNumber is null)
		-- AND (@identityNumber is null OR identity_number ilike '@identityNumber%')
		-- AND (@name is null OR name ilike '%@name%')
		AND (@gender is null OR gender = @gender)
		-- AND (@departmentID is null OR department_id = @departmentID)
	OFFSET @offset
	LIMIT @limit`
	queryCreateEmployee = `
	INSERT INTO employees(name, gender, identity_number, department_id, user_id, employee_image_uri)
	VALUES (@name, @gender, @identityNumber, @departmentID, @userID, @employeeImageUri)
	RETURNING name, identity_number, gender, department_id, employee_image_uri;`
	// TODO:
	// queryUpdateEmployee = ""
	queryDeleteEmployee = "DELETE FROM employees WHERE user_id = @userID AND identity_number = @identityNumber;"
)

func (r *EmployeeRepository) CheckIfEmployeeExists(ctx context.Context, userID int, identityNumber string) (bool, error) {
	var isExist bool
	args := pgx.NamedArgs{
		"userID":         userID,
		"identityNumber": identityNumber,
	}

	err := r.pool.QueryRow(ctx, queryCheckIfEmployeeExists, args).Scan(&isExist)
	if err != nil {
		return false, errors.Wrap(err, "failed to check if employee exists")
	}

	return isExist, nil
}

func (r *EmployeeRepository) CreateEmployee(ctx context.Context, userID int, payload *dto.CreateEmployeePayload) (*dto.Employee, error) {
	var employee dto.Employee

	args := pgx.NamedArgs{
		"name":             payload.Name,
		"gender":           payload.Gender,
		"identityNumber":   payload.IdentityNumber,
		"departmentID":     payload.DepartmentId,
		"userID":           userID,
		"employeeImageUri": payload.EmployeeImageUri,
	}
	err := r.pool.QueryRow(ctx, queryCreateEmployee, args).Scan(
		&employee.Name,
		&employee.IdentityNumber,
		&employee.Gender,
		&employee.DepartmentId,
		&employee.EmployeeImageUri,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create employee")
	}

	return &employee, nil
}

func (r *EmployeeRepository) GetListEmployee(ctx context.Context, userID int, payload *dto.GetEmployeeParams) (*[]dto.Employee, error) {
	var employees []dto.Employee
	args := pgx.NamedArgs{
		"userID":         userID,
		"identityNumber": &payload.IdentityNumber,
		"name":           &payload.Name,
		"gender":         payload.Gender,
		"departmentID":   &payload.DepartmentId,
		"limit":          payload.Limit,
		"offset":         payload.Offset,
	}

	rows, err := r.pool.Query(ctx, queryGetListEmployee, args)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get list employee")
	}

	for rows.Next() {
		employee := dto.Employee{}
		imgUri := new(pgtype.Text)

		err := rows.Scan(
			&employee.Name,
			&employee.IdentityNumber,
			&employee.Gender,
			&employee.DepartmentId,
			imgUri,
		)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse sql response")
		}

		employee.EmployeeImageUri = imgUri.String
		employees = append(employees, employee)
	}

	return &employees, nil
}

func (r *EmployeeRepository) UpdateEmployee(ctx context.Context, userID int, identityNumber string, payload *dto.PatchEmployeePayload) (*dto.Employee, error) {
	var employee dto.Employee
	return &employee, nil
}

func (r *EmployeeRepository) DeleteEmployee(ctx context.Context, userID int, identityNumber string) error {
	args := pgx.NamedArgs{
		"userID":         userID,
		"identityNumber": identityNumber,
	}

	_, err := r.pool.Exec(ctx, queryDeleteEmployee, args)
	if err != nil {
		return errors.Wrap(err, "failed to delete employee")
	}

	return nil
}
