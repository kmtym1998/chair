package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/kmtym1998/chair/generator"
	"github.com/kmtym1998/chair/postgres/client"
	"github.com/kr/pretty"
)

type SchemaLoader struct {
	DB     *sql.DB
	schema string
}

func NewSchemaLoader(dsn, schema string) (*SchemaLoader, error) {
	pgClient, err := client.New(client.Opts{
		DataSourceName: dsn,
	})
	if err != nil {
		return nil, err
	}

	return &SchemaLoader{
		DB:     pgClient.DB(),
		schema: schema,
	}, nil
}

func (s *SchemaLoader) LoadSchema(ctx context.Context) ([]generator.Table, error) {
	pgClasses, err := s.listPGClass(ctx, s.schema)
	if err != nil {
		return nil, err
	}

	pretty.Println(pgClasses)

	return nil, nil
}

// https://www.postgresql.org/docs/current/catalog-pg-class.html
// https://www.postgresql.org/docs/current/catalog-pg-namespace.html
type PGClass struct {
	OID     int    `db:"oid"`
	RelName string `db:"relname"`
	RelType int    `db:"reltype"`
}

func (s *SchemaLoader) listPGClass(ctx context.Context, schema string) ([]PGClass, error) {
	// TODO: get comment and columns  https://www.postgresql.jp/document/9.3/html/functions-info.html
	const query = "SELECT c.oid, c.relname, c.reltype " +
		"FROM pg_class c INNER JOIN pg_namespace ns ON c.relnamespace = ns.oid " +
		"WHERE relkind = 'r' AND ns.nspname = $1;"
	slog.Debug("executing query", "query", query, "schema", schema)

	rows, err := s.DB.QueryContext(ctx, query, schema)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var pgClass []PGClass
	for rows.Next() {
		var class PGClass
		if err := rows.Scan(
			&class.OID,
			&class.RelName,
			&class.RelType,
		); err != nil {
			return nil, fmt.Errorf("failed to scan pg_class: %w", err)
		}

		pgClass = append(pgClass, class)
	}

	return pgClass, nil
}
