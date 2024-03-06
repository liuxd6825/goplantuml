package tags

import "github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram/utils"

type TagView struct {
	BaseTag
}

func NewTagView(text string) (*TagView, error) {
	tagData := &TagView{}
	tagData.TagType = TagTypeView
	if err := utils.TagUnmarshal("{"+text+"}", tagData); err != nil {
		return nil, err
	}
	return tagData, nil
}

func ParseTagView(text string) (tag Tag, ok bool, err error) {
	dataText, dataHas := GetTagText(text, TagTypeView)
	if dataHas {
		tag, err = NewTagView(dataText)
	}
	if tag != nil {
		ok = true
	}
	return
}
