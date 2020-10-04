package domain

type Storable interface {
	Id() string
	Label() Label
}

type CountableStorable interface {
	Storable
	Count() uint
	Put(uint)
	Take(uint) error
}