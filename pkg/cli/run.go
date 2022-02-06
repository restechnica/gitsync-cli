package cli

import (
	"github.com/restechnica/gitsync-cli/pkg/cli/commands"
)

func Run() error {
	var cmd = commands.NewRootCommand()
	return cmd.Execute()
}
