package tags

import (
	"github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram/utils"
)

type TagAggregate struct {
	BaseTag
}

func NewTagAggregate(text string) (*TagAggregate, error) {
	tagData := &TagAggregate{}
	tagData.TagType = TagTypeAggregate
	if err := utils.TagUnmarshal("{"+text+"}", tagData); err != nil {
		return nil, err
	}
	return tagData, nil
}

func ParseTagAggregate(text string) (tag Tag, ok bool, err error) {
	dataText, dataHas := GetTagText(text, TagTypeAggregate)
	if dataHas {
		tag, err = NewTagAggregate(dataText)
	}
	if tag != nil {
		ok = true
	}
	return tag, ok, err
}

func (t *TagAggregate) GetTagType() TagType {
	return t.TagType
}
