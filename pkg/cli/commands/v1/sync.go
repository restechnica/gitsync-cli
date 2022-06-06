package v1

import (
	"github.com/spf13/cobra"

	"github.com/restechnica/gitsync-cli/pkg/cli"
	"github.com/restechnica/gitsync-cli/pkg/core"
)

// NewSyncCommand creates a new V1 sync command.
// Returns the new V1 sync command.
func NewSyncCommand() *cobra.Command {
	var command = &cobra.Command{
		PreRunE: SyncCommandPreRunE,
		RunE:    SyncCommandRunE,
		Use:     "sync",
	}

	command.Flags().StringVarP(&cli.DestinationPathFlag, "destination-path", "d", "", "")
	command.Flags().StringVarP(&cli.TargetUrlFlag, "target-url", "t", "", "")

	return command
}

// SyncCommandPreRunE runs before the command runs.
// Returns an error if it failed.
func SyncCommandPreRunE(command *cobra.Command, args []string) (err error) {
	return command.MarkFlagRequired("target-url")
}

// SyncCommandRunE runs the command.
// Returns an error if the command fails.
func SyncCommandRunE(command *cobra.Command, args []string) (err error) {
	var options = &core.SyncOptions{
		DestinationPath: command.Flags().Lookup("destination-path").Value.String(),
		TargetUrl:       command.Flags().Lookup("target-url").Value.String(),
	}

	return core.Sync(options)
}
