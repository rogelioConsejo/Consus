package storable

import "testing"

func TestStorable(t *testing.T) {
	var s Type
	const testKey = "testKey"
	const testValue = "testValue"
	label := make(map[string]string)
	label[testKey] = testValue
	s = Storable(label)
	print(s.Label()[testKey])
	if testValue != s.Label()[testKey] {
		t.Error("label error")
	}
}
