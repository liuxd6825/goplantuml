package classdiagram

import (
	"context"
	"github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram/utils"
	"strings"
)

type Namespace struct {
	NameElement
	Elements    map[string]Element    `json:"elements,omitempty"`
	Namespaces  map[string]*Namespace `json:"namespaces,omitempty"`
	Connections []*Connection         `json:"connections,omitempty"`
	Parent      *Namespace            `json:"-"`
}

func NewNamespace(line string, parent *Namespace, notes []*Comment) *Namespace {
	ns := &Namespace{
		Elements:    make(map[string]Element),
		Connections: make([]*Connection, 0),
		Namespaces:  make(map[string]*Namespace),
		Parent:      parent,
	}
	ns.InitName(utils.ParseNamespaceName(line))
	namespace := "root"
	if parent != nil {
		namespace = parent.GetNamespaceName() + "." + ns.Name
	}

	ns.InitBase(line, "namespace", namespace, notes)
	return ns
}

func (n *Namespace) Parse(ctx context.Context, reader ParseReader) (resNotes []*Comment, err error) {
	notes := make([]*Comment, 0)
	var note *Note
	var element Element
	for reader.Scan() {
		line := reader.ReadLine()
		if line == "}" {
			return
		} else if strings.Contains(line, "'") {
			note := NewComment(ctx, line)
			notes = append(notes, note)
			continue
		} else if line == "" {
			continue
		} else if strings.HasPrefix(line, "note") {
			notes, element, err = n.ParseLine(ctx, line, notes, reader)
			if element != nil {
				if n, ok := element.(*Note); ok {
					note = n
				}
			}
		} else if utils.StringIsEndNote(line) {

		} else {
			notes, element, err = n.ParseLine(ctx, line, notes, reader)
			if element != nil {
				note = nil
			} else if element == nil && note != nil {
				note.AddText(line)
			} else {
				note = nil
			}
			if err != nil {
				return
			}
		}
	}
	return
}

func (n *Namespace) ParseLine(ctx context.Context, line string, comments []*Comment, reader ParseReader) (resComments []*Comment, element Element, err error) {
	resComments = nil
	element = nil
	if line == "}" || line == "" {
		return comments, element, nil
	} else if strings.HasPrefix(line, "namespace") {
		ns := NewNamespace(line, n, comments)
		n.Namespaces[ns.Name] = ns
		element = ns
		if _, err = ns.Parse(ctx, reader); err != nil {
			return
		}
	} else if strings.HasPrefix(line, "class") {
		class := NewClass(ctx, line, n.GetNamespaceName(), comments)
		n.Elements[class.Name] = class
		element = class
		registerNode(ctx, class)
		if err = class.Parse(ctx, reader); err != nil {
			return
		}
	} else if strings.HasPrefix(line, "enum") {
		enum := NewEnum(ctx, line, n.GetNamespaceName(), comments)
		n.Elements[enum.Name] = enum
		element = enum
		registerNode(ctx, enum)
		if err = enum.Parse(ctx, reader); err != nil {
			return
		}
	} else if strings.HasPrefix(line, "interface") {
		inter := NewInterface(ctx, line, n.GetNamespaceName(), comments)
		n.Elements[inter.Name] = inter
		element = inter
		registerNode(ctx, inter)
		if err = inter.Parse(ctx, reader); err != nil {
			return
		}
	} else if IsConnection(line) {
		conn := NewConnection(ctx, line, n.GetNamespaceName(), comments)
		element = conn
		n.Connections = append(n.Connections, conn)
		registerConnection(ctx, conn)
		return
	} else if strings.HasPrefix(line, "note") {
		note := NewNote(line)
		registerNode(ctx, note)
		element = note
	} else {
		resComments = comments
		element = nil
	}
	return resComments, element, nil
}

func registerNode(ctx context.Context, node Element) {
	pctx := GetParseContext(ctx)
	if pctx != nil {
		pctx.AddNode(node)
	}
}
func registerConnection(ctx context.Context, conn ...*Connection) {
	pctx := GetParseContext(ctx)
	if pctx != nil {
		pctx.AddConnection(conn...)
	}
}
