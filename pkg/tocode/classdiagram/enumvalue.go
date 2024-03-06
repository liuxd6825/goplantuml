package classdiagram

import "context"

type EnumValue struct {
	BaseElement
	Name string `json:"name"`
	Line string `json:"-"`
}

func NewEnumValue(ctx context.Context, line string, namespaceName string, notes []*Comment) *EnumValue {
	member := &EnumValue{
		Name: line,
		Line: line,
	}
	member.InitBase(line, "EnumValue", namespaceName, notes)
	return member
}
