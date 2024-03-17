package main

import (
	"log"

	"github.com/kmtym1998/chair/command"

	_ "github.com/kr/pretty" // for debugging
)

func main() {
	rootCmd := command.NewRootCommand()
	postgresCmd := command.NewPostgresCommand()

	rootCmd.AddCommand(postgresCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}
