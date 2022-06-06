package exec

import (
	"os"

	"github.com/restechnica/gitsync-cli/pkg/cli/commands"
)

func Run() (err error) {
	var command = commands.NewRootCommand()

	if err = command.Execute(); err != nil {
		os.Exit(1)
	}

	return err
}
