package workspace

import (
	"os"

	"github.com/rs/zerolog/log"
)

// Clean cleans up a workspace.
// Returns an error if it failed.
func Clean(workdir string) (err error) {
	log.Debug().Msg("Cleaning up workspace...")
	err = os.RemoveAll(workdir)

	if err == nil {
		log.Debug().Msg("Cleaning all done!")
	}

	return err
}

// CleanNoError clean up a workspace without returning errors.
func CleanNoError(workdir string) {
	_ = Clean(workdir)
}

// Create creates a workspace.
// Returns a directory path or an error if it failed.
func Create() (dir string, err error) {
	log.Debug().Msg("Creating workspace...")
	return os.MkdirTemp("", "tmp.*")
}

// SetUp creates and uses a workspace.
// Returns a directory path or and error if it failed.
func SetUp() (workdir string, err error) {
	if workdir, err = Create(); err != nil {
		return workdir, err
	}

	if err = Use(workdir); err != nil {
		return workdir, err
	}

	return workdir, err
}

// Use uses a workspace by changing to the directory.
// Returns an error if it failed.
func Use(workdir string) (err error) {
	err = os.Chdir(workdir)

	if err == nil {
		log.Debug().Str("path", workdir).Msg("Using workspace")
	}

	return err
}
