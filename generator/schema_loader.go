package generator

import "context"

type Table struct {
	Name    string
	Comment string
	Columns []Column
}

// TODO: Relation
type Column struct {
	Name       string
	Comment    string
	Type       string
	IsNullable bool
	OrderAsc   int
}

type SchemaLoader interface {
	LoadTableSchemas(ctx context.Context) ([]Table, error)
}
