package command

import (
	"fmt"
	"log"
	"os"

	"github.com/kmtym1998/chair/generator"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:  "chair",
		Long: "chair is a tool to generate Go struct from relational database schema",
		Args: cobra.NoArgs,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			log.Println("PersistentPreRunE")

			cfgFileName, err := cmd.Flags().GetString("config")
			if err != nil {
				return fmt.Errorf("failed to get config flag: %w", err)
			}

			log.Println("cfgFileName", cfgFileName)
			cfgFile, err := os.ReadFile(cfgFileName)
			if err != nil {
				log.Fatalf("failed to read config file: %v", err)
			}

			var cfg generator.Config
			if err := yaml.Unmarshal(cfgFile, &cfg); err != nil {
				log.Fatalf("failed to unmarshal config: %v", err)
			}

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
