package classdiagram

import (
	"context"
	"github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram/utils"
	"strings"
)

type Namespace struct {
	NameElement
	Nodes       map[string]Element    `json:"nodes,omitempty"`
	Namespaces  map[string]*Namespace `json:"namespaces,omitempty"`
	Connections []*Connection         `json:"connections,omitempty"`
	Parent      *Namespace            `json:"-"`
}

func NewNamespace(line string, parent *Namespace, notes []*Note) *Namespace {
	ns := &Namespace{
		Nodes:       make(map[string]Element),
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

func (n *Namespace) Parse(ctx context.Context, reader ParseReader) (resNotes []*Note, err error) {
	notes := make([]*Note, 0)
	for reader.Scan() {
		line := reader.ReadLine()
		if line == "}" {
			return
		} else if strings.Contains(line, "'") {
			note := NewNote(ctx, line)
			notes = append(notes, note)
			continue
		} else if line == "" {
			continue
		} else {
			notes, err = n.ParseLine(ctx, line, notes, reader)
			if err != nil {
				return
			}
		}
	}
	return
}

func (n *Namespace) ParseLine(ctx context.Context, line string, notes []*Note, reader ParseReader) (resNotes []*Note, err error) {
	resNotes = nil
	if line == "}" || line == "" {
		return notes, nil
	} else if strings.HasPrefix(line, "namespace") {
		ns := NewNamespace(line, n, notes)
		n.Namespaces[ns.Name] = ns
		if _, err = ns.Parse(ctx, reader); err != nil {
			return
		}
	} else if strings.HasPrefix(line, "class") {
		class := NewClass(ctx, line, n.GetNamespaceName(), notes)
		n.Nodes[class.Name] = class
		registerNode(ctx, class)
		if err = class.Parse(ctx, reader); err != nil {
			return
		}
	} else if strings.HasPrefix(line, "enum") {
		enum := NewEnum(ctx, line, n.GetNamespaceName(), notes)
		n.Nodes[enum.Name] = enum
		registerNode(ctx, enum)
		if err = enum.Parse(ctx, reader); err != nil {
			return
		}
	} else if strings.HasPrefix(line, "interface") {
		inter := NewInterface(ctx, line, n.GetNamespaceName(), notes)
		n.Nodes[inter.Name] = inter
		registerNode(ctx, inter)
		if err = inter.Parse(ctx, reader); err != nil {
			return
		}
	} else if IsConnection(line) {
		conn := NewConnection(ctx, line, n.GetNamespaceName(), notes)
		n.Connections = append(n.Connections, conn)
		registerConnection(ctx, conn)
		return
	} else {
		resNotes = notes
	}
	return resNotes, nil
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
