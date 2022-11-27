package core

import (
	"github.com/rs/zerolog/log"

	"github.com/restechnica/gitsync-cli/pkg/target"
)

type SyncOptions struct {
	// AvailableTargets the available targets for destination and source.
	AvailableTargets []target.Target
	// Destination the id of the destination.
	Destination string
	// Source the id of the source.
	Source string
}

// Sync pulls files from a source and pushes them to a destination.
// It detects how to handle different sources and destinations based on available targets.
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

	if source, err = target.SelectTarget(options.Source, options.AvailableTargets); err != nil {
		return err
	}

	log.Info().
		Str("target", source.GetName()).
		Str("src", options.Source).
		Msg("Pulling from source...")

	if err = source.Pull(options.Source); err != nil {
		return err
	}

	log.Info().Msg("Pull done!")

	log.Info().Msg("Selecting compatible destination...")

	var destination target.Target

	if destination, err = target.SelectTarget(options.Destination, options.AvailableTargets); err != nil {
		return err
	}

	log.Info().
		Str("target", destination.GetName()).
		Str("dst", options.Destination).
		Msg("Pushing to destination...")

	if err = destination.Push(options.Destination); err != nil {
		return err
	}

	log.Info().Msg("Push done!")

	log.Info().Msg("Sync completed!")

	return err
}
