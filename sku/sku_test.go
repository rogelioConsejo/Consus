package sku

import (
	"errors"
	"github.com/rogelioConsejo/Consus/sku/unit"
	"math/big"
	"testing"
)

func TestSKU(t *testing.T) {
	var sku Type
	var u unit.Unit
	u.Name = "TestUnit"
	sku = Sku(u)

	ten := big.NewFloat(10)
	twenty := big.NewFloat(20)

	sku.Add(ten)
	if sku.Amount().Cmp(ten) != 0  {
		t.Error(errors.New("error adding to sku"))
	}
	sku.Add(ten)
	if sku.Amount().Cmp(twenty) != 0 {
		t.Error(errors.New("error adding again to sku"))
	}

	err := sku.Take(ten)
	if sku.Amount().Cmp(ten) != 0 || err != nil {
		t.Error(errors.New("error taking from sku"))
	}

	thirtyMillion := big.NewFloat(30000000)
	pointTwo := big.NewFloat(.2)
	thirtyMillionAndTenPointTwo := big.NewFloat(30000010.2)
	sku.Add(thirtyMillion)
	sku.Add(pointTwo)
	if sku.Amount().Cmp(thirtyMillionAndTenPointTwo) != 0 {
		t.Error(errors.New("error adding big and precise number to sku"))
	}

	zero := big.NewFloat(0)
	sku.Empty()
	if sku.Amount().Cmp(zero) != 0 {
		t.Error(errors.New("error emptying sku"))
	}

	one := big.NewFloat(1)
	err = sku.Take(one)
	if err == nil {
		t.Error(errors.New("did not throw an error when taking from an empty sku"))
	}
}
