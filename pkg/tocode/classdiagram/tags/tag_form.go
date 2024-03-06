package tags

import (
	"github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram/utils"
)

type TagForm struct {
	BaseTag
}

func NewTagForm(text string) (*TagForm, error) {
	tagForm := &TagForm{}
	tagForm.TagType = TagTypeForm
	if err := utils.TagUnmarshal(text, tagForm); err != nil {
		return nil, err
	}
	return tagForm, nil
}
