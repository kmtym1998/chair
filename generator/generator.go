package generator

import (
	"context"

	"github.com/kmtym1998/chair/generator/config"
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
	_, err := g.SchemaLoader.LoadSchema(ctx)
	if err != nil {
		return err
	}

	// Generate code

	return nil
}
