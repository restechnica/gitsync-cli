package target

import (
	"fmt"
)

func SelectCompatibleTarget(id string) (target Target, err error) {
	var targets = []Target{
		NewGitTarget(),
	}

	for _, target = range targets {
		if target.IsCompatible(id) {
			return target, err
		}
	}

	return target, fmt.Errorf(`failed to determine target for id "%s"`, id)
}
