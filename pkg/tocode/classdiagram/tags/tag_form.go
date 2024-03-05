package tags

import (
	"github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram/utils"
)

type TagForm struct {
	BaseTag
}

func NewTagForm(text string) (*TagForm, error) {
	tagForm := &TagForm{}
	tagForm.TagType = TagTypeForm
	if err := utils.TagUnmarshal("{"+text+"}", tagForm); err != nil {
		return nil, err
	}
	return tagForm, nil
}

func ParseTagForm(text string) (tag Tag, ok bool, err error) {
	dataText, dataHas := GetTagText(text, TagTypeForm)
	if dataHas {
		tag, err = NewTagForm(dataText)
	}
	if tag != nil {
		ok = true
	}
	return
}
