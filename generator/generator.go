package generator

import (
	"context"

	"github.com/kmtym1998/chair/generator/config"
	"github.com/kr/pretty"
)

type Generator struct {
	Config         *config.Config
	DefaultMappers []config.TypeMapping
	SchemaLoader   SchemaLoader
}

func New(
	config *config.Config,
	defaultMappers []config.TypeMapping,
	schemaLoader SchemaLoader,
) *Generator {
	return &Generator{
		Config:         config,
		DefaultMappers: defaultMappers,
		SchemaLoader:   schemaLoader,
	}
}

func (g *Generator) Run(ctx context.Context) error {
	// Load schema
	tables, err := g.SchemaLoader.LoadTableSchemas(ctx)
	if err != nil {
		return err
	}

	pretty.Println(tables)

	// Generate code

	return nil
}
