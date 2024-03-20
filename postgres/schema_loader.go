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
					Name:       column.ColumnName,
					Comment:    column.Comment.String,
					Type:       column.DataType,
					IsNullable: strings.ToUpper(column.IsNullable) != "NO",
					OrderAsc:   column.Position,
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
	DataType   string         `db:"data_type"`
	IsNullable string         `db:"is_nullable"`
	Position   int            `db:"ordinal_position"`
	Comment    sql.NullString `db:"description"`
}

func (s *SchemaLoader) listColumns(ctx context.Context, schema string) ([]Column, error) {
	// TODO: Relation
	const query = `
WITH column_list AS (
SELECT
	c.table_schema || '_' || c.table_name || '_' || c.column_name AS column_key,
	c.table_schema,
	c.table_name,
	c.column_name,
	c.data_type,
	c.is_nullable,
	c.ordinal_position,
	(
		SELECT
			description
		FROM
			pg_description
		WHERE
			pg_description.objoid = pg_stat_user_tables.relid
			AND pg_description.objsubid = c.ordinal_position
	) AS description
FROM
	pg_stat_user_tables,
	information_schema.columns c
WHERE
	pg_stat_user_tables.relname = c.table_name
	AND pg_stat_user_tables.schemaname = $1
),
relation_list AS (
select
	t.table_schema || '_' || t.table_name || '_' || k.column_name AS column_key,
	t.table_schema AS table_schema,
	k.table_name AS from_table_name,
	k.column_name AS from_column_name,
	c.table_name AS to_table_name,
	c.column_name AS to_colmun_name
from
	information_schema.table_constraints as t,
	information_schema.key_column_usage as k,
	information_schema.constraint_column_usage as c
where
	t.constraint_type = 'FOREIGN KEY'
	AND t.constraint_name = k.constraint_name
	AND t.constraint_name = c.constraint_name
)
SELECT
	col.table_schema,
	col.table_name,
	col.column_name,
	col.data_type,
	col.is_nullable,
	col.ordinal_position,
	col.description
FROM
	column_list AS col
	LEFT JOIN relation_list AS rel ON col.column_key = rel.column_key
ORDER BY
	col.table_name ASC,
	col.ordinal_position ASC
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
			&column.IsNullable,
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
