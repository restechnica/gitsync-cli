package cli

import (
	v1 "github.com/restechnica/gitsync-cli/pkg/cli/v1"
	"github.com/spf13/cobra"
)

// NewRootCommand creates a new root command.
// Returns the new root command.
func NewRootCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "gitsync",
	}

	command.AddCommand(v1.NewV1Command())
	command.AddCommand(v1.NewVersionCommand())

	return command
}
