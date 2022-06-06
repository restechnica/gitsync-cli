package git

import cmder "github.com/restechnica/go-cmder/pkg"

// CLI a git.API to interact with the git CLI.
type CLI struct {
	Commander cmder.Commander
}

// NewCLI creates a new CLI with a commander to run git commands.
// Returns the new CLI.
func NewCLI() CLI {
	return CLI{Commander: cmder.NewExecCommander()}
}

// AddConfig adds a git config.
// Returns an error if the command fails
func (api CLI) AddConfig(key string, value string) (output string, err error) {
	return api.Commander.Output("git", "config", "--add", key, value)
}

func (api CLI) FetchAll() (output string, err error) {
	return api.Commander.Output("git", "fetch", "--all")
}

// InitBareRepository initializes a bare git repository.
// Returns an error if the command fails.
func (api CLI) InitBareRepository(path string) (output string, err error) {
	return api.Commander.Output("git", "init", "--bare", path)
}

// GetStatus gets the git status.
// Returns an error if the command fails.
func (api CLI) GetStatus() (output string, err error) {
	return api.Commander.Output("git", "status")
}

// PushMirror mirror pushes to a remote url.
func (api CLI) PushMirror(url string) (output string, err error) {
	return api.Commander.Output("git", "push", "--mirror", url)
}

// SetConfig sets a git config.
// Returns an error if the command fails
func (api CLI) SetConfig(key string, value string) (output string, err error) {
	return api.Commander.Output("git", "config", key, value)
}
