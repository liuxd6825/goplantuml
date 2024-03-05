package tags

import (
	"github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram/utils"
)

type TagData struct {
	BaseTag
	Size   int    `json:"size"`
	Title  string `json:"title"`
	Titles Titles `json:"titles"`
}

type Titles struct {
	CN string `json:"cn"`
	EN string `json:"en"`
	DE string `json:"de"`
}

func NewTagData(text string) (*TagData, error) {
	tagData := &TagData{}
	tagData.TagType = TagTypeData
	if err := utils.TagUnmarshal("{"+text+"}", tagData); err != nil {
		return nil, err
	}
	if tagData.Title != "" && tagData.Titles.CN == "" {
		tagData.Titles.CN = tagData.Title
	}
	return tagData, nil
}

func ParseTagData(text string) (tag Tag, ok bool, err error) {
	dataText, dataHas := GetTagText(text, TagTypeData)
	if dataHas {
		tag, err = NewTagData(dataText)
	}
	if tag != nil {
		ok = true
	}
	return
}

func (t *TagData) GetTagType() TagType {
	return TagTypeData
}
