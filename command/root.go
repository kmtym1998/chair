package command

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/kmtym1998/chair/generator/config"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:  "chair",
		Long: "chair is a tool to generate Go struct from relational database schema",
		Args: cobra.NoArgs,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cfgFileName, err := cmd.Flags().GetString("config")
			if err != nil {
				return fmt.Errorf("failed to get config flag: %w", err)
			}

			cfg, err := config.Parse(cfgFileName)
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					slog.Warn("config file not found")

					return nil
				}
				return fmt.Errorf("failed to parse config file: %w", err)
			}

			ctx := cmd.Context()
			ctx = config.With(ctx, cfg)
			cmd.SetContext(ctx)

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	const defaultCfgFileName = ".chair.yml"
	rootCmd.PersistentFlags().StringP("config", "c", defaultCfgFileName, "config file path")

	return rootCmd
}
