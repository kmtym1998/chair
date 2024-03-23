package generator

import (
	"context"
)

type SchemaLoaderMock struct {
	returnTables []Table
	returnError  error
}

func NewSchemaLoaderMock() SchemaLoader {
	return SchemaLoaderMock{
		returnTables: []Table{
			{
				Name: "character_types",
				Columns: []Column{
					{Name: "id", Type: "integer", IsNullable: false, OrderAsc: 1},
					{Name: "character_value_nullable", Type: "character", IsNullable: true, OrderAsc: 2},
					{Name: "character_varying_value_nullable", Type: "character varying", IsNullable: true, OrderAsc: 3},
					{Name: "text_value_nullable", Type: "text", IsNullable: true, OrderAsc: 4},
					{Name: "character_value", Type: "character", IsNullable: false, OrderAsc: 5},
					{Name: "character_varying_value", Type: "character varying", IsNullable: false, OrderAsc: 6},
					{Name: "text_value", Type: "text", IsNullable: false, OrderAsc: 7},
				},
			},
			{
				Name:    "numeric_types",
				Comment: "numeric types",
				Columns: []Column{
					{Name: "id", Type: "integer", IsNullable: false, OrderAsc: 1},
					{Name: "smallint_value_nullable", Type: "smallint", IsNullable: true, OrderAsc: 2, Comment: "smallint value nullable"},
					{Name: "integer_value_nullable", Type: "integer", IsNullable: true, OrderAsc: 3, Comment: "integer value nullable"},
					{Name: "bigint_value_nullable", Type: "bigint", IsNullable: true, OrderAsc: 4},
					{Name: "decimal_value_nullable", Type: "numeric", IsNullable: true, OrderAsc: 5},
					{Name: "numeric_value_nullable", Type: "numeric", IsNullable: true, OrderAsc: 6},
					{Name: "real_value_nullable", Type: "real", IsNullable: true, OrderAsc: 7},
					{Name: "double_precision_value_nullable", Type: "double precision", IsNullable: true, OrderAsc: 8},
					{Name: "smallint_value", Type: "smallint", IsNullable: false, OrderAsc: 9},
					{Name: "integer_value", Type: "integer", IsNullable: false, OrderAsc: 10},
					{Name: "bigint_value", Type: "bigint", IsNullable: false, OrderAsc: 11},
					{Name: "decimal_value", Type: "numeric", IsNullable: false, OrderAsc: 12},
					{Name: "numeric_value", Type: "numeric", IsNullable: false, OrderAsc: 13},
					{Name: "real_value", Type: "real", IsNullable: false, OrderAsc: 14},
					{Name: "double_precision_value", Type: "double precision", IsNullable: false, OrderAsc: 15},
					{Name: "smallserial_value", Type: "smallint", IsNullable: false, OrderAsc: 16},
					{Name: "serial_value", Type: "integer", IsNullable: false, OrderAsc: 17},
					{Name: "bigserial_value", Type: "bigint", IsNullable: false, OrderAsc: 18},
				},
			},
			{
				Name: "datetime_types",
				Columns: []Column{
					{Name: "id", Type: "integer", IsNullable: false, OrderAsc: 1},
					{Name: "date_value_nullable", Type: "date", IsNullable: true, OrderAsc: 2},
					{Name: "time_value_nullable", Type: "time without time zone", IsNullable: true, OrderAsc: 3},
					{Name: "timestamp_value_nullable", Type: "timestamp without time zone", IsNullable: true, OrderAsc: 4},
					{Name: "timestamptz_value_nullable", Type: "timestamp with time zone", IsNullable: true, OrderAsc: 5},
					{Name: "interval_value_nullable", Type: "interval", IsNullable: true, OrderAsc: 6},
					{Name: "date_value", Type: "date", IsNullable: false, OrderAsc: 7},
					{Name: "time_value", Type: "time without time zone", IsNullable: false, OrderAsc: 8},
					{Name: "timestamp_value", Type: "timestamp without time zone", IsNullable: false, OrderAsc: 9},
					{Name: "timestamptz_value", Type: "timestamp with time zone", IsNullable: false, OrderAsc: 10},
					{Name: "interval_value", Type: "interval", IsNullable: false, OrderAsc: 11},
				},
			},
			{
				Name: "uuid_types",
				Columns: []Column{
					{Name: "id", Type: "integer", IsNullable: false, OrderAsc: 1},
					{Name: "uuid_value_nullable", Type: "uuid", IsNullable: true, OrderAsc: 2},
					{Name: "uuid_value", Type: "uuid", IsNullable: false, OrderAsc: 3},
				},
			},
			{
				Name: "money_types",
				Columns: []Column{
					{Name: "id", Type: "integer", IsNullable: false, OrderAsc: 1},
					{Name: "money_value_nullable", Type: "money", IsNullable: true, OrderAsc: 2},
					{Name: "money_value", Type: "money", IsNullable: false, OrderAsc: 3},
				},
			},
			{
				Name: "boolean_types",
				Columns: []Column{
					{Name: "id", Type: "integer", IsNullable: false, OrderAsc: 1},
					{Name: "boolean_value_nullable", Type: "boolean", IsNullable: true, OrderAsc: 2},
					{Name: "boolean_value", Type: "boolean", IsNullable: false, OrderAsc: 3},
				},
			},
		},
		returnError: nil,
	}
}

func (m SchemaLoaderMock) LoadTableSchemas(ctx context.Context) ([]Table, error) {
	return m.returnTables, m.returnError
}

func (m SchemaLoaderMock) WithTable(returnTables []Table) SchemaLoaderMock {
	m.returnTables = returnTables

	return m
}

func (m SchemaLoaderMock) WithError(returnError error) SchemaLoaderMock {
	m.returnError = returnError

	return m
}
