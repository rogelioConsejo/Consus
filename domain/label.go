package domain

type Label interface {
	Info() map[string]string
	SetInfo(map[string]string)
}