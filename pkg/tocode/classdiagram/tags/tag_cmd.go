package tags

import "github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram/utils"

type TagCmd struct {
	BaseTag
}

func NewTagCmd(text string) (*TagCmd, error) {
	tagData := &TagCmd{}
	tagData.TagType = TagTypeCmd
	if err := utils.TagUnmarshal(text, tagData); err != nil {
		return nil, err
	}
	return tagData, nil
}
