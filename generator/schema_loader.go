package generator

import "context"

type Table struct {
	Name    string
	Comment string
	Columns []Column
}

type Column struct {
	Name     string
	Comment  string
	Type     string
	OrderAsc int
}

type SchemaLoader interface {
	LoadTableSchemas(ctx context.Context) ([]Table, error)
}
