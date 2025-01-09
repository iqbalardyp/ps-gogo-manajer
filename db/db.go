package db

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

var (
	pgInstance *Postgres
	pgOnce     sync.Once
)

func Connect(ctx context.Context, dbUrl string) (*Postgres, error) {
	var err error

	pgOnce.Do(func() {
		pool, connErr := pgxpool.New(ctx, dbUrl)
		if connErr != nil {
			err = connErr
			return
		}
		pgInstance = &Postgres{pool}
	})

	return pgInstance, err
}
