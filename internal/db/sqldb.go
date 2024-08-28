package db

import (
	"context"
	"database/sql"
)

type SQLDatabase struct {
	DB *sql.DB
}

func (s *SQLDatabase) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return s.DB.QueryContext(ctx, query, args...)
}

func (s *SQLDatabase) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return s.DB.ExecContext(ctx, query, args...)
}

func (s *SQLDatabase) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return s.DB.BeginTx(ctx, opts)
}

func (s *SQLDatabase) Close() error {
	return s.DB.Close()
}
