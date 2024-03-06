package tags

import (
	"github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram/utils"
)

type TagService struct {
	BaseTag
}

func NewTagService(text string) (*TagService, error) {
	tagData := &TagService{}
	tagData.TagType = TagTypeService
	if err := utils.TagUnmarshal("{"+text+"}", tagData); err != nil {
		return nil, err
	}
	return tagData, nil
}

func ParseTagService(text string) (tag Tag, ok bool, err error) {
	dataText, dataHas := GetTagText(text, TagTypeService)
	if dataHas {
		tag, err = NewTagService(dataText)
	}
	if tag != nil {
		ok = true
	}
	return
}
