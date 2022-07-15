package core

import (
	"github.com/rs/zerolog/log"

	"github.com/restechnica/gitsync-cli/pkg/target"
)

type SyncOptions struct {
	// Destination the id of the destination.
	Destination string
	// Source the id of the source.
	Source string
}

// Sync a `source` Git repository into a `destination` Git repository.
// Returns an error if the sync fails.
func Sync(options *SyncOptions) (err error) {
	log.Info().Msg("Starting sync...")

	defer func() {
		if err != nil {
			log.Warn().Msg("An error interrupted the syncing process")
		}
	}()

	log.Info().Msg("Selecting compatible source...")

	var source target.Target

	if source, err = target.SelectCompatibleTarget(options.Source); err != nil {
		return err
	}

	log.Info().
		Str("src", options.Source).
		Str("target", source.GetName()).
		Msg("Pulling from source...")

	if err = source.Pull(options.Source); err != nil {
		return err
	}

	log.Info().Msg("Pull done!")

	log.Info().Msg("Selecting compatible destination...")

	var destination target.Target

	if destination, err = target.SelectCompatibleTarget(options.Destination); err != nil {
		return err
	}

	log.Info().
		Str("dst", options.Destination).
		Str("target", destination.GetName()).
		Msg("Pushing to destination...")

	if err = destination.Push(options.Destination); err != nil {
		return err
	}

	log.Info().Msg("Push done!")

	log.Info().Msg("Sync completed!")

	return err
}
