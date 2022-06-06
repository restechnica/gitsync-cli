package core

import (
	"fmt"
	"os"

	"github.com/restechnica/gitsync-cli/pkg/git"
)

type SyncOptions struct {
	DestinationPath string
	TargetUrl       string
}

func Sync(options *SyncOptions) (err error) {
	var workdir string

	if workdir, err = os.MkdirTemp("", "tmp.*"); err != nil {
		return err
	}

	if err = os.Chdir(workdir); err != nil {
		return err
	}

	var gitAPI git.API = git.NewCLI()
	var output string

	if output, err = gitAPI.InitBareRepository("."); err != nil {
		return err
	}

	if output, err = gitAPI.SetConfig("remote.origin.url", options.TargetUrl); err != nil {
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

	if output, err = gitAPI.PushMirror(options.DestinationPath); err != nil {
		return err
	}

	fmt.Println(output)

	return os.RemoveAll(workdir)
}
