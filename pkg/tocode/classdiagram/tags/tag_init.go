package tags

import (
	"github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram/utils"
	"strings"
)

type ParseTag func(note string) (Tag, bool, error)

type Note interface {
	GetText() string
}

var parseTags []ParseTag

func init() {
	parseTags = []ParseTag{
		ParseTagData,
		ParseTagGrid,
		ParseTagEnum,
		ParseTagForm,
		ParseTagAggregate,
		ParseService,
	}
}

func ParseNotes(notes []string) ([]Tag, error) {
	var list []Tag
	for _, text := range notes {
		tags, ok, err := ParseTags(text)
		if err != nil {
			return nil, err
		}
		if ok {
			list = append(list, tags...)
		}
	}
	return list, nil
}

func ParseTags(text string) ([]Tag, bool, error) {
	if !strings.Contains(text, "@") {
		return nil, false, nil
	}
	var list []Tag
	for _, parseTag := range parseTags {
		tag, ok, err := parseTag(text)
		if err != nil {
			return nil, false, err
		} else if ok {
			list = append(list, tag)
		}
	}
	return list, len(list) > 0, nil
}

func GetTagText(text string, tagType TagType) (string, bool) {
	isHas, res, _, _, _ := utils.StringBetween(text, string(tagType), "}")
	if isHas {
		return res[1:], true
	}
	return "", false
}
