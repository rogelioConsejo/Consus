package uniqueItemStorage

import (
	"errors"
	"strings"
)

type tag struct {
	name string
	info map[string]string
}

func Tag(name string) *tag {
	var t *tag
	t = new(tag)
	t.name = name
	t.info = make(map[string]string)
	return t
}

func (t *tag) Name() string {
	return t.name
}

func (t *tag) SetName(name string) error {
	var err error
	if strings.TrimSpace(name) == "" {
		err = errors.New("invalid name")
	}
	if err == nil {
		t.name = name
	}
	return err
}

func (t *tag) Info() map[string]string {
	return t.info
}

func (t *tag) SetInfo(info map[string]string) {
	t.info = info
}