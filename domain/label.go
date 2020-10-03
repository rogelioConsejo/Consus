package domain

type Label interface {
	Name() string
	SetName(name string) error
	Info() map[string]string
	SetInfo(map[string]string)
}