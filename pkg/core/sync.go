package core

import (
	"os"

	"github.com/rs/zerolog/log"

	"github.com/restechnica/gitsync-cli/pkg/git"
)

type SyncOptions struct {
	Destination string
	Source      string
}

func Sync(options *SyncOptions) (err error) {
	log.Info().
		Str("dst", options.Destination).
		Str("src", options.Source).
		Msg("Starting sync...")

	var workdir string
	log.Debug().Msg("Creating workspace...")

	if workdir, err = os.MkdirTemp("", "tmp.*"); err != nil {
		return err
	}

	defer func() {
		if err != nil {
			log.Warn().Msg("An error interrupted the syncing process")
		}

		log.Debug().Msg("Cleaning up workspace...")
		_ = os.RemoveAll(workdir)
		log.Debug().Msg("Cleaning all done!")
	}()

	if err = os.Chdir(workdir); err != nil {
		return err
	}

	log.Debug().Str("path", workdir).Msg("Using workspace")

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

	if output, err = gitAPI.FetchAll(); err != nil {
		return err
	}

	log.Debug().Msg(output)

	if output, err = gitAPI.PushMirror(options.Destination); err != nil {
		return err
	}

	log.Debug().Msg(output)

	log.Info().Msg("Sync completed")

	return err
}
