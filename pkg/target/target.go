package target

type Target interface {
	GetName() string
	IsCompatible(id string) bool
	Pull(id string) error
	Push(id string) error
}
