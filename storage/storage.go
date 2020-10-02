package storage

import "errors"

type item struct {
	id   uint
	name string
	info map[string]string
}

func (i *item) Id() uint {
	return i.id
}

func (i *item) Name() string {
	return i.name
}

func (i *item) SetName(name string) error {
	i.name = name
	return errors.New("not implemented")
}

func (i *item) Info(key string) (err error, value string) {
	return errors.New("not implemented"), ""
}

func (i *item) SetInfo(key string, value string) error {
	return errors.New("not implemented")
}

func (i *item) Store() (err error, id uint) {
	return errors.New("not implemented"), 0
}

func (i *item) Use() (err error) {
	return errors.New("not implemented")
}


func Item(name string, info map[string]string) (err error, newItem *item) {
	return errors.New("not implemented"), nil
}

func RetrieveItem(id uint) (err error, retrievedItem *item) {
	return errors.New("not implemented"), nil
}