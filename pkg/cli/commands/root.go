package commands

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/restechnica/gitsync-cli/pkg/cli"
	v1 "github.com/restechnica/gitsync-cli/pkg/cli/commands/v1"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

// NewRootCommand creates a new root command.
// Returns the new root command.
func NewRootCommand() *cobra.Command {
	var command = &cobra.Command{
		PersistentPreRunE: RootCommandPersistentPreRunE,
		SilenceErrors:     true,
		Use:               "gitsync",
	}

	command.PersistentFlags().BoolVarP(&cli.VerboseFlag, "verbose", "v", false, "increase log level verbosity")

	command.AddCommand(v1.NewV1Command())
	command.AddCommand(v1.NewSyncCommand())
	command.AddCommand(NewVersionCommand())

	return command
}

// RootCommandPersistentPreRunE runs before the command and any subcommand runs.
// Returns an error if it failed.
func RootCommandPersistentPreRunE(cmd *cobra.Command, args []string) (err error) {
	ConfigureLogging()

	return err
}

func ConfigureLogging() {
	SetLogLevel()
}

func SetLogLevel() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if cli.VerboseFlag {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
