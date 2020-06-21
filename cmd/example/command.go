package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version should be the semantic version tagged to the binary
	Version = "dev"
	// Commit is the commit hash/identifier which the binary was built from
	Commit = "unknown_commit"
	// Timestamp is the date and time which the binary was built
	Timestamp = "unknown_timestamp"
)

// GetCommand returns the command object associated with the entrypoint
// of the application
func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use: "example",
		Run: run,
	}
	cmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Displays version information",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("%s-%s / %s\n", Version, Commit, Timestamp)
		},
	})
	return &cmd
}

func run(cmd *cobra.Command, args []string) {
	cmd.Help()
}
