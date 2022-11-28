package cli

import "github.com/restechnica/gitsync-cli/pkg/target"

var (
	// DefaultTargets the default available targets for sources and destinations.
	DefaultTargets = []target.Target{
		target.NewGitTarget(),
		target.NewS3Target(),
	}
)
