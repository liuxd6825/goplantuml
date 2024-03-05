package tags

import (
	"github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram/utils"
)

type TagRepository struct {
	BaseTag
}

func NewTagRepository(text string) (*TagRepository, error) {
	tagData := &TagRepository{}
	tagData.TagType = TagTypeRepository
	if err := utils.TagUnmarshal("{"+text+"}", tagData); err != nil {
		return nil, err
	}
	return tagData, nil
}

func ParseTagRepository(text string) (tag Tag, ok bool, err error) {
	dataText, dataHas := GetTagText(text, TagTypeRepository)
	if dataHas {
		tag, err = NewTagRepository(dataText)
	}
	if tag != nil {
		ok = true
	}
	return
}
