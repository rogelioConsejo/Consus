package persistence

import (
	"errors"
	"fmt"
	"github.com/rogelioConsejo/Consus/storage/storable"
)

type mock struct {
	stored map[string]storable.Type
}

func Mock() *mock {
	m := new(mock)
	m.stored = make(map[string]storable.Type)
	return m
}

func (m *mock) Create(s storable.Type) (err error, id string) {
	var key string
	key = CreateToken(12)
	for m.stored[key] != nil {
		key = CreateToken(10)
	}

	m.stored[key] = s
	return nil, key
}

func (m *mock) Read(id string) (err error, s storable.Type) {
	if m.stored[id] == nil {
		err = errors.New(fmt.Sprintf("storable {%s} not found", id))
	} else {
		s = m.stored[id]
	}

	return err, s
}

func (m *mock) Delete(id string) error{
	delete(m.stored, id)
	return nil
}