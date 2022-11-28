package v1

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/restechnica/gitsync-cli/pkg/cli"
	"github.com/restechnica/gitsync-cli/pkg/core"
	"github.com/restechnica/gitsync-cli/pkg/workspace"
)

// NewSyncCommand creates a new V1 sync command.
// Returns the new V1 sync command.
func NewSyncCommand() *cobra.Command {
	var command = &cobra.Command{
		PreRunE: SyncCommandPreRunE,
		RunE:    SyncCommandRunE,
		Short:   "Syncs two Git repositories.",
		Use:     "sync",
	}

	command.Flags().StringVarP(&cli.DestinationFlag, "destination", "d", "", "the destination Git repository")
	command.Flags().StringVarP(&cli.SourceFlag, "source", "s", "", "the source Git repository")

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
	// silence usage and errors because errors at this point are unrelated to CLI usage errors
	command.SilenceErrors = true
	command.SilenceUsage = true

	var workdir string

	if workdir, err = workspace.SetUp(); err != nil {
		return err
	}

	defer workspace.CleanNoError(workdir)

	var options = &core.SyncOptions{
		AvailableTargets: cli.DefaultTargets,
		Destination:      command.Flags().Lookup("destination").Value.String(),
		Source:           command.Flags().Lookup("source").Value.String(),
	}

	log.Debug().Msg(fmt.Sprintf("available sync targets: %v", options.AvailableTargets))

	if err = core.Sync(options); err != nil {
		log.Error().Err(err).Msg("")
	}

	return err
}
