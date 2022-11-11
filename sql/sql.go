package sql

import (
	"database/sql"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
)

func Open(driverName, dsn string) (*sql.DB, error) {
	db, err := otelsql.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
