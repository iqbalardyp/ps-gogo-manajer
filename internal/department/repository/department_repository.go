package repository

import (
	"context"
	"ps-gogo-manajer/internal/department/dto"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)


type DepartmentRepository struct {
	pool *pgxpool.Pool
}


func NewDepartmentRepository(pool *pgxpool.Pool) *DepartmentRepository {
	return &DepartmentRepository{pool: pool}
}

const (
	queryCreateDepartment = `
	INSERT INTO departments
	(
		name,
		user_id,
	)
	VALUES (@name,@userID)
	RETURNING name,user_id;`
)

func(r *DepartmentRepository) CreateDepartment(ctx context.Context, userID int, payload *dto.CreateDepartmentPayload)(*dto.Department,error){
	var department dto.Department

	args := pgx.NamedArgs{
		"id" : payload.DepartmentId,
		"name": payload.Name,
	}

	err := r.pool.QueryRow(ctx,queryCreateDepartment,args).Scan(
		&department.Name,
	)

	if err != nil {
		return nil, errors.Wrap(err,"failed to create department")
	}

	return &department,nil
}