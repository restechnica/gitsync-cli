package target

type Target interface {
	GetName() string
	IsCompatible(id string) bool
	Pull(id string, directory string) error
	Push(directory string, id string) error
	String() string
}
