package git

// API interface to interact with git.
type API interface {
	AddConfig(key string, value string) (output string, err error)
	Clone(repository string, path string) (output string, err error)
	FetchAll() (output string, err error)
	InitBareRepository(path string) (output string, err error)
	GetStatus() (output string, err error)
	PushMirror(url string) (output string, err error)
	SetConfig(key string, value string) (output string, err error)
}
