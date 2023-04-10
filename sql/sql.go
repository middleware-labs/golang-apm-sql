package sql

import (
	"database/sql"
	"github.com/middleware-labs/opentelemetry-go-extra/otelsql"
	"go.opentelemetry.io/otel/attribute"
)

func Open(driverName, dsn string, opts ...otelsql.Option) (*sql.DB, error) {
	db, err := otelsql.Open(driverName, dsn, opts...)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func WithDBSystem(system string) otelsql.Option {
	return otelsql.WithDBSystem(system)
}

func WithAttributes(attrs ...attribute.KeyValue) otelsql.Option {
	return otelsql.WithAttributes(attrs...)
}
