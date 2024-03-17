package postgres

import (
	"context"
	"database/sql"

	"github.com/kmtym1998/chair/generator"
	"github.com/kmtym1998/chair/postgres/client"
)

type SchemaLoader struct {
	DB *sql.DB
}

func NewSchemaLoader(dsn string) (*SchemaLoader, error) {
	pgClient, err := client.New(client.Opts{
		DataSourceName: dsn,
	})
	if err != nil {
		return nil, err
	}

	return &SchemaLoader{
		DB: pgClient.DB(),
	}, nil
}

func (s *SchemaLoader) LoadSchema(ctx context.Context) ([]generator.Table, error) {
	return nil, nil
}
