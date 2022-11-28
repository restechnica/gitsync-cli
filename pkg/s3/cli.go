package s3

import cmder "github.com/restechnica/go-cmder/pkg"

// CLI a s3.API to interact with the aws s3 CLI.
type CLI struct {
	Commander cmder.Commander
}

// NewCLI creates a new CLI with a commander to run s3 commands.
// Returns the new CLI.
func NewCLI() *CLI {
	return &CLI{Commander: cmder.NewExecCommander()}
}

// Sync syncs to and from s3.
// Returns an error if the command fails.
func (api CLI) Sync(source string, destination string) (output string, err error) {
	return api.Commander.Output("aws", "s3", "sync", source, destination, "--delete")
}
