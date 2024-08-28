package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"
)

type Database interface {
	Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Close() error
}

func GetDb() Database {
	dsn := os.Getenv("DATABASE_URL")

	sqlDB, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}

	// Configure the connection pool
	sqlDB.SetMaxOpenConns(25)                 // Maximum number of open connections to the database
	sqlDB.SetMaxIdleConns(25)                 // Maximum number of idle connections in the pool
	sqlDB.SetConnMaxIdleTime(5 * time.Minute) // Maximum amount of time a connection may be idle
	sqlDB.SetConnMaxLifetime(1 * time.Hour)   // Maximum amount of time a connection may be reused

	return &SQLDatabase{DB: sqlDB}
}
