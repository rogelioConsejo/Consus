package sku

import (
	"github.com/rogelioConsejo/Consus/sku/unit"
	"github.com/rogelioConsejo/Consus/storage/storable"
	"math/big"
)

type Type interface {
	Add(float *big.Float)
	Take(float *big.Float) error
	Empty()
	Amount() *big.Float
}

type sku struct {
	storable storable.Type
	amount   *big.Float
	unit     unit.Unit
}

func Sku(u unit.Unit) *sku {
	var s *sku
	s = new(sku)
	s.unit = u

	var label storable.Label
	label = make(storable.Label)
	s.storable = storable.Storable(label)

	s.amount = big.NewFloat(0)

	return s
}

func (s *sku) Add(float *big.Float)  {
}

func (s *sku) Take(float *big.Float) error{
	return nil
}

func (s *sku) Empty() {
}

func (s *sku) Amount() *big.Float {
	return big.NewFloat(0)
}