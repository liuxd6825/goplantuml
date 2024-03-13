package ddd

import "testing"

func TestBuilder(t *testing.T) {
	err := Builder("aaa", "../../template/ddd", "../../test/ddd/person.puml")
	if err != nil {
		t.Error(err)
	}
}
