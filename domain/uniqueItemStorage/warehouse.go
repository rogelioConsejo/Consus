package uniqueItemStorage

import (
	"errors"
	"github.com/rogelioConsejo/Consus/domain"
)

type warehouse struct {

}

func Warehouse() *warehouse {
	return new(warehouse)
}

//TODO
func (w *warehouse) Store(storable domain.Storable) (err error, id uint) {
	return errors.New("not implemented"), 0
}

//TODO
func (w *warehouse) Use(storable domain.Storable) error  {
	return errors.New("not implemented")
}

