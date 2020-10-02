package Consus

type Storable interface {
	Id() uint
	Name() string
	SetName(name string) error
	Info(key string) (err error, value string)
	SetInfo(key string, value string) error
	Store() (err error, id uint)
	Use() (err error)
}