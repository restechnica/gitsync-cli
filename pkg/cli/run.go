package cli

func Run() error {
	var command = NewRootCommand()
	return command.Execute()
}
