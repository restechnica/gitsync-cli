package v1

import (
	"fmt"
	"github.com/spf13/cobra"
)

// NewSyncCommand creates a new V1 sync command.
// Returns the new V1 sync command.
func NewSyncCommand() *cobra.Command {
	var command = &cobra.Command{
		Run: SyncCommandRun,
		Use: "sync",
	}

	return command
}

func SyncCommandRun(cmd *cobra.Command, args []string) {
	fmt.Println("syncing...")
}
