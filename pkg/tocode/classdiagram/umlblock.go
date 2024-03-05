package classdiagram

import (
	"context"
	"strings"
)

type UMLBlock struct {
	FileName      string         `json:"fileName,omitempty"`
	AlisNode      map[string]any `json:"-"`
	Root          *Namespace     `json:"root"`
	allNode       map[string]Element
	allConnection []*Connection
}

func NewUMLBlock(fileName string) *UMLBlock {
	root := NewNamespace("namespace root", nil, nil)
	block := &UMLBlock{
		Root:     root,
		FileName: fileName,
	}
	return block
}

func (b *UMLBlock) Parse(ctx context.Context, reader ParseReader) (err error) {
	err = b.parse(ctx, reader, "@startuml", "@enduml")
	if err == nil {
		b.setConnection(ctx)
	}
	return err
}

func (b *UMLBlock) setConnection(ctx context.Context) {
	pctx := GetParseContext(ctx)
	if pctx != nil {
		//allConnections := pctx.GetConnections()
		//allNodes := pctx.GetNodes()
	}
}

func (b *UMLBlock) getNode(namespace string, nodeName string) Element {
	return nil
}

func (b *UMLBlock) parse(ctx context.Context, reader ParseReader, startTag string, endTag string) (err error) {
	notes := make([]*Note, 0)
	for reader.Scan() {
		line := reader.ReadLine()
		if strings.HasPrefix(line, startTag) {
			for reader.Scan() {
				line := reader.ReadLine()
				if line == "" {
					continue
				} else if strings.Contains(line, "'") {
					notes = append(notes, NewNote(ctx, line))
					continue
				} else if strings.HasPrefix(line, "}") {
					return nil
				}
				notes, err = b.Root.ParseLine(ctx, line, notes, reader)
				if err != nil {
					return err
				}
			}
		}
		if strings.HasPrefix(line, endTag) {
			return nil
		}
	}
	return nil
}
