package tags

import "github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram/utils"

type TagEvent struct {
	BaseTag
}

func NewTagEvent(text string) (*TagEvent, error) {
	tagData := &TagEvent{}
	tagData.TagType = TagTypeEvent
	if err := utils.TagUnmarshal("{"+text+"}", tagData); err != nil {
		return nil, err
	}
	return tagData, nil
}

func ParseTagEvent(text string) (tag Tag, ok bool, err error) {
	dataText, dataHas := GetTagText(text, TagTypeEvent)
	if dataHas {
		tag, err = NewTagEvent(dataText)
	}
	if tag != nil {
		ok = true
	}
	return
}
