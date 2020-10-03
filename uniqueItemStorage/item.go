package uniqueItemStorage

import (
	"github.com/rogelioConsejo/Consus/domain"
)

type item struct {
	label domain.Label
}

func (i *item) GetLabel() (tag domain.Label) {
	return i.label
}

func Item(name string, info map[string]string) (error, *item) {
	var err error

	newItem	:= new(item)
	newItem.label = Tag(name)
	newItem.label.SetInfo(info)
	err = newItem.label.SetName(name)

	return err, newItem
}