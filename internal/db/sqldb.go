package db

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"time"
)

type SQLDatabase struct {
	DB *sql.DB
}

func (s *SQLDatabase) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return s.DB.QueryContext(ctx, query, args...)
}

func (s *SQLDatabase) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return s.DB.QueryRowContext(ctx, query, args...)
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

// flattenStructFields recursively flattens a struct's fields, including those from embedded structs.
func flattenStructFields(v reflect.Value) ([]interface{}, error) {
	var fields []interface{}

	// If the value is a pointer, dereference it
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Ensure the value is a struct
	if v.Kind() != reflect.Struct {
		return nil, errors.New("value must be a struct")
	}

	// Iterate through the struct's fields
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		// Recursively flatten embedded struct fields or add the field directly
		if field.Kind() == reflect.Struct {
			// Special handling for time.Time type
			if field.Type() == reflect.TypeOf(time.Time{}) {
				fields = append(fields, field.Addr().Interface())
			} else {
				embeddedFields, err := flattenStructFields(field)
				if err != nil {
					return nil, err
				}
				fields = append(fields, embeddedFields...)
			}
		} else {
			fields = append(fields, field.Addr().Interface())
		}
	}

	return fields, nil
}

// MarshalRowToStruct maps a single row result into a struct, supporting embedded structs.
func MarshalRowToStruct(row *sql.Row, dest interface{}) error {
	destValue := reflect.ValueOf(dest).Elem()
	fields, err := flattenStructFields(destValue)
	if err != nil {
		return err
	}

	if len(fields) == 0 {
		return errors.New("no fields to scan into")
	}

	if err := row.Scan(fields...); err != nil {
		return err
	}

	return nil
}

// MarshalRowsToStructs maps multiple row results into a slice of structs, supporting embedded structs.
func MarshalRowsToStructs(rows *sql.Rows, destSlice interface{}) error {
	destSliceValue := reflect.ValueOf(destSlice)
	if destSliceValue.Kind() != reflect.Ptr || destSliceValue.Elem().Kind() != reflect.Slice {
		return errors.New("destSlice must be a pointer to a slice")
	}

	destSliceElemType := destSliceValue.Elem().Type().Elem()

	for rows.Next() {
		// Create a new instance of the element type
		newElem := reflect.New(destSliceElemType.Elem())
		fields, err := flattenStructFields(newElem)
		if err != nil {
			return err
		}

		if len(fields) == 0 {
			return errors.New("no fields to scan into")
		}

		if err := rows.Scan(fields...); err != nil {
			return err
		}

		destSliceValue.Elem().Set(reflect.Append(destSliceValue.Elem(), newElem))
	}

	return rows.Err()
}
