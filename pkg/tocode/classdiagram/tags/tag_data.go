package tags

import (
	"github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram/utils"
)

type TagData struct {
	BaseTag
	Size   int    `json:"size,omitempty"`
	Titles Titles `json:"titles,omitempty"`
}

type Titles struct {
	CN string `json:"cn,omitempty"`
	EN string `json:"en,omitempty"`
	DE string `json:"de,omitempty"`
}

func NewTagData(text string) (*TagData, error) {
	tagData := &TagData{}
	tagData.TagType = TagTypeData
	if err := utils.TagUnmarshal(text, tagData); err != nil {
		return nil, err
	}
	return tagData, nil
}
