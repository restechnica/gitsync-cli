package cli

import (
	"github.com/restechnica/gitsync-cli/pkg/cli/cmd"
)

func Run() error {
	var command = cmd.NewRootCommand()
	return command.Execute()
}
