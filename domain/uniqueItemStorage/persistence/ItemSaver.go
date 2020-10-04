package persistence

import (
	"github.com/rogelioConsejo/Consus/domain"
)

type ItemSaver interface {
	Save(*domain.Storable) (error, string)
	Retrieve(string) (error, *domain.Storable)
	Erase(string) error
}