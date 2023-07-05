package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

const (
	defaultMaxPoolSize  = 1
	defaultConnAttempts = 10
	defaultConnTimeout  = time.Second
)

type PgxPool interface {
	Close()
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	Ping(ctx context.Context) error
}

type Postgres struct {
	Pool PgxPool
}

func NewClient(ctx context.Context, url string) (*Postgres, error) {
	var pg = new(Postgres)
	var err error
	pg.Pool, err = pgxpool.New(ctx, url)
	if err != nil {
		return pg, fmt.Errorf("Unable to create connecction pool: %v\n", err)
	}

	return pg, nil
}
