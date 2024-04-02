package target

import (
	"os"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/restechnica/gitsync-cli/pkg/git"
)

// RemoteGitTarget a Target to pull and push remote git repositories.
type RemoteGitTarget struct {
}

// NewRemoteGitTarget creates a new RemoteGitTarget.
// Returns the new RemoteGitTarget.
func NewRemoteGitTarget() RemoteGitTarget {
	return RemoteGitTarget{}
}

// GetName gets a unique id used for all RemoteGitTarget instances.
// Returns the name.
func (target RemoteGitTarget) GetName() string {
	return "remote-git"
}

// IsCompatible checks whether an id can be used with a RemoteGitTarget.
// Returns true if the id is compatible.
func (target RemoteGitTarget) IsCompatible(id string) bool {
	var isHTTPS = strings.HasPrefix(id, "https")
	var isSSH = strings.HasPrefix(id, "git@")

	return (isHTTPS || isSSH) && strings.HasSuffix(id, ".git")
}

// Pull pulls a remote git repository into a directory.
// The id parameter has to be a valid git origin URL.
// The remote git repository can be a URL or a filesystem path.
// Returns an error if something went wrong.
func (target RemoteGitTarget) Pull(id string, directory string) (err error) {
	var gitAPI git.API = git.NewCLI()
	var output string

	var previousDirectory string

	if previousDirectory, err = os.Getwd(); err != nil {
		return err
	}

	if err = os.Chdir(directory); err != nil {
		return err
	}

	if output, err = gitAPI.InitBareRepository("."); err != nil {
		return err
	}

	log.Debug().Msg(output)

	if output, err = gitAPI.SetConfig("remote.origin.url", id); err != nil {
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

	if err = os.Chdir(previousDirectory); err != nil {
		return err
	}

	return err
}

// Push pushes a directory to a remote git repository.
// The current working directory has to be a git repository.
// The remote git repository can be a URL or a filesystem path.
// Returns an error if something went wrong.
func (target RemoteGitTarget) Push(directory string, id string) (err error) {
	var gitAPI git.API = git.NewCLI()
	var output string

	var previousDirectory string

	if previousDirectory, err = os.Getwd(); err != nil {
		return err
	}

	if err = os.Chdir(directory); err != nil {
		return err
	}

	if output, err = gitAPI.PushMirror(id); err != nil {
		return err
	}

	log.Debug().Msg(output)

	if err = os.Chdir(previousDirectory); err != nil {
		return err
	}

	return err
}

// String converts an RemoteGitTarget to a string representation
// returns a string representation of an RemoteGitTarget
func (target RemoteGitTarget) String() string {
	return target.GetName()
}
