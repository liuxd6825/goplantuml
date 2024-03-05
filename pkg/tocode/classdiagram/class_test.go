package classdiagram

import (
	"testing"
)

func Test_ParseClassName(t *testing.T) {
	res := ParseClassName("class pack.UserMapper <<User,Entity >> { ", "class")
	t.Log(res.NamespaceName)
	t.Log(res.Name)
	t.Log(res.GenericType)
}
