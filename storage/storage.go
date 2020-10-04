package storage

import (
	"github.com/rogelioConsejo/Consus/storage/persistence"
	"github.com/rogelioConsejo/Consus/storage/storable"
)

type storage struct {
	dto StorableDto
}

type StorableDto interface {
	Create(storable.Type) (err error, id string)
	Read(id string) (error, storable.Type)
	Delete(id string) error
}

func Storage() *storage {
	s := new(storage)
	s.dto = persistence.Mock()

	return s
}

func (s *storage) Store(storable storable.Type) (err error, id string) {
	return s.dto.Create(storable)
}

func (s *storage) CheckOut(id string) (err error, storable storable.Type) {
	err, storable = s.dto.Read(id)
	if err == nil {
		err = s.dto.Delete(id)
	}

	return
}
