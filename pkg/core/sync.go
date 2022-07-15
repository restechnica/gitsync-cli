package core

import (
	"github.com/rs/zerolog/log"

	"github.com/restechnica/gitsync-cli/pkg/git"
)

type SyncOptions struct {
	Destination string
	Source      string
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

	var gitAPI git.API = git.NewCLI()
	var output string

	if output, err = gitAPI.InitBareRepository("."); err != nil {
		return err
	}

	log.Debug().Msg(output)

	if output, err = gitAPI.SetConfig("remote.origin.url", options.Source); err != nil {
		return err
	}

	if output, err = gitAPI.AddConfig("remote.origin.fetch", "+refs/heads/*:refs/heads/*"); err != nil {
		return err
	}

	if output, err = gitAPI.AddConfig("remote.origin.fetch", "+refs/tags/*:refs/tags/*"); err != nil {
		return err
	}

	if output, err = gitAPI.AddConfig("remote.origin.fetch", "+refs/notes/*:refs/notes/*"); err != nil {
		return err
	}

	if output, err = gitAPI.SetConfig("remote.origin.mirror", "true"); err != nil {
		return err
	}

	log.Info().Str("src", options.Source).Msg("Fetching everything from source...")

	if output, err = gitAPI.FetchAll(); err != nil {
		return err
	}

	log.Debug().Msg(output)
	log.Info().Msg("Fetch done!")

	log.Info().Str("dst", options.Destination).Msg("Pushing to destination...")

	if output, err = gitAPI.PushMirror(options.Destination); err != nil {
		return err
	}

	log.Debug().Msg(output)
	log.Info().Msg("Push done!")

	log.Info().Msg("Sync completed")

	return err
}
