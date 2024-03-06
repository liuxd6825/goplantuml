package tags

import "github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram/utils"

type TagCmd struct {
	BaseTag
}

func NewTagCmd(text string) (*TagCmd, error) {
	tagData := &TagCmd{}
	tagData.TagType = TagTypeCmd
	if err := utils.TagUnmarshal("{"+text+"}", tagData); err != nil {
		return nil, err
	}
	return tagData, nil
}

func ParseTagCmd(text string) (tag Tag, ok bool, err error) {
	dataText, dataHas := GetTagText(text, TagTypeCmd)
	if dataHas {
		tag, err = NewTagQuery(dataText)
	}
	if tag != nil {
		ok = true
	}
	return
}
