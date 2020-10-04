package storage

import (
	"errors"
	"github.com/rogelioConsejo/Consus/domain/storable"
)

type storage struct {
}

func Storage() *storage {
	return new(storage)
}

//TODO
func (s *storage) Store(storable storable.Type) (err error, id string) {
	return errors.New("not implemented"), ""
}

//TODO
func (s *storage) Use(id string) (error, storable.Type) {
	return errors.New("not implemented"), nil
}
