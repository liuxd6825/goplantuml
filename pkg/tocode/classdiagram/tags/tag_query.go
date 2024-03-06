package tags

import "github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram/utils"

type TagQuery struct {
	BaseTag
}

func NewTagQuery(text string) (*TagView, error) {
	tagData := &TagView{}
	tagData.TagType = TagTypeQuery
	if err := utils.TagUnmarshal("{"+text+"}", tagData); err != nil {
		return nil, err
	}
	return tagData, nil
}

func ParseTagQuery(text string) (tag Tag, ok bool, err error) {
	dataText, dataHas := GetTagText(text, TagTypeQuery)
	if dataHas {
		tag, err = NewTagQuery(dataText)
	}
	if tag != nil {
		ok = true
	}
	return
}
