package cli

import "github.com/restechnica/gitsync-cli/pkg/target"

var (
	DefaultTargets = []target.Target{
		target.NewGitTarget(),
	}
)
