package main

import (
	"log"

	command "github.com/kmtym1998/chair/commands"
)

func main() {
	rootCmd := command.NewRootCommand()
	postgresCmd := command.NewPostgresCommand()

	rootCmd.AddCommand(postgresCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}
