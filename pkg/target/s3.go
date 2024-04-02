package target

import (
	"strings"

	"github.com/restechnica/gitsync-cli/pkg/s3"
	"github.com/rs/zerolog/log"
)

// S3Target a Target to pull and push S3 resources.
type S3Target struct {
}

// NewS3Target creates a new S3Target.
// Returns the new S3Target.
func NewS3Target() S3Target {
	return S3Target{}
}

// GetName gets a unique id used for all S3Target instances.
// Returns the name.
func (target S3Target) GetName() string {
	return "s3"
}

// IsCompatible checks whether an id can be used with an S3Target.
// Returns true if the id is compatible.
func (target S3Target) IsCompatible(id string) bool {
	return strings.HasPrefix(id, "s3")
}

// Pull pulls an S3 bucket resource into a directory.
// The id parameter has to be a valid S3 URI.
// Returns an error if something went wrong.
func (target S3Target) Pull(id string, directory string) (err error) {
	var s3API = s3.NewCLI()
	var output string

	if output, err = s3API.Sync(id, directory); err != nil {
		return err
	}

	log.Debug().Msg(output)

	return err
}

// Push pushes a directory to an S3 bucket.
// The id parameter is an S3 URI.
// Returns an error if something went wrong.
func (target S3Target) Push(directory string, id string) (err error) {
	var s3API = s3.NewCLI()
	var output string

	if output, err = s3API.Sync(directory, id); err != nil {
		return err
	}

	log.Debug().Msg(output)

	return err
}

// String converts an S3Target to a string representation
// returns a string representation of an S3Target
func (target S3Target) String() string {
	return target.GetName()
}
