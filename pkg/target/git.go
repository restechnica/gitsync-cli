package target

import (
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/restechnica/gitsync-cli/pkg/git"
)

type GitTarget struct {
}

func NewGitTarget() GitTarget {
	return GitTarget{}
}

func (target GitTarget) GetName() string {
	return "git"
}

func (target GitTarget) IsCompatible(id string) bool {
	return strings.HasSuffix(id, ".git")
}

func (target GitTarget) Pull(id string) (err error) {
	var gitAPI git.API = git.NewCLI()
	var output string

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

	return err
}

// Push pushes the current working directory to a remote git repository.
// The current working directory has to be a git repository.
// The remote git repository can be a URL or a filesystem path.
// Returns an error if something went wrong.
func (target GitTarget) Push(id string) (err error) {
	var gitAPI git.API = git.NewCLI()
	var output string

	if output, err = gitAPI.PushMirror(id); err != nil {
		return err
	}

	log.Debug().Msg(output)

	return err
}