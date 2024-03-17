package generator

import (
	"context"
)

type Generator struct {
	Config         *Config
	DefaultMappers []TypeMapping
	SchemaLoader   SchemaLoader
}

func New(
	config *Config,
	defaultMappers []TypeMapping,
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

	// Generate code

	return nil
}
