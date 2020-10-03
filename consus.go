package main

import (
	"errors"
	"github.com/rogelioConsejo/Consus/domain"
)

func main() {
}

//TODO
func Store(storable domain.Storable) (error, uint) {
	return errors.New("not implemented"), 0
}

//TODO
func GetName(storableId uint) (error, string) {
	return errors.New("not implemented"), ""
}

//TODO
func Use(storableId uint) error {
	return errors.New("not implemented")
}