package uniqueItemStorage

type tag struct {
	info map[string]string
}

func Tag() *tag {
	var t *tag
	t = new(tag)
	t.info = make(map[string]string)
	return t
}

func (t *tag) Info() map[string]string {
	return t.info
}

func (t *tag) SetInfo(info map[string]string) {
	t.info = info
}