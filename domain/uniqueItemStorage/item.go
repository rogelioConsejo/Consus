package uniqueItemStorage

import (
	"github.com/rogelioConsejo/Consus/domain"
)

type item struct {
	id    string
	label domain.Label
}

func (i *item) Id() string {
	return ""
}

func (i *item) Label() (tag domain.Label) {
	return i.label
}

func Item(info map[string]string) *item {

	newItem := new(item)
	newItem.label = Tag()
	newItem.label.SetInfo(info)

	return newItem
}
