package s3

// API interface to interact with s3.
type API interface {
	Sync(source string, destination string) (output string, err error)
}
