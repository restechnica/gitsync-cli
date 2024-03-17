package commands

import (
	"errors"
	"fmt"
	"runtime/debug"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/restechnica/gitsync-cli/internal/ldflags"
)

// NewVersionCommand creates a new version command.
// It prints out useful version information about gitsync.
// Returns the new version command.
func NewVersionCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:  "version",
		RunE: VersionCommandRunE,
	}

	return command
}

// VersionCommandRunE runs the command.
// Returns an error if the command fails.
func VersionCommandRunE(cmd *cobra.Command, args []string) (err error) {
	log.Debug().Str("command", "version").Msg("starting run...")

	var info *debug.BuildInfo
	var ok bool

	if info, ok = debug.ReadBuildInfo(); !ok {
		return errors.New("failed to read build info")
	}

	var arch, os string

	for _, setting := range info.Settings {
		if setting.Key == "GOARCH" {
			arch = setting.Value
		}

		if setting.Key == "GOOS" {
			os = setting.Value
		}
	}

	fmt.Printf(
		"gitsync-cli %s %s %s/%s\n",
		ldflags.Version,
		info.GoVersion,
		os,
		arch,
	)

	return err
}
