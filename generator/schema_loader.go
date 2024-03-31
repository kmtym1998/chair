package generator

import "context"

type Table struct {
	Name    string
	Comment string
	Columns []Column
}

type Column struct {
	Name       string
	Comment    string
	Type       string
	IsNullable bool
	OrderAsc   int
}

type RelationType string

const (
	RelationTypeOneToOne  RelationType = "one_to_one"
	RelationTypeOneToMany RelationType = "one_to_many"
	RelationTypeManyToOne RelationType = "many_to_one"
)

func (r RelationType) String() string {
	return string(r)
}

type SchemaLoader interface {
	LoadTableSchemas(ctx context.Context) ([]Table, error)
}
