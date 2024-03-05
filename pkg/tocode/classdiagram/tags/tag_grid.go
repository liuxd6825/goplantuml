package tags

import (
	"github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram/utils"
)

type TagGrid struct {
	BaseTag
	Width int
}

func NewTagGrid(text string) (*TagGrid, error) {
	tagGrid := &TagGrid{}
	tagGrid.TagType = TagTypeGrid
	if err := utils.TagUnmarshal("{"+text+"}", tagGrid); err != nil {
		return nil, err
	}
	return tagGrid, nil
}

func ParseTagGrid(text string) (tag Tag, ok bool, err error) {
	dataText, dataHas := GetTagText(text, TagTypeGrid)
	if dataHas {
		tag, err = NewTagGrid(dataText)
	}
	if tag != nil {
		ok = true
	}
	return
}
