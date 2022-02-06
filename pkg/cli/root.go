package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewRootCommand creates a new root command.
// Returns the new root command.
func NewRootCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "gitsync",
	}

	fmt.Println("hello")

	return command
}
