package generator

import "context"

type Table struct {
	Name    string
	Columns []Column
}

type Column struct {
	Name string
	Type string
}

type SchemaLoader interface {
	LoadSchema(ctx context.Context) ([]Table, error)
}
