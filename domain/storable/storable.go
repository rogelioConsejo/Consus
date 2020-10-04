package storable

//This interface is used only to enforce the use of the storable.storable constructor, so it is only used to declare types
type Type interface {
	Id() string
	Label() Label
}

type storable struct {
	id    string
	label Label
}

type Label map[string]string

func (i *storable) Id() string {
	return ""
}

func (i *storable) Label() Label {
	return i.label
}

func Storable(info Label) *storable {

	newItem := new(storable)

	if info == nil {
		newItem.label = make(map[string]string)
	} else {
		newItem.label = info
	}


	return newItem
}
