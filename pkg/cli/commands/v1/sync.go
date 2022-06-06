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

	command.Flags().StringVarP(&cli.DestinationFlag, "destination", "d", "", "")
	command.Flags().StringVarP(&cli.SourceFlag, "source", "s", "", "")

	return command
}

// SyncCommandPreRunE runs before the command runs.
// Returns an error if it failed.
func SyncCommandPreRunE(command *cobra.Command, args []string) (err error) {
	if err = command.MarkFlagRequired("destination"); err != nil {
		return err
	}

	if err = command.MarkFlagRequired("source"); err != nil {
		return err
	}

	return err
}

// SyncCommandRunE runs the command.
// Returns an error if the command fails.
func SyncCommandRunE(command *cobra.Command, args []string) (err error) {
	var options = &core.SyncOptions{
		Destination: command.Flags().Lookup("destination").Value.String(),
		Source:      command.Flags().Lookup("source").Value.String(),
	}

	// silence usage because errors at this point are unrelated to CLI usage errors
	command.SilenceUsage = true

	return core.Sync(options)
}
