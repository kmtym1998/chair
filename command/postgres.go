package command

import (
	"fmt"

	"github.com/kmtym1998/chair/generator"
	"github.com/kmtym1998/chair/generator/config"
	"github.com/kmtym1998/chair/postgres"
	"github.com/spf13/cobra"
)

func NewPostgresCommand() *cobra.Command {
	postgresCmd := &cobra.Command{
		Use:  "postgres",
		Long: "generate Go struct from PostgreSQL table schema",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			dsn, err := cmd.Flags().GetString("dsn")
			if err != nil {
				return fmt.Errorf("failed to get dsn flag: %w", err)
			}

			cfg, _ := config.From(cmd.Context())

			pgLoader, err := postgres.NewSchemaLoader(dsn, cfg.Postgres.Schema)
			if err != nil {
				return fmt.Errorf("failed to create postgres client: %w", err)
			}

			g := generator.New(
				cfg,
				postgres.DefaultMappers(),
				pgLoader,
			)

			return g.Run(cmd.Context())
		},
	}

	postgresCmd.Flags().String("dsn", "", "PostgreSQL data source name")
	if err := postgresCmd.MarkFlagRequired("dsn"); err != nil {
		panic(err)
	}

	return postgresCmd
}
