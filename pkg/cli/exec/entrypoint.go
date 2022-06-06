package exec

import (
	"os"

	"github.com/rs/zerolog/log"

	"github.com/restechnica/gitsync-cli/pkg/cli/commands"
)

func Run() (err error) {
	var command = commands.NewRootCommand()

	if err = command.Execute(); err != nil {
		log.Error().Err(err).Msg("")
		os.Exit(1)
	}

	return err
}
