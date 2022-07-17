package target

import (
	"fmt"
)

func SelectCompatibleTarget(id string, targets []Target) (target Target, err error) {
	for _, target = range targets {
		if target.IsCompatible(id) {
			return target, err
		}
	}

	return target, fmt.Errorf(`failed to select target for id "%s"`, id)
}
