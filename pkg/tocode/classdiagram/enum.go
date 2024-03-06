package classdiagram

import (
	"context"
	"github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram/utils"
	"strings"
)

type Enum struct {
	NameElement
	Members []*EnumValue `json:"members,omitempty"`
}

func NewEnum(ctx context.Context, line string, namespace string, notes []*Comment) *Enum {
	enum := &Enum{
		Members: make([]*EnumValue, 0),
	}
	name, alias, ns := utils.ParseEnumName(line)
	if namespace != "" && ns != "" {
		ns = "." + ns
	}
	enum.InitBase(line, "enum", namespace+ns, notes)
	enum.InitName(name, alias)
	enum.InitDataTag(name)
	return enum
}

func (e *Enum) Parse(ctx context.Context, reader ParseReader) error {
	var notes []*Comment
	for reader.Scan() {
		line := reader.ReadLine()
		if line == "}" {
			return nil
		} else if strings.HasPrefix(line, "'") {
			notes = append(notes, NewComment(ctx, line))
		} else {
			e.AddEnumValue(NewEnumValue(ctx, line, e.GetNamespaceName(), notes))
			notes = nil
		}
	}
	return nil
}

func (e *Enum) AddEnumValue(member *EnumValue) {
	e.Members = append(e.Members, member)
}
