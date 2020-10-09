package sku

import (
	"errors"
	"fmt"
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
	s.amount.Add(s.amount, float)
}

func (s *sku) Take(float *big.Float) error{
	var err error

	if s.amount.Cmp(float) < 0 {
		err = errors.New(fmt.Sprintf("%s is not enough to take %s from sku", s.amount.String(), float.String()))
	} else {
		s.amount.Sub(s.amount, float)
	}

	return err
}

func (s *sku) Empty() {
	_ = s.Take(s.Amount())
}

func (s *sku) Amount() *big.Float {
	return s.amount
}