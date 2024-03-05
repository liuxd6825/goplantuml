package tags

import (
	"testing"
)

func TestParseTags(t *testing.T) {
	str := "姓名 @data{size:100} @grid{width:200} @form{}"
	tags, err := ParseTags(str)
	if err != nil {
		t.Error(err)
	}
	for _, tag := range tags {
		t.Log(tag)
	}
}
