package exec

import "github.com/restechnica/gitsync-cli/pkg/cli/commands"

func Run() error {
	var command = commands.NewRootCommand()
	return command.Execute()
}
