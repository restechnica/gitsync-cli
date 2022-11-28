package target

import (
	"fmt"
	"strings"
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

// Pull pulls an S3 bucket resource into the current working directory.
// The id parameter has to be a valid S3 URI.
// Returns an error if something went wrong.
func (target S3Target) Pull(id string) (err error) {
	err = fmt.Errorf("pulling an S3 target is not supported yet")
	return err
}

// Push pushes the current working directory to an S3 bucket.
// The id parameter is an S3 URI.
// Returns an error if something went wrong.
func (target S3Target) Push(id string) (err error) {
	err = fmt.Errorf("pushing an S3 target is not supported yet")
	return err
}

// String converts an S3Target to a string representation
// returns a string representation of an S3Target
func (target S3Target) String() string {
	return target.GetName()
}