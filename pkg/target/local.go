package target

import (
	"path/filepath"

	"github.com/restechnica/gitsync-cli/pkg/git"
	"github.com/rs/zerolog/log"
)

const LocalGit = "local-git"

// LocalGitTarget a Target to pull and push local git repositories.
type LocalGitTarget struct {
}

// NewLocalGitTarget creates a new LocalGitTarget.
// Returns the new LocalGitTarget.
func NewLocalGitTarget() LocalGitTarget {
	return LocalGitTarget{}
}

// GetName gets a unique id used for all LocalGitTarget instances.
// Returns the name.
func (target LocalGitTarget) GetName() string {
	return LocalGit
}

// IsCompatible checks whether an id can be used with a LocalGitTarget.
// Returns true if the id is compatible.
func (target LocalGitTarget) IsCompatible(id string) bool {
	var isAbsolutePath = filepath.IsAbs(id)
	var isLocalPath = filepath.IsLocal(id)

	return isAbsolutePath || isLocalPath
}

// Pull pulls a local git repository into the current working directory.
// The id parameter has to be a valid filesystem path.
// The resulting repository is a full repository.
// Returns an error if something went wrong.
func (target LocalGitTarget) Pull(id string, directory string) (err error) {
	var gitAPI git.API = git.NewCLI()
	var output string

	if output, err = gitAPI.Clone(id, directory); err != nil {
		return err
	}

	log.Debug().Msg(output)

	return err
}

// Push pushes the current working directory to a local git repository.
// The current working directory has to be a git repository.
// The id parameter has to be a valid filesystem path.
// The resulting repository is a full repository.
// Returns an error if something went wrong.
func (target LocalGitTarget) Push(directory string, id string) (err error) {
	var gitAPI git.API = git.NewCLI()
	var output string

	if output, err = gitAPI.Clone(directory, id); err != nil {
		return err
	}

	log.Debug().Msg(output)

	return err
}

// String converts an LocalGitTarget to a string representation
// returns a string representation of an LocalGitTarget
func (target LocalGitTarget) String() string {
	return target.GetName()
}
