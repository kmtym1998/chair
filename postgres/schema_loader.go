package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"regexp"
	"strings"

	"github.com/kmtym1998/chair/generator"
	"github.com/kmtym1998/chair/postgres/client"
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

func (s *SchemaLoader) LoadTableSchemas(ctx context.Context) ([]generator.Table, error) {
	tables, err := s.listTables(ctx, s.schema)
	if err != nil {
		return nil, err
	}

	columns, err := s.listColumns(ctx, s.schema)
	if err != nil {
		return nil, err
	}

	tableSchemas := make([]generator.Table, len(tables))
	for i, table := range tables {
		columnSchemas := make([]generator.Column, 0, len(columns))
		for _, column := range columns {
			if table.TableName == column.TableName && table.SchemaName == column.SchemaName {
				columnSchemas = append(columnSchemas, generator.Column{
					Name:     column.ColumnName,
					Comment:  column.Comment.String,
					Type:     column.DataType,
					OrderAsc: column.Position,
				})
			}
		}

		tableSchemas[i] = generator.Table{
			Name:    table.TableName,
			Comment: table.Comment.String,
			Columns: columnSchemas,
		}
	}

	return tableSchemas, nil
}

type Table struct {
	ID         int            `db:"relid"`
	SchemaName string         `db:"schemaname"`
	TableName  string         `db:"relname"`
	Comment    sql.NullString `db:"description"`
}

func (s *SchemaLoader) listTables(ctx context.Context, schema string) ([]Table, error) {
	const query = "SELECT relid, schemaname, relname, d.description " +
		"FROM pg_stat_user_tables sut " +
		"INNER JOIN pg_description d ON d.objsubid = 0 AND sut.relid = d.objoid " +
		"WHERE sut.schemaname = $1;"
	slog.Debug("executing query", "query", query, "schema", schema)

	rows, err := s.DB.QueryContext(ctx, query, schema)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var tables []Table
	for rows.Next() {
		var table Table
		if err := rows.Scan(
			&table.ID,
			&table.SchemaName,
			&table.TableName,
			&table.Comment,
		); err != nil {
			return nil, fmt.Errorf("failed to scan tables: %w", err)
		}

		tables = append(tables, table)
	}

	return tables, nil
}

type Column struct {
	SchemaName string         `db:"table_schema"`
	TableName  string         `db:"table_name"`
	ColumnName string         `db:"column_name"`
	DataType   string         `db:"udt_name"`
	Position   int            `db:"ordinal_position"`
	Comment    sql.NullString `db:"description"`
}

func (s *SchemaLoader) listColumns(ctx context.Context, schema string) ([]Column, error) {
	const query = `
SELECT
	information_schema.columns.table_schema,
	information_schema.columns.table_name,
	information_schema.columns.column_name,
	information_schema.columns.udt_name,
	information_schema.columns.ordinal_position,
	(
		SELECT
			description
		FROM
			pg_description
		WHERE
			pg_description.objoid = pg_stat_user_tables.relid
			AND pg_description.objsubid = information_schema.columns.ordinal_position
	) AS description
FROM
	pg_stat_user_tables,
	information_schema.columns
WHERE
	pg_stat_user_tables.relname = information_schema.columns.table_name
	AND pg_stat_user_tables.schemaname = $1
ORDER BY
	information_schema.columns.table_name ASC,
	information_schema.columns.ordinal_position ASC
;`
	slog.Debug("executing query", "query", normalizeQuery(query), "schema", schema)

	rows, err := s.DB.QueryContext(ctx, query, schema)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var columns []Column
	for rows.Next() {
		var column Column
		if err := rows.Scan(
			&column.SchemaName,
			&column.TableName,
			&column.ColumnName,
			&column.DataType,
			&column.Position,
			&column.Comment,
		); err != nil {
			return nil, fmt.Errorf("failed to scan columns: %w", err)
		}

		columns = append(columns, column)
	}

	return columns, nil
}

func normalizeQuery(query string) string {
	tabAndNewlineRegex := regexp.MustCompile("[\t\n]")
	replaced := tabAndNewlineRegex.ReplaceAllString(query, " ")

	spaceRegex := regexp.MustCompile(`\s+`)
	replaced = spaceRegex.ReplaceAllString(replaced, " ")

	replaced = strings.Trim(replaced, " ")

	return replaced
}
