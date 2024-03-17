package generator

type Table struct {
	Name    string
	Columns []Column
}

type Column struct {
	Name string
	Type string
}

type SchemaLoader interface {
	LoadSchema() ([]Table, error)
}
