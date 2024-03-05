package tags

import (
	"github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram/utils"
)

type TagEnum struct {
	BaseTag
}

func NewTagEnum(text string) (*TagEnum, error) {
	tagEnum := &TagEnum{}
	tagEnum.TagType = TagTypeEnum
	if err := utils.TagUnmarshal("{"+text+"}", tagEnum); err != nil {
		return nil, err
	}
	return tagEnum, nil
}

func ParseTagEnum(text string) (tag Tag, ok bool, err error) {
	dataText, dataHas := GetTagText(text, TagTypeEnum)
	if dataHas {
		tag, err = NewTagEnum(dataText)
	}
	if tag != nil {
		ok = true
	}
	return
}

func (t *TagEnum) GetTagType() TagType {
	return t.TagType
}
