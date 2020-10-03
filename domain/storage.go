package domain

type Storage interface {
	Store(storable Storable) (err error, id uint)
	Use(storable Storable) error
}