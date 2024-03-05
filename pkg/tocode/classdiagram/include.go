package classdiagram

import "context"

type Include struct {
	BaseElement
	FileName string `json:"fileName"`
}

func NewInclude(line string, namespaceName string, notes []*Note) *Include {
	i := &Include{}
	i.InitBase(line, "include", namespaceName, notes)
	return i
}

func (b *Include) Parse(ctx context.Context, reader ParseReader) (err error) {
	return nil
}
