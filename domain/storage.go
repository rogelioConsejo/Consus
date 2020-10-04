package domain

type Storage interface {
	Store(Storable) (err error, id uint)
	Use(Storable) error
}