package command

import (
	"fmt"

	"github.com/kmtym1998/chair/generator"
	"github.com/kmtym1998/chair/postgres"
	"github.com/spf13/cobra"
)

func NewPostgresCommand() *cobra.Command {
	postgresCmd := &cobra.Command{
		Use:  "postgres",
		Long: "generate Go struct from PostgreSQL table schema",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			pgLoader, err := postgres.New("postgres://postgres:password@localhost:5432/postgres?sslmode=disable")
			if err != nil {
				return fmt.Errorf("failed to create postgres client: %w", err)
			}

			g := generator.New(
				&generator.Config{},
				[]generator.TypeMapping{},
				pgLoader,
			)

			return g.Run(cmd.Context())
		},
	}

	return postgresCmd
}
